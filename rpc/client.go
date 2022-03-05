package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

type Params struct {
	Weight, Height int
}

func main() {
	conn, err := jsonrpc.Dial("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}

	ret := 0
	p := Params{10, 20}
	err2 := conn.Call("Rect.Area", p, &ret)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println("面积：", ret)

	err3 := conn.Call("Rect.Perimeter", p, &ret)
	if err3 != nil {
		log.Fatal(err3)
	}
	fmt.Println("周长：", ret)
}
