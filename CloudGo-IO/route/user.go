package route

import (
	"github.com/Jiahonzheng/CloudGo-IO/controller"
	"net/http"
)

var userController = new(controller.UserController)

func init() {
	http.HandleFunc("/api/user", userController.GetInfo)
	http.HandleFunc("/api/user/register", userController.SignUp)
}
