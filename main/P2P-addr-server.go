package main

import (
	"log"
	"net/http"
	"strings"
)

var(
	mp map[string] string
)


func main() {
	mp = make(map[string]string)
	http.HandleFunc("/SetAddr", SetAddr)
	http.HandleFunc("/GetAddr", GetAddr)
	http.HandleFunc("/RemoveAddr", RemoveAddr)
	if err := http.ListenAndServe(":20017", nil); err != nil {
		log.Fatal(err)
	}
}

func RemoveAddr(writer http.ResponseWriter, request *http.Request) {
	var err error
	err = request.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}
	delete(mp, request.Form.Get("UID"))
	_,err = writer.Write([]byte("ok"))
	if err != nil {
		log.Println(err)
		return
	}
}

func GetAddr(writer http.ResponseWriter, request *http.Request) {
	var err error
	err = request.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}
	addr,ok := mp[request.Form.Get("UID")]
	if !ok {
		log.Printf("there is no UID : %s",request.Form.Get("UID"))
		return
	}
	_,err = writer.Write([]byte(addr))
	if err != nil {
		log.Println(err)
		return
	}
}

func SetAddr(writer http.ResponseWriter, request *http.Request) {
	var err error
	err = request.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}
	Addr := strings.Split(request.RemoteAddr, ":")[0]
	Addr = Addr + ":" + request.Form.Get( "Port")
	mp[request.Form.Get("UID")] = Addr
	log.Printf("UID=%s,Addr=%s", request.Form.Get("UID"), Addr)
	_,err = writer.Write([]byte("ok"))
	if err != nil {
		log.Println(err)
		return
	}
}
