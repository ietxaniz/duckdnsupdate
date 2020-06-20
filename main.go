package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"
)

func main() {
	changed, err := regexp.Compile("UPDATE")
	if err != nil {
		log.Fatalln(err)
	}
	newIP, err := regexp.Compile("\\d+\\.\\d+\\.\\d+\\.\\d+")
	if err != nil {
		log.Fatalln(err)
	}
	for true {
		url := fmt.Sprintf("https://duckdns.org/update?domains=%s&token=%s&verbose=true", os.Getenv("DOMAINS"), os.Getenv("TOKEN"))
		resp, err := http.Get(url)
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		if changed.Match(body) {
			log.Printf("UPDATE %s", string(newIP.Find(body)))
		}
		time.Sleep(20 * time.Second)
	}
}
