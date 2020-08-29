package config

import (
	"log"
	"testing"
)

func TestGetSelfIP(t *testing.T) {
	log.Println(GetSelfIP())
}