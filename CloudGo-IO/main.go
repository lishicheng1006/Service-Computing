package main

import "net/http"
import _ "github.com/Jiahonzheng/CloudGo-IO/route"

func main() {
	_ = http.ListenAndServe(":4444", nil)
}
