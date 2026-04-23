package main

import (
	"fmt"
	"log"
	"net/http"
)

var storedValue string
var db = DatabaseInit()

func main() {
	//DatabaseReCreateTables()
	//DatabaseInsertStat()
	//DatabaseInsertPacketa()
	//DatabaseInsertDPD()
	//DatabaseInsertPosta()
	http.HandleFunc("/", home)
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))
	//http.HandleFunc("/add", add)

	fmt.Println("Server běží na http://localhost:80")
	log.Fatal(http.ListenAndServe(":80", nil))
}
