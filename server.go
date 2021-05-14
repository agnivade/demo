package demo

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/rpc"
	"os"
	"runtime"

	"github.com/hashicorp/yamux"
)

type Server struct {
	l      net.Listener
	m      *MyStruct
	broker *MuxBroker
}

func NewServer() (*Server, error) {
	suiteL, err := serverListener()
	if err != nil {
		return nil, err
	}

	fmt.Println("address----", suiteL.Addr().String())
	return &Server{
		l: suiteL,
	}, nil
}

// func (s *Server) Big() {
// 	s.m.impl.Get("lolery")
// }

func (s *Server) Start() {
	for {
		conn, err := s.l.Accept()
		if err != nil {
			if !errors.Is(err, net.ErrClosed) {
				fmt.Println("error accepting rpc conn", err)
			}
			return
		}

		// cli := rpc.NewClient(conn)
		// var out *int
		// err = cli.Call("MyStruct.Multiply", 5, &out)
		// if err != nil {
		// 	mlog.Error("hahahaha------", mlog.Err(err))
		// 	continue
		// }
		// mlog.Info("got this-----------", mlog.Int("out", *out))

		go s.ServeConn(conn)
	}
}

func (s *Server) ServeConn(conn net.Conn) {
	mux, err := yamux.Server(conn, nil)
	if err != nil {
		conn.Close()
		log.Printf("[ERR] plugin: error creating yamux server: %s", err)
		return
	}

	// Accept the control connection
	control, err := mux.Accept()
	if err != nil {
		mux.Close()
		if err != io.EOF {
			log.Printf("[ERR] plugin: error accepting control connection: %s", err)
		}
		return
	}

	// Create the broker and start it up
	broker := newMuxBroker(mux)
	go broker.Run()

	s.broker = broker
	// Use the control connection to build the dispenser and serve the
	// connection.
	rpcServer := rpc.NewServer()
	// s.m = &MyStruct{}
	rpcServer.Register(NewGlue(s))
	// server.RegisterName("Control", &controlServer{
	// 	server: s,
	// })
	// server.RegisterName("Dispenser", &dispenseServer{
	// 	broker:  broker,
	// 	plugins: s.Plugins,
	// })
	rpcServer.ServeConn(control)
}

func serverListener() (net.Listener, error) {
	if runtime.GOOS == "windows" {
		return serverListener_tcp()
	}

	return serverListener_unix()
}

func serverListener_tcp() (net.Listener, error) {
	listener, err := net.Listen("tcp", "localhost:")
	if err != nil {
		return nil, err
	}

	return listener, nil
}

func serverListener_unix() (net.Listener, error) {
	tf, err := ioutil.TempFile("", "suite")
	if err != nil {
		return nil, err
	}
	path := tf.Name()

	// Close the file and remove it because it has to not exist for
	// the domain socket.
	if err := tf.Close(); err != nil {
		return nil, err
	}
	if err := os.Remove(path); err != nil {
		return nil, err
	}

	l, err := net.Listen("unix", path)
	if err != nil {
		return nil, err
	}

	// Wrap the listener in rmListener so that the Unix domain socket file
	// is removed on close.
	return &rmListener{
		Listener: l,
		Path:     path,
	}, nil
}

// rmListener is an implementation of net.Listener that forwards most
// calls to the listener but also removes a file as part of the close. We
// use this to cleanup the unix domain socket on close.
type rmListener struct {
	net.Listener
	Path string
}

func (l *rmListener) Close() error {
	println("trying to close from here-----------")
	// Close the listener itself
	// l.Listener.Close()
	return nil
	// Remove the file
	// return os.Remove(l.Path)
}
