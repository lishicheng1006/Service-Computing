package route

import "net/http"

func init() {
	http.Handle("/", http.FileServer(http.Dir("public")))
}
