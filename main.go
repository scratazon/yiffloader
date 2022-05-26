package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type E6Generated struct {
	Posts []struct {
		Sample struct {
			Url string
		}
		Tags struct {
			General []string
			Species []string
		}
	}
}

func ec(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}

func main() {

	jsonURL := "https://e621.net/posts.json"
	ua := "Never gonna give you up / I don't know how to use JSON sorry"

	client := &http.Client{}

	tags := "?tags=shark+intersex+gaping_anus"
	rq, e := http.NewRequest("GET", jsonURL+tags, nil)
	ec(e)

	rq.Header.Add("User-Agent", ua)

	r, e := client.Do(rq)
	ec(e)

	defer r.Body.Close()

	var bj E6Generated

	in, e := ioutil.ReadAll(r.Body)
	ec(e)

	json.Unmarshal(in, &bj)

	for i := 0; i <= 75; i++ {
		time.Sleep(1000 * time.Millisecond)
		if bj.Posts[i].Sample.Url == "" {
			fmt.Printf("Post %d: ROODPOSTING\n", i)
		}
		for _, v := range bj.Posts[i].Tags.Species {
			fmt.Print(v, " ")
		}
		for _, v := range bj.Posts[i].Tags.General {
			fmt.Print(v, " ")
		}
		time.Sleep(1000 * time.Millisecond)
	}
}
