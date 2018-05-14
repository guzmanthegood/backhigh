package main

import (
	"fmt"
    "net/http"
    "backhigh/models"
)

func main() {
    models.InitDB("postgres://backhigh:backhigh@localhost/backhigh")
    http.HandleFunc("/highlights", highlightsIndex)
    http.ListenAndServe(":8080", nil)
}

func highlightsIndex(w http.ResponseWriter, r *http.Request) {
    if r.Method != "GET" {
        http.Error(w, http.StatusText(405), 405)
        return
    }

	hls, err := models.AllHighlights()
    if err != nil {
        http.Error(w, http.StatusText(500), 500)
        return
    }
	
	for _, rw := range hls {
        fmt.Fprintf(w, "%v, %v, %v, $%.2f\n", rw.Id, rw.Title, rw.Country, rw.Price)
	}
}