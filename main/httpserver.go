package main

import (
	"log"
	"net/http"
)

func test(w http.ResponseWriter, r *http.Request) {
	var (
		err error
		name string
	)
	err = r.ParseForm()
	name = r.Form.Get("name")
	_, err = w.Write([]byte("echo" + name))
	if err != nil {
		log.Println(err)
	}
}

func main()  {
	http.HandleFunc("/", test)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}