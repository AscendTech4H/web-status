package main

import (
	"net/http"
	"flag"
	"fmt"
	"time"
	"log"
)

func check(tocheck string) bool {
	g, err := (&http.Client{Timeout: time.Second*3}).Get(tocheck)
	if err != nil {
		log.Printf("GET ERROR: %s\n", err.Error())
		return false
	}
	defer g.Body.Close()
	if g.StatusCode != http.StatusOK {
		log.Printf("BAD STATUS: %s\n", g.StatusCode)
		return false
	}
	return true
}

func main() {
	flag.Parse()
	checks := flag.Args()
	http.ListenAndServe(":8000", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, chk := range checks {
			color := "red"
			if check(chk) {
				color = "green"
			}
			fmt.Fprintf(w, `<p style="color:%s">%s</p>`, color, chk)
		}
	}))
}
