package main

import(
	"fmt" 
	"net/http"
	"log"
)

func helloHandler(w http.ResponseWriter, r* http.Request){
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method != "GET"{
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "This is cool innit!")
}

func formHandler(w http.ResponseWriter, r* http.Request){
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}


func main(){
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	// fmt.Println("Server failed to start")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		// fmt.Println("Server failed to start")
		// fmt.Println(err)
		log.Fatal(err)
	}
}