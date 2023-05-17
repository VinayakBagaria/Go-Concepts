package contextual

import (
	"fmt"
	"net/http"
	"time"
)

func StartRequestServer() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		fmt.Println("Processing request")

		select {
		case <-time.After(4 * time.Second):
			fmt.Fprintf(w, "Request accepted")
		case <-ctx.Done():
			err := ctx.Err()
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.ListenAndServe(":8000", nil)
}
