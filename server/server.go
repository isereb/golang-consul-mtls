package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/consul/connect"
	"net/http"
)

func main() {

	// Create a Consul API client
	client, _ := api.NewClient(api.DefaultConfig())

	// Create an instance representing this service. "my-service" is the
	// name of _this_ service. The service should be cleaned up via Close.
	err := client.Agent().ServiceRegister(&api.AgentServiceRegistration{
		Name: "my-server",
		Port: 8080,
		Connect: &api.AgentServiceConnect{
			Native:         true,
			SidecarService: nil,
		},
	})
	if err != nil {
		panic(err)
	}
	svc, _ := connect.NewService("my-server", client)
	defer svc.Close()

	// Creating an HTTP server that serves via Connect
	server := &http.Server{
		Addr:      ":8080",
		TLSConfig: svc.ServerTLSConfig(),
		// ... other standard fields
	}

	// Serve!
	done := make(chan bool)
	go func() {
		err := server.ListenAndServeTLS("", "")
		if err != nil {
			fmt.Println("Err: ", err)
		}
		svc.Close()
		done <- true
	}()
	<-done
}

func Healthcheck(c *gin.Context) {
	c.JSON(200, Response{
		Status:  200,
		Message: "Success!!!!@#$@",
	})
}

type Response struct {
	Status  int
	Message string
}
