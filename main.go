package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/DE-Karpov/pngtime/converter"
)

var port = flag.String("port", "6769", "Port to listen on")

func main() {
	flag.Parse()

	log.Printf("flag args: %s", os.Args[1:])

	http.HandleFunc("/now", handlerTime)
	*port = ":" + *port
	if err := http.ListenAndServe(*port, nil); err != nil {
        log.Fatal(err)
    }

}

func handlerTime(w http.ResponseWriter, r *http.Request) {
	k, ok := r.URL.Query()["k"]
	if !ok {
		log.Fatalln("Invalid URL query")
		return
	}

	log.Printf("k: %s", k)

	koef, err := strconv.Atoi(k[0])
	if err != nil {
		log.Fatalln("Wrong coefficient")
	}

	converter.BuildTimeInPng(koef)

	fileBytes, err := ioutil.ReadFile("tmp/img.png")
	if err != nil {
		panic(err)
	}
	
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(fileBytes)
	
}
