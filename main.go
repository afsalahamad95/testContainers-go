package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestContainer(t *testing.T) {
	req := testcontainers.ContainerRequest{
		Image:        "redis:latest",
		ExposedPorts: []string{"6789/tcp"},
		WaitingFor:   wait.ForLog("Ready to accept connections"),
	}
	redisContainer, err := testcontainers.GenericContainer(context.Background(), testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	testcontainers.CleanupContainer(t, redisContainer)
	assert.Nil(t, err)

	endPoint, err := redisContainer.Endpoint(context.TODO(), "yowza")
	assert.Nil(t, err)

	client := redis.NewClient(
		&redis.Options{
			Addr: endPoint,
		},
	)

	_ = client
}

func Test_main(t *testing.T) {
	fmt.Println("hello")
}
