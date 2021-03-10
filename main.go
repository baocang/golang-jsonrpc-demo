package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// Args holds arguments to be passed to service Arith in RPC call
type Args struct {
	// Multiplier
	A int
	// Multiplicand
	B int
}

// Arith represents service Arith with method Multiply
type Arith int

// Result of RPC call is of this type
type Result int

// Multiply procedure is invoked by rpc and calls Arith.Multiply which stores product of args.A and args.B in result pointer
func (t *Arith) Multiply(args Args, result *Result) error {
	return Multiply(args, result)
}

// Multiply stores product of args.A and args.B in result pointer
func Multiply(args Args, result *Result) error {
	log.Printf("Multiplying %d with %d\n", args.A, args.B)
	*result = Result(args.A * args.B)
	return nil
}

type rpcRequest struct {
	r *http.Request
	w *http.ResponseWriter
}

func (r *rpcRequest) Read(p []byte) (n int, err error) {
	return r.r.Body.Read(p)
}

func (r *rpcRequest) Write(p []byte) (n int, err error) {
	return (*r.w).Write(p)
}

func (r *rpcRequest) Close() error {
	return r.r.Body.Close()
}

func serveHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	jsonrpc.ServeConn(&rpcRequest{req, &res})
}

func main() {
	auth := new(Arith)
	err := rpc.Register(auth)

	if err != nil {
		log.Fatalf("Format of service Arith isn't correct, %s", err)
	}

	rpc.HandleHTTP()

	l, e := net.Listen("tcp", ":8080")

	if e != nil {
		log.Fatalf("Couldn't start listening on port 8080, %s", e)
	}

	log.Println("Register RPC handler")

	http.HandleFunc("/rpc", serveHTTP)

	log.Println("Serving RPC")

	err = http.Serve(l, nil)

	if err != nil {
		log.Fatalf("Error serving: %s", err)
	}
}

// curl http://127.0.0.1:8080/rpc -d '{"id": 1, "method": "Arith.Multiply", "params": [{"a": 3, "b": 2}]}'
// curl http://127.0.0.1:8080/rpc -d '{"id": 1, "method": "Arith.Multiply", "params": [{"A": 3, "B": 2}]}'
