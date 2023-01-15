package main

import "fmt"

type Client struct {
}

func (c *Client) startNavigation(transport Transport) {
	fmt.Println("Client 가 네비게이션을 시작하고 있다")
	transport.navigateToDestination()
}
