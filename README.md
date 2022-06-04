# imageAPI

HTTP API for uploading, optimizing, and serving images.

## Prerequisites

- [libvips](https://github.com/libvips/libvips) 8.3+ (8.8+ recommended)
- C compatible compiler such as gcc 4.6+ or clang 3.0+
- Go 1.7+
- Docker

## Installation

- Install [Docker](https://docs.docker.com/desktop/#download-and-install), then run follow command in terminal 
```bash
docker run -d --hostname my-rabbit --name some-rabbit -p 15672:15672 -p 5672:5672 rabbitmq:3-management

```
Now you can login to RabbitMQ in http://localhost:15672 use the username `guest` and password `guest` to login.


- In project run following command to start server in `:8080` port
```bash
go run cmd/main.go | go run cmd/consumer.go
```
- If you have problems building the vips library like `invalid flag in pkg-config --cflags: -Xpreprocessor`, try run
```bash
CGO_CFLAGS_ALLOW=-Xpreprocessor go vet ./...
```
Then go back to the previous step

## Production

- Open index.html

You can upload photos that will be processed by [bimg](https://github.com/h2non/bimg) package and will be available at `:8080/{photoName}`

- Use query parameters `?quality=100/75/50/25` after the photo name to get an optimized image with 100% 75%, 50% and 25% quality respectively. 

