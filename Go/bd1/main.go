package main

import (
	"bd1/model"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	_ "os"
)

func Decoder1(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var genre = model.Genre{}
	err := decoder.Decode(&genre)
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println(genre)
}

func Decoder2(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var rrating = model.RRating{}
	err := decoder.Decode(&rrating)
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println(rrating)
}

func Decoder3(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var sale = model.Sale{}
	err := decoder.Decode(&sale)
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println(sale)
}

func Decoder4(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var mode = model.Mode{}
	err := decoder.Decode(&mode)
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println(mode)
}

func Decoder5(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var devphaze = model.DevPhaze{}
	err := decoder.Decode(&devphaze)
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println(devphaze)
}

func Decoder6(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var order = model.Order{}
	err := decoder.Decode(&order)
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println(order)
}

func Decoder7(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var game = model.Genre{}
	err := decoder.Decode(&game)
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println(game)
}

func Hello2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Kirill!")
}

func main() {

	//msgHandler := msg("Hello from Web Server in Go")
	fmt.Println("Server is listening...")

	router := mux.NewRouter()
	router.HandleFunc("/users", Decoder1)
	router.HandleFunc("/rrating", Decoder2)
	router.HandleFunc("/sale", Decoder3)
	router.HandleFunc("/mode", Decoder4)
	router.HandleFunc("/devphaze", Decoder5)
	router.HandleFunc("/order", Decoder6)
	router.HandleFunc("/game", Decoder7)
	router.HandleFunc("/game", Decoder7)
	http.ListenAndServe("localhost:8080", router)
}
