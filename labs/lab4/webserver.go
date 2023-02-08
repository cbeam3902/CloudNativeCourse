package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func main() {
	var db database                                      // Database is a struct that contains a map and a mutex
	db.dbm = map[string]dollars{"shoes": 50, "socks": 5} // Initialize the database (Would need to use make(map[string]dollars) normally)

	mux := http.NewServeMux()                             // Function mux
	mux.HandleFunc("/list", db.list)                      // Mux handle for list
	mux.HandleFunc("/price", db.price)                    // Mux handle for price
	mux.HandleFunc("/create", db.create)                  // Mux handle for create
	mux.HandleFunc("/update", db.update)                  // Mux handle for update
	mux.HandleFunc("/delete", db.delete)                  // Mux handle for delete
	log.Fatal(http.ListenAndServe("localhost:8000", mux)) // Begin and log the server
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) } // Have dollars print correctly

type database struct { // Database struct
	dbm map[string]dollars // Uses a map of key type string and value type dollars
	rwm sync.RWMutex       // Mutex for this database instance
}

// var rwm sync.RWMutex

func (db *database) list(w http.ResponseWriter, req *http.Request) { // List every item and it's price in the database
	db.rwm.RLock()                    // Read lock
	defer db.rwm.RUnlock()            // Unlock at the end
	for item, price := range db.dbm { // Collect every item and price in the database
		fmt.Fprintf(w, "%s: %s\n", item, price) // Print every item and price
	}
}

func (db *database) price(w http.ResponseWriter, req *http.Request) { // List the price for the requested item
	db.rwm.RLock()                      // Read lock
	defer db.rwm.RUnlock()              // Unlock at the end
	item := req.URL.Query().Get("item") // Get the requested item and store it as a string
	if price, ok := db.dbm[item]; ok {  // Check if it exists in the database
		fmt.Fprintf(w, "%s\n", price) // Print price if it exits
	} else { // If it doesn't exist
		w.WriteHeader(http.StatusNotFound)         // 404
		fmt.Fprintf(w, "no such item: %q\n", item) // No item in database
	}
}

func (db *database) create(w http.ResponseWriter, req *http.Request) { // Place a new item and price in the database
	db.rwm.Lock()                                 // Write lock
	defer db.rwm.Unlock()                         // Unlock at the end
	item := req.URL.Query().Get("item")           // Collect the requested item
	price := req.URL.Query().Get("price")         // Collect the requested price
	tempPrice, e := strconv.ParseFloat(price, 32) // Parse the price
	if e != nil {                                 // Check if the price can become a float
		w.WriteHeader(http.StatusNotFound) // Price can't become a float
		fmt.Fprintf(w, "Price not supported: %s\n", price)
	} else if _, ok := db.dbm[item]; ok { // Check if it already exists in the database
		w.WriteHeader(http.StatusNotAcceptable) // Tell them to use update
		fmt.Fprintf(w, "%s is already in the database!\nPlease use update instead.\n", item)
	} else { // If it's not in the database
		db.dbm[item] = dollars(tempPrice) // Place in the database
		fmt.Fprintf(w, "%s added in the database with price %s\n", item, dollars(tempPrice))
	}
}

func (db *database) update(w http.ResponseWriter, req *http.Request) { // Update an item's price in the database
	db.rwm.Lock()                                 // Write lock
	defer db.rwm.Unlock()                         // Unlock at the end
	item := req.URL.Query().Get("item")           // Collect the requested item
	price := req.URL.Query().Get("price")         // Collect the requested price
	tempPrice, e := strconv.ParseFloat(price, 32) // Turn the price into a float
	if e != nil {                                 // Check if the price is supported
		w.WriteHeader(http.StatusNotFound) // The price is not a float
		fmt.Fprintf(w, "Price not supported: %s\n", price)
	} else if _, ok := db.dbm[item]; ok { // Check if it exists in the database
		db.dbm[item] = dollars(tempPrice) // Update price in the database
		fmt.Fprintf(w, "%s is now %s\n", item, dollars(tempPrice))
	} else { // If it's not in the database, tell them it doesn't exist
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db *database) delete(w http.ResponseWriter, req *http.Request) { // Delete an item from the database
	db.rwm.Lock()                       // Write lock
	defer db.rwm.Unlock()               // Unlock at the end
	item := req.URL.Query().Get("item") // Get the requested item
	if _, ok := db.dbm[item]; ok {      // Check if it's in the database
		delete(db.dbm, item) // Remove the item from the database
		fmt.Fprintf(w, "%s was removed from the database\n", item)
	} else { // If it's not in the database, tell them it doesn't exist
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}

}
