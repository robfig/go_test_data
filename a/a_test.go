package a

import (
	"log"
	"os"
	"testing"
)

func init() {
	f, err := os.Open("config.txt")
	if err != nil {
		log.Fatalf("init: %v", err)
	}
	f.Close()
}

func TestA(t *testing.T) {
	f, err := os.Open("config.txt")
	if err != nil {
		t.Errorf("TestA: error opening config.txt: %v", err)
	}
	f.Close()
}
