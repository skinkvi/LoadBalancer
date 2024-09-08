package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func(port string) {
			defer wg.Done()

			http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, "Hello from server %s", port)
			})

			fmt.Printf("Server running on %s\n", port)
			http.ListenAndServe(port, nil)
		}(":" + fmt.Sprintf("%d", 8080+i))
	}

	wg.Wait()
}
