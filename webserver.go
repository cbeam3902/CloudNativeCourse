package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	mux.HandleFunc("/list", db.list)
	mux.HandleFunc("/price", db.price)
	mux.HandleFunc("/create", db.create)
	mux.HandleFunc("/update", db.update)
	mux.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

var rwm sync.RWMutex

func (db database) list(w http.ResponseWriter, req *http.Request) {
	rwm.RLock()
	defer rwm.RUnlock()
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	rwm.RLock()
	defer rwm.RUnlock()
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	rwm.Lock()
	defer rwm.Unlock()
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	tempPrice, e := strconv.ParseFloat(price, 32)
	if e != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Price not supported: %s\n", price)
	} else if _, ok := db[item]; ok {
		w.WriteHeader(http.StatusNotAcceptable)
		fmt.Fprintf(w, "%s is already in the database!\nPlease use update instead.\n", item)
	} else {
		db[item] = dollars(tempPrice)
		fmt.Fprintf(w, "%s added in the database with price %s\n", item, dollars(tempPrice))
	}
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	rwm.Lock()
	defer rwm.Unlock()
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	tempPrice, e := strconv.ParseFloat(price, 32)
	if e != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Price not supported: %s\n", price)
	} else if _, ok := db[item]; ok {
		db[item] = dollars(tempPrice)
		fmt.Fprintf(w, "%s is now %s\n", item, dollars(tempPrice))
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	rwm.Lock()
	defer rwm.Unlock()
	item := req.URL.Query().Get("item")
	if _, ok := db[item]; ok {
		delete(db, item)
		fmt.Fprintf(w, "%s was removed from the database\n", item)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}

}
