package main

import (
	"fmt"
	"net/http"
	"sync"

	"load_balancer/internal/config"
)

func main() {
	c, err := config.LoadConfig("config/config.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}

	var wg sync.WaitGroup
	for _, server := range c.Servers {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()

			http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, "Hello from server %s", url)
			})

			fmt.Printf("Server running on %s\n", url)
			http.ListenAndServe(url[7:], nil)
		}(server.URL)
	}

	wg.Wait()
}
