package hutplate

import "net/http"

type Response struct {
	http.ResponseWriter
	request *http.Request
	session Session
}

func (response Response) Redirect(url string, code ...int) Response {
	redirectCode := 302
	if len(code) > 0 {
		redirectCode = code[0]
	}
	http.Redirect(response, response.request, url, redirectCode)
	return response
}

func (response Response) With(key string, value interface{}) Response {
	response.session.SetFlash(key, value)
	return response
}
