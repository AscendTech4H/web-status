package main

import (
	"net/http"
	"flag"
	"fmt"
	"time"
)

func check(tocheck string) bool {
	g, err := (&http.Client{Timeout: time.Second*3}).Get(tocheck)
	if err != nil {
		return false
	}
	defer g.Body.Close()
	return g.StatusCode == http.StatusOK
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
