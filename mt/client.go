package mt

import "google.golang.org/grpc"

type Client struct {
	TargetsClient
}

func NewClient(cc *grpc.ClientConn) *Client {
	return &Client{NewTargetsClient(cc)}
}
