package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func getUserinfo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello ai ma user")

}
func getUserinfoSecond(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "I am Here")
}
func main() {
	http.HandleFunc("/", getUserinfo)
	http.HandleFunc("/getUserinfo-second", getUserinfoSecond)
	fmt.Println("Server is running at http://localhost:8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("server failed", err)
	}

}
