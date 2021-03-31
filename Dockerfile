FROM golang:1.16.2 as build

WORKDIR /golang-jsonrpc-demo
EXPOSE 8080
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOARM=6 go build -ldflags '-w -s' -o app

FROM scratch

WORKDIR /golang-jsonrpc-demo
COPY --from=build /golang-jsonrpc-demo/app ./app

EXPOSE 8080

CMD ["./app"]
