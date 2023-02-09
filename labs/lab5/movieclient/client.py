import logging

import grpc
import os, sys
sys.path.insert(0, os.path.abspath(os.path.join(os.path.dirname(__file__), '../movieapi')))

import movieapi_pb2
import movieapi_pb2_grpc

address = "localhost:50051"
defaultTitle = "Pulp fiction"

def main():
    # Set up a connection to the server
    with grpc.insecure_channel(address) as channel:
        c = movieapi_pb2_grpc.MovieInfoStub(channel)
        title = defaultTitle
        r = c.GetMovieInfo(movieapi_pb2.MovieRequest(title=title))
        logging.log(20, "Movie Info for %s %d %s %v", title, r.year, r.director, r.cast)
        print("Movie Info for", title, r.year, r.director, ", ".join(r.cast))

        r = c.SetMovieInfo(movieapi_pb2.MovieData(title="test", director="A B", year=1234, cast=["^a b^", "^c d^", "^e f^"]))
        logging.log(20, "Movie Set status %d", r.code)
        print("Movie Set status", r.code)

        r = c.SetMovieInfo(movieapi_pb2.MovieData(title="Black Panther: Wakanda Forever", year=2022, director="Ryan Coogler", cast=["Letitia Wright", "Tenoch Huerta",
		"Angela Bassett", "Michael B. Jordan", "Dominique Thorne", "Lupita Nyong'o", "Mabel Cadena", "Danai Gurira"]))
        logging.log(20, "Movie Set status %d", r.code)
        print("Movie Set status", r.code)

if __name__ == "__main__":
    main()