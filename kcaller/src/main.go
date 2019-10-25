package main

import (
	"fmt"
	"net/http"
	"os"
)
var cnt int = 0

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("unable to get hostname")
	}
	cnt++
	fmt.Fprintf(w, "Go Hello on %s:%d\n", hostname, cnt)
}
func HealthzHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
}
func main()  {
	fmt.Println("cfg")
	http.HandleFunc("/", HelloHandler)
	http.HandleFunc("/healthz", HealthzHandler)

	http.ListenAndServe(":8080", nil)
}
