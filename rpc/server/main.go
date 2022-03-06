package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Params struct {
	Weight, Height int
}

type Rect struct{}

// RPC服务端方法，计算矩阵的面积
// 方法名称必须大写，第二个参数是返回参数，必须是指针类型
// 函数返回必须是error
func (r *Rect) Area(p Params, ret *int) error {
	*ret = p.Weight * p.Height
	return nil
}

// 计算矩阵的周长
func (r *Rect) Perimeter(p Params, ret *int) error {
	*ret = 2 * (p.Weight + p.Height)
	return nil
}

func main() {
	// 注册一个rpc服务
	rect := new(Rect)
	rpc.Register(rect)

	// 服务处理绑定到http协议上
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Panicln(err)
	}

	// 采用jsonrpc进行通信，这样可以支持跨语言调用
	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Panicln(err)
		}

		go func(conn net.Conn) {
			jsonrpc.ServeConn(conn)
		}(conn)

	}
}
