package main

import (
	"net/rpc"

	"github.com/mattermost/mattermost-server/v5/model"
)

type Client struct {
	client *rpc.Client
}

func NewClient(addr string) (*Client, error) {
	client, err := rpc.Dial("unix", addr)
	if err != nil {
		return nil, err
	}

	return &Client{
		client: client,
	}, nil
}

func (c *Client) GetChannel(channelID string) (*model.Channel, error) {
	var out *model.GetChannelRPCResponse
	err := c.client.Call("MyStruct.MyTest", channelID, &out)
	if err != nil {
		return nil, err
	}
	return out.Channel, nil
}
