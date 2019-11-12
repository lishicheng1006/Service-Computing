package controller

import (
	"encoding/json"
	"github.com/Jiahonzheng/CloudGo-IO/dto"
	"github.com/Jiahonzheng/CloudGo-IO/errors"
	"github.com/Jiahonzheng/CloudGo-IO/model"
	"github.com/Jiahonzheng/CloudGo-IO/service/user"
	"html/template"
	"net/http"
)

type UserController struct{}

var userService = new(user.Service)

func handleError(w http.ResponseWriter, err error) {
	if errorCode, ok := errors.FromErrorCode(err); ok {
		w.WriteHeader(errorCode.HTTPStatusCode)
		_ = json.NewEncoder(w).Encode(errorCode)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (uc *UserController) SignUp(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	payload := &dto.UserSignUp{}
	err := json.NewDecoder(req.Body).Decode(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err, u := userService.SignUp(payload)
	if err != nil {
		handleError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(errors.OKCode(u))
}

func (uc *UserController) GetInfo(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	q := req.URL.Query()
	err, u := userService.GetUserInfo(q.Get("stu_id"))
	if err != nil {
		u = &model.User{}
	}
	tmpl, _ := template.ParseFiles("template/user_info.tmpl")
	_ = tmpl.Execute(w, u)
	w.Header().Set("Content-Type", "text/html")
}
