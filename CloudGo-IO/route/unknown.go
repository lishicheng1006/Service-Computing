package route

import "net/http"

func init() {
	http.HandleFunc("/unknown", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
	})
}
