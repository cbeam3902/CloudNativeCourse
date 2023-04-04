// Package main implements a server for movieinfo service.
package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"

	"github.com/cbeam3902/CloudNativeCourse/labs/lab5/movieapi"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement movieapi.MovieInfoServer
type server struct {
	movieapi.UnimplementedMovieInfoServer
}

// Map representing a database
var moviedb = map[string][]string{"Pulp fiction": []string{"1994", "Quentin Tarantino", "John Travolta,Samuel Jackson,Uma Thurman,Bruce Willis"}}

// GetMovieInfo implements movieapi.MovieInfoServer
func (s *server) GetMovieInfo(ctx context.Context, in *movieapi.MovieRequest) (*movieapi.MovieReply, error) {
	title := in.GetTitle()
	log.Printf("Received: %v", title)
	reply := &movieapi.MovieReply{}
	if val, ok := moviedb[title]; !ok { // Title not present in database
		return reply, errors.New(fmt.Sprintf("%s does not exist!\n", title))
	} else {
		if year, err := strconv.Atoi(val[0]); err != nil {
			reply.Year = -1
		} else {
			reply.Year = int32(year)
		}
		reply.Director = val[1]
		cast := strings.Split(val[2], ",")
		reply.Cast = append(reply.Cast, cast...)

	}

	return reply, nil

}

func (s *server) SetMovieInfo(ctx context.Context, in *movieapi.MovieData) (*movieapi.Status, error) {
	title := in.GetTitle() // Collect all of the variables
	year := in.GetYear()
	director := in.GetDirector()
	cast := in.GetCast()

	log.Printf("Received: %v", title)   // Print to log
	status := &movieapi.Status{}        // Declare Status
	status.Code = 0                     // Set to a default number
	year_str := strconv.Itoa(int(year)) // Convert string to int

	// log.Println("Hi")
	moviedb[title] = []string{year_str, director, strings.Join(cast, ",")} // Store in map

	log.Println("Moviedb: ", moviedb) // Print moviedb

	return status, nil

}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	movieapi.RegisterMovieInfoServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
