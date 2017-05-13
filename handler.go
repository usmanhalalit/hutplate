package hutplate

import (
	"net/http"
	"log"
)

type Handler func(hut Http) interface{}

func (handleFunc Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hut := NewHttp(w, r)
	response := handleFunc(hut)
	switch response.(type) {
	case error:
		if Config.HandleError != nil {
			Config.HandleError(response.(error), hut)
			break
		}

		http.Error(w, "Oops! An error occurred", 500)
		log.Print(response.(error).Error())
	case string:
		w.Write([]byte(response.(string)))
	default:

	}
}
