FROM golang:1.17.1 as build
ENV DB_CONN=sqlite
RUN apt-get update && apt-get install sqlite3 -y
WORKDIR /src
COPY go.sum go.mod ./
RUN go mod download
COPY . .
RUN go test -v -cover ./...
RUN go build -a -ldflags "-linkmode external -extldflags '-static' -s -w" -o /bin/app .

FROM alpine
ENV DB_CONN=sqlite
ENV ENV=local
WORKDIR /app
COPY --from=build /bin/app .
CMD ["./app"]