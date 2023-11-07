package main

import (
	"flag"
	"log"
	"net/http"
	"path"
)

var (
	hubs      map[string]*Hub = make(map[string]*Hub)
	terminate chan *Hub       = make(chan *Hub)
)

var addr = flag.String("addr", ":8080", "http service address")

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	// if r.URL.Path != "/" {
	// 	http.Error(w, "Not found", http.StatusNotFound)
	// 	return
	// }
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "www/index.html")
}

func main() {
	flag.Parse()
	go cleanEmptyHub()

	http.HandleFunc("/", serveHome)
	http.HandleFunc("/rooms/", func(w http.ResponseWriter, r *http.Request) {
		hubName := path.Base(r.URL.Path)
		hub, exists := hubs[hubName]
		if !exists {
			log.Println("Creating room:" + hubName)
			hub = newHub(hubName, terminate)
			go hub.run()
			hubs[hubName] = hub
		}
		serveWs(hub, w, r)
	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func cleanEmptyHub() {
	for {
		select {
		case hub := <-terminate:
			log.Println("Deleting hub:", hub.name)
			log.Printf("Total %d clients after %d seconds.", len(hub.clients), 10)
			delete(hubs, hub.name)
		}
	}
}
