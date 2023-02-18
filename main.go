package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("HELLO WEB SERVER")
	fileServer:=http.FileServer(http.Dir("./static"))
	http.Handle("/",fileServer)
	http.HandleFunc("/hello",helloHandler)
	http.HandleFunc("/form",formHandler)
	if err:=http.ListenAndServe(":8080",nil); err!=nil{
		log.Fatal(err)
	}
}
func helloHandler(w http.ResponseWriter,r *http.Request){
	if r.URL.Path!="/hello"{
		http.Error(w,"404 page not found",http.StatusNotFound)
		return
	}else if r.Method!="GET"{
		http.Error(w,"method not supported",http.StatusNotFound)
		return
	}
	fmt.Fprintf(w,"HELLO!")
}

func formHandler(w http.ResponseWriter,r *http.Request){
	if err:=r.ParseForm();err!=nil{
		fmt.Fprintf(w,"ParseForm error : %v",err)
	}
	name:=r.FormValue("name")
	address:=r.FormValue("address")
	fmt.Fprintf(w,"name : %v\n",name)
	fmt.Fprintf(w,"name : %v",address)
}
