package main

import (
	"io/ioutil"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func hiHandler(w http.ResponseWriter, r *http.Request) {
	leakyFunction()
	resp, err := http.Get("https://swapi.dev/api/people/1")
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	log.Printf(sb)
	w.Write([]byte(sb))

}

func main() {
	http.HandleFunc("/", hiHandler)
	http.ListenAndServe(":8080", nil)
}

func leakyFunction() []string {
	s := make([]string, 3)
	for i := 0; i < 10000000; i++ {
		s = append(s, "magical pandas")
		if (i % 100000) == 0 {
			time.Sleep(500 * time.Millisecond)
		}
	}
	return s
}
