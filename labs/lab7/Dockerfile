FROM golang:1.18-alpine AS build

WORKDIR /src/
COPY movieapi_grpc.pb.go movieapi.pb.go server.go /src/
RUN go mod init src
RUN go mod tidy
RUN CGO_ENABLED=0 go build -o /bin/movieserver server.go

FROM scratch
COPY --from=build /bin/movieserver /bin/movieserver
ENTRYPOINT [ "/bin/movieserver"]