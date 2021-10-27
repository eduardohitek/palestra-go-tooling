package main

import (
	"math/rand"
	"net/http"
	_ "net/http/pprof"
)

type Person struct {
	Name string
	ID   int
}

func randPerson() *Person {
	name := make([]byte, 50)
	for i := 0; i < 50; i++ {
		name[i] = byte(65 + rand.Intn(25))
	}
	id := rand.Intn(10e6)
	return &Person{
		Name: string(name),
		ID:   id,
	}
}

func leakyFunction(w http.ResponseWriter, r *http.Request) {
	n := int(10e6)
	people := make(map[int]*Person)
	for i := 1; i < n; i++ {
		person := randPerson()
		people[person.ID] = person
	}
	println(len(people), "#####", n)
	w.Write([]byte("done"))
}

func main() {
	http.HandleFunc("/", leakyFunction)
	http.ListenAndServe(":8080", nil)
}
