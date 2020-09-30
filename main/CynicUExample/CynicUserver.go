package main

import (
	CynicUServer "github.com/dxyinme/LukaComm/CynicU/Server"
	"log"
)

func main() {
	s := &CynicUServer.Server{}
	gServe := s.NewCynicUServer(":8080", "Luka")
	if err := gServe.Serve(s.Lis); err != nil {
		log.Fatal(err)
	}
}
