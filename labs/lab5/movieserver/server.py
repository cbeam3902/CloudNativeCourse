from concurrent import futures
import logging

import grpc
import os, sys
sys.path.insert(0, os.path.abspath(os.path.join(os.path.dirname(__file__), '../movieapi')))

import movieapi_pb2
import movieapi_pb2_grpc

port = ":50051"

moviedb = {"Pulp fiction": ["1994", "Quentin Tarantino", "John Travolta,Samuel Jackson,Uma Thurman,Bruce Willis"]}

class MovieDatabase(movieapi_pb2_grpc.MovieInfoServicer):
    def GetMovieInfo(self, request, context):
        year = None
        director = None
        cast = []
        title = request.title
        print("Received:", title)
        ok = title in moviedb
        if not ok:
            movieinfo = movieapi_pb2.MovieReply(year=year, director=director, cast=cast)
            return movieinfo
        else:
            v = moviedb[title]
            year = int(v[0])
            director = v[1]
            cast = v[2].split(",")
            movieinfo = movieapi_pb2.MovieReply(year=year, director=director, cast=cast)
            return movieinfo

    def SetMovieInfo(self, request, context):
        moviedb[request.title] = [str(request.year), request.director, ",".join(request.cast)]
        status = movieapi_pb2.Status(code=0)
        print("moviedb: ", moviedb)
        return status

    
def main():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=1))
    movieapi_pb2_grpc.add_MovieInfoServicer_to_server(MovieDatabase(), server)
    server.add_insecure_port('[::]'+port)
    server.start()
    print("Server started, listening on " + port)
    server.wait_for_termination()

if __name__ == "__main__":
    main()