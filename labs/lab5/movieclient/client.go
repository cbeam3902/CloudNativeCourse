// Package main imlements a client for movieinfo service
package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/cbeam3902/CloudNativeCourse/labs/lab5/movieapi"

	"google.golang.org/grpc"
)

const (
	address      = "localhost:50051"
	defaultTitle = "Pulp fiction"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := movieapi.NewMovieInfoClient(conn)

	// Contact the server and print out its response.
	title := defaultTitle
	if len(os.Args) > 1 {
		title = os.Args[1]
	}
	// Timeout if server doesn't respond
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetMovieInfo(ctx, &movieapi.MovieRequest{Title: title})
	if err != nil {
		log.Fatalf("could not get movie info: %v", err)
	}
	log.Printf("Movie Info for %s %d %s %v", title, r.GetYear(), r.GetDirector(), r.GetCast())

	// Call SetMovieInfo with test data
	s, err := c.SetMovieInfo(ctx, &movieapi.MovieData{Title: "Test", Year: 1234, Director: "A B", Cast: []string{"^a b^", "^c d^", "^e f^"}})
	if err != nil {
		log.Fatalf("could not set movie data: %v", err)
	}
	log.Printf("Movie Set status %d", s.GetCode())

	// Call SetMovieInfo with an actual movie
	s, err = c.SetMovieInfo(ctx, &movieapi.MovieData{Title: "Black Panther: Wakanda Forever", Year: 2022, Director: "Ryan Coogler", Cast: []string{"Letitia Wright", "Tenoch Huerta",
		"Angela Bassett", "Michael B. Jordan", "Dominique Thorne", "Lupita Nyong'o", "Mabel Cadena", "Danai Gurira"}})
	if err != nil {
		log.Fatalf("could not set movie data: %v", err)
	}
	log.Printf("Movie Set status %d", s.GetCode())

	// Call GetMovieInfo to get back the actual movie
	r, err = c.GetMovieInfo(ctx, &movieapi.MovieRequest{Title: "Black Panther: Wakanda Forever"})
	if err != nil {
		log.Fatalf("could not get movie info: %v", err)
	}
	log.Printf("Movie Info for %s %d %s %v", "Black Panther: Wakanda Forever", r.GetYear(), r.GetDirector(), r.GetCast())

	// Call GetMovieInfo to get an error
	r, err = c.GetMovieInfo(ctx, &movieapi.MovieRequest{Title: "Blah"})
	if err != nil {
		log.Fatalf("could not get movie info: %v", err)
	}
	log.Printf("Movie Info for %s %d %s %v", "Blah", r.GetYear(), r.GetDirector(), r.GetCast())
}
