package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request)  {
    if r.URL.Path != "/hello"{
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	if r.Method != "GET"{
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
	}
	fmt.Fprintf(w,"Hello World By Kunal!")
}

func formHandler(w http.ResponseWriter, r *http.Request)  {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() error: %v\n", err)
		return
	}
	fmt.Fprintf(w, "POST Request Successful!\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name: %s\n", name)
	fmt.Fprintf(w, "Address: %s\n", address)
}

func main()  {
	fmt.Println("Simple Web Server")
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting a Server at 8080 port...")
	if err := http.ListenAndServe(":8080", nil); err != nil{
		fmt.Println("Error starting server:", err)
        return
	}
}	