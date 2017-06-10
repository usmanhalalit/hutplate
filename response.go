package hutplate

import "net/http"

type response struct {
	http.ResponseWriter
	request *http.Request
	session session
}

func newResponse(w http.ResponseWriter, r *http.Request, s session) response {
	return response{
		w,
		r,
		s,
	}
}

func (r response) Redirect(url string, code ...int) response {
	redirectCode := 302
	if len(code) > 0 {
		redirectCode = code[0]
	}
	http.Redirect(r, r.request, url, redirectCode)
	return r
}

func (r response) With(key string, value interface{}) response {
	r.session.SetFlash(key, value)
	return r
}
