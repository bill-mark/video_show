package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

type middleWareHandler struct {
	r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r
	return m
}

func (m middleWareHandler) ServerHTTP(w http.ResponseWriter, r *http.Request){
	validateUserSession(r)
    m.r.ServerHTTP(w,r)
}


func RegisterHandlers() * httprouter.Router{
	router := httprouter.New()

	router.POST("/user",CreateUser)

	router.POST("/user/:user_name",Login)

	return router
}

func main(){
	 r := RegisterHandlers()
	 mh := NewMiddleWareHandler(r)
	 http.ListenAndServe(":8091",mh)
}

