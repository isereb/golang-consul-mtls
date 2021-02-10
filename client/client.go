package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/consul/connect"
)

func main() {
	// Create a Consul API client
	client, _ := api.NewClient(api.DefaultConfig())

	// Create an instance representing this service. "my-service" is the
	// name of _this_ service. The service should be cleaned up via Close.
	svc, _ := connect.NewService("my-client", client)
	defer svc.Close()

	// Get an HTTP client
	httpClient := svc.HTTPClient()

	// Perform a request, then use the standard response
	resp, err := httpClient.Get("https://my-server.service.consul/user/mitchellh")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp.StatusCode)
}
