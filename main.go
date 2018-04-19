// gohttps redirects it's requests to https.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func redirectHandler(port string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		to := "https://" + strings.Split(r.Host, ":")[0]
		if port != "443" {
			to += ":" + port
		}
		to += r.URL.Path
		if r.URL.RawQuery != "" {
			to += "?" + r.URL.RawQuery
		}
		http.Redirect(w, r, to, http.StatusTemporaryRedirect)
	}
}

var Usage = `usage: gohttps addr [https-port=443]

  gohttps :80
  gohttps :80 8443
`

func main() {
	args := os.Args[1:]
	if len(args) == 0 || len(args) > 2 {
		fmt.Fprintln(os.Stderr, Usage)
		os.Exit(1)
	}
	addr := args[0]
	httpsPort := "443"
	if len(args) == 2 {
		httpsPort = args[1]
	}
	if _, err := strconv.Atoi(httpsPort); err != nil {
		fmt.Fprintln(os.Stderr, Usage)
		fmt.Fprintln(os.Stderr, "https-port follows same origin policy. In example, do '443' instead of ':443'.")
		os.Exit(1)
	}
	http.HandleFunc("/", redirectHandler(httpsPort))
	log.Fatal(http.ListenAndServe(addr, nil))
}
