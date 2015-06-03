package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	url := "https://dl.dropbox.com/s/92fkimwaceaw5sn/Slack_2015-06-03_6_pm-50-32.png"
	json := "{ \"text\": \"" + url + "\"}"
	fmt.Println("Returning: ", json)
	io.WriteString(w, json+"\n")
}

func main() {
	http.HandleFunc("/nobueno", HelloServer)
	port := os.Getenv("PORT")
	fmt.Println("Listening on port: ", port+"\n")
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
