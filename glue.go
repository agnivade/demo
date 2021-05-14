package demo

import (
	"fmt"
	"net/rpc"
)

type Inter struct {
}

func (i *Inter) Get(s string, out *error) error {
	fmt.Println(s)
	return nil
}

type MyStruct struct {
	s  *Server
	cl *rpc.Client
}

func NewGlue(s *Server) *MyStruct {
	return &MyStruct{s: s}
}

func (m *MyStruct) MyTest(channelID string, out *int) error {
	fmt.Println(channelID)
	*out = 10
	return nil
}

func (m *MyStruct) Register(muxID uint32, out *error) error {
	fmt.Println("registering")
	conn, err := m.s.broker.Dial(muxID)
	if err != nil {
		return err
	}

	m.cl = rpc.NewClient(conn)
	// m.impl = in
	// in.Get("hello")
	// *out = 20
	return nil
}

func (m *MyStruct) Callback(s string, out *error) error {
	err := m.cl.Call("Inter.Get", s, nil)
	if err != nil {
		return err
	}
	return nil

}

func (m *MyStruct) Multiply(arg int, reply *int) error {
	*reply = arg * 5
	return nil
}
