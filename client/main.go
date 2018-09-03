package client

import (
	consul "github.com/hashicorp/consul/api"
	"github.com/pkg/errors"
)

// Client is a struct to wrap a consul connection and provide helper methods
type Client struct {
	consul *consul.Client
}

// NewClient creates a configured instance of the Client struct
func NewClient(addr string) (*Client, error) {
	config := consul.DefaultConfig()
	config.Address = addr
	c, err := consul.NewClient(config)
	if err != nil {
		return nil, errors.Wrap(err, "Couldn't create consul client")
	}
	return &Client{consul: c}, nil
}

// Service returns a consul service entry
func (c *Client) Service(service, tag string) ([]*consul.ServiceEntry, *consul.QueryMeta, error) {
	addrs, meta, err := c.consul.Health().Service(service, tag, true, nil)
	if len(addrs) == 0 && err == nil {
		return nil, nil, errors.Wrapf(err, "service ( %s ) was not found", service)
	}
	if err != nil {
		return nil, nil, errors.Wrap(err, "Error talking to consul")
	}
	return addrs, meta, nil
}
