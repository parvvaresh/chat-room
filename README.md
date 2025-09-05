# TCP Chat Application

A simple TCP chat application built in Go. The project includes a server and client programs, which allow multiple users to connect and chat in real-time. The project also provides Docker and Docker Compose setup for easy deployment and testing.

---

## Features

* TCP-based chat server and client
* Support for multiple clients simultaneously
* Dockerized server and clients for easy deployment
* Can be extended to WebSocket support or custom features

---

## Project Structure

```
tcp-chat/
├── server.go          # Server source code
├── client.go          # Client source code
├── Dockerfile         # Dockerfile to build server and client images
├── docker-compose.yml # Compose file to run server and multiple clients
├── go.mod             # Go module file
├── go.sum             # Checksums for dependencies
└── README.md          # Project documentation
```

---

## Prerequisites

* Go 1.21 or later
* Docker (for containerized setup)
* Docker Compose (for running multiple services)

---

## Running Locally

### 1. Build and run the server manually

```bash
# Build server binary
go build -o chat-server ./server.go

# Run server
./chat-server
```

### 2. Build and run the client manually

```bash
# Build client binary
go build -o chat-client ./client.go

# Run client (can run multiple instances)
./chat-client
```

---

## Running with Docker

### 1. Build Docker images

```bash
docker-compose build
```

### 2. Run server and clients

```bash
docker-compose up
```

* The server will start listening on port 8080.
* Clients will connect automatically.
* You can open multiple terminal sessions to interact with different clients.

### 3. Stop containers

```bash
docker-compose down
```

---

## CI/CD with GitHub Actions

* The project includes a GitHub Actions workflow to:

  * Build Go binaries
  * Run tests
  * Build Docker images
  * Optional: push Docker images and deploy

---

## Extending the Project

* Add username support for clients
* Upgrade to WebSocket for browser-based chat
* Add logging or message persistence

---

## License

This project is licensed under the MIT License.


