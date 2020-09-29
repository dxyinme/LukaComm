package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

func httpCall(name string) {
	resp, err := http.Get("http://localhost:8080/?name=" + name)
	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
		return
	} else {
		body, errRead := ioutil.ReadAll(resp.Body)
		if errRead != nil {
			log.Println(errRead)
			return
		}
		if string(body) != "echo" + name {
			log.Println(errRead)
		}
	}
}

func main()  {
	var reqCount int = 25000
	startTime := time.Now()
	var wg sync.WaitGroup
	wg.Add(reqCount)
	for i := 0 ; i < reqCount; i ++ {
		go func() {
			httpCall("UUUUU")
			wg.Done()
		}()
	}
	wg.Wait()
	log.Printf("use time : %v ms", time.Now().Sub(startTime).Milliseconds())
}