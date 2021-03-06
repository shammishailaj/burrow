package main

import (
	"burrow"
	"burrow/lru"
	"log"
	"net/http"
)

var db = map[string]string{
	"6.824": "MIT",
	"15213": "CMU",
	"15445": "CMU",
}

func main() {
	burrow.NewBurrow("test", 5, burrow.FuncGetter(
		func(key string) (lru.Value, bool) {
			log.Println("Fetch data from datasource by: ", key)
			if v, ok := db[key]; ok {
				return v, true
			}
			return nil, false
		}))
	servers := []string{"localhost:5001", "localhost:5002", "localhost:5003"}
	for _, serverURL := range servers {
		server := burrow.NewHTTPPoolWithServers(serverURL, servers)
		go func(serverURL string) {
			http.ListenAndServe(serverURL, server)
		}(serverURL)
	}
	select {}
}
