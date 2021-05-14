package demo

import (
	"fmt"
	"net"
	"net/rpc"

	"github.com/hashicorp/yamux"
)

type Client struct {
	client *rpc.Client
	broker *MuxBroker
}

func NewClient(addr string) (*Client, error) {
	conn, err := net.Dial("unix", addr)
	if err != nil {
		return nil, err
	}
	if tcpConn, ok := conn.(*net.TCPConn); ok {
		// Make sure to set keep alive so that the connection doesn't die
		tcpConn.SetKeepAlive(true)
	}

	// Create the yamux client so we can multiplex
	mux, err := yamux.Client(conn, nil)
	if err != nil {
		conn.Close()
		return nil, err
	}

	// Connect to the control stream.
	control, err := mux.Open()
	if err != nil {
		mux.Close()
		return nil, err
	}

	// Create the broker and start it up
	broker := newMuxBroker(mux)
	go broker.Run()

	// Build the client using our broker and control channel.
	return &Client{
		broker: broker,
		client: rpc.NewClient(control),
	}, nil

	// client, err := rpc.Dial("unix", addr)
	// if err != nil {
	// 	return nil, err
	// }

	// return &Client{
	// 	client: client,
	// 	inter:  &Inter{},
	// }, nil
}

func (c *Client) GetChannel(channelID string) (int, error) {
	var out *int
	err := c.client.Call("MyStruct.MyTest", channelID, &out)
	if err != nil {
		return 0, err
	}
	return *out, nil
}

func (c *Client) Register() error {
	muxID := c.broker.NextId()
	go c.broker.AcceptAndServe(muxID, &Inter{})

	var out *error
	err := c.client.Call("MyStruct.Register", muxID, &out)
	if err != nil {
		return err
	}
	return *out
}

func (c *Client) Callback(s string) error {
	go func() {
		var out *error
		err := c.client.Call("MyStruct.Callback", s, &out)
		if err != nil {
			fmt.Println(err)
		}
	}()
	// return *out
	return nil
}
