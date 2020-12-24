package httpserver

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Person struct {
	Name string
	Age  int16
	Mail string
	City string
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req)
	if req.Method == "GET" {
		p := Person{
			Age:  27,
			Name: "Martin",
			Mail: "martin.steinhauer@stud.ur.de",
			City: "Regensburg",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(p)
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}

func RunHttpServer() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}
