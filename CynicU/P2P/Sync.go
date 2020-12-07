package P2P

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

var (
	addrServer = flag.String("addrServer", "localhost:8080", "addrServer")
)



// get user's port from addr-server
func GetAddr(UID string) (addr string) {
	resp, err := http.PostForm("http://" + *addrServer + "/GetAddr", url.Values{
		"UID": {UID},
	})
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	addr = strings.TrimSpace(string(body))
	return
}


// set now user addr to addr-server
func SyncAddr(UID, uAddr string) (err error) {
	resp, err := http.PostForm("http://" + *addrServer + "/SetAddr", url.Values{
		"UID": {UID},
		"Port": {uAddr},
	})
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
	return
}

func RemoveAddr(UID string) (err error) {

	return
}