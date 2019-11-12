package route

import "net/http"

func init() {
	http.HandleFunc("/api/unknown", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
		_, _ = w.Write([]byte("Method Not Implemented!"))
	})
}
