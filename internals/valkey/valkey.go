package valkey

import (
	"os"
	"strconv"

	glide "github.com/valkey-io/valkey-glide/go/v2"
	"github.com/valkey-io/valkey-glide/go/v2/config"
)

var client *glide.Client

func GetValKeyClient() *glide.Client {
	if client != nil {
		return client
	}

	host, ok := os.LookupEnv("VALKEY_HOST")
	if !ok {
		panic("VALKEY_HOST not set")
	}

	port, ok := os.LookupEnv("VALKEY_PORT")

	if !ok {
		panic("VALKEY_PORT not set")
	}

	parsedPort, err := strconv.Atoi(port)
	if err != nil {
		panic("Error parsing VALKEY_PORT")
	}

	config := config.NewClientConfiguration().
		WithAddress(&config.NodeAddress{Host: host, Port: parsedPort})

	newClient, err := glide.NewClient(config)
	if err != nil {
		panic("Error connecting to valkey:")
	}

	client = newClient
	return newClient
}
