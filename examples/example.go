package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	log.SetFlags(log.Llongfile | log.LstdFlags | log.Lmicroseconds)
	log.Println("こんにちわーるど")

	var port = flag.String("p", "8080", "-p: port number")
	flag.Parse()

	routeMap := map[string]func(http.ResponseWriter, *http.Request){
		"get-started":        getStarted,
		"line-simple":        lineSimple,
		"line-smooth":        lineSmooth,
		"confidence-band":    confidenceBand,
		"bar-simple":         barSimple,
		"candlestick-simple": candlestickSimple,
		"candlestick-large":  candlestickLarge,
		"mix-line-bar":       mixLineBar,
	}
	for path, f := range routeMap {
		http.HandleFunc("/"+path, f)
		log.Println("http://localhost:" + *port + "/" + path)
	}

	log.Println(http.ListenAndServe(":"+*port, nil))
}
