FROM golang:1.18-alpine AS build

WORKDIR /src/
COPY webserver.go /src/
RUN go mod init src
RUN go mod tidy
RUN CGO_ENABLED=0 go build -o /bin/webserver webserver.go

FROM scratch
COPY --from=build /bin/webserver /bin/webserver
ENTRYPOINT [ "/bin/webserver"]