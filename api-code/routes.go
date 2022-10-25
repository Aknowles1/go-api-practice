package main

import (
	"fmt"
	"net/http"
)

type application struct {
}

func startAPI() {
	fmt.Println("Starting webserver....")
	InitialiseDB()
	app := application{}
	//create new mux
	mux := http.NewServeMux()

	//route traffic that hits /getbooks to the getBooks methodß
	mux.HandleFunc("/incriment", app.IncrementHitCount)
	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./html"))))
	mux.HandleFunc("/templatedhome", app.homePage)
	mux.HandleFunc("/edit", app.editHome)
	mux.HandleFunc("/save", app.saveHome)
	//Advertise webserver on localhost:8080, so localhost:8080/incriment
	//will trigger incrimentHitCount() method
	http.ListenAndServe(":8080", mux)
}
