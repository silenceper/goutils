package goutils

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

//PostFormWithFile 上传带文件的表单
//ref:https://stackoverflow.com/questions/20205796/
func PostFormWithFile(client *http.Client, url string, values map[string]io.Reader) (*http.Response, error) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	for key, r := range values {
		var fw io.Writer
		var err error
		if x, ok := r.(io.Closer); ok {
			defer x.Close()
		}
		if x, ok := r.(*os.File); ok {
			if fw, err = w.CreateFormFile(key, x.Name()); err != nil {
				return nil, err
			}
		} else {
			//  添加其他字段
			if fw, err = w.CreateFormField(key); err != nil {
				return nil, err
			}
		}
		if _, err := io.Copy(fw, r); err != nil {
			return nil, err
		}

	}

	err := w.Close()
	if err != nil {
		return nil, err
	}

	// Now that you have a form, you can submit it to your handler.
	req, err := http.NewRequest("POST", url, &b)
	if err != nil {
		return nil, err
	}
	// Don't forget to set the content type, this will contain the boundary.
	req.Header.Set("Content-Type", w.FormDataContentType())
	return client.Do(req)
}
