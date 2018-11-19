package main

import (
	"io"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func CreateUser(w http.ResponseWriter,r *http.Request,p httprouter.Params){
   io.WriteString(w,"create user handler")
}

func Login(w http.ResponseWriter,r *http.Request,p httprouter.Params){
  usname := p.ByName("user_name")
  io.WriteString(w,usname)
}

