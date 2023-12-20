package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/", Whois)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Whois(rw http.ResponseWriter, rq *http.Request) {
	if len(rq.URL.Query()) != 2 {
		name := "whois"
		c := exec.Command(name, rq.URL.Query().Get("s"))
		stdout, err := c.Output()
		if err != nil {
			fmt.Fprintf(rw, "%s\n", err)
			log.Panic(err)
		}
		_, err = rw.Write(stdout)
		if err != nil {
			rw.Write([]byte(err.Error()))
			rw.WriteHeader(500)
		}
	} else {
		fmt.Fprintf(rw, "%s\n", "Minimum argument count is 2 !!!")
	}
}
