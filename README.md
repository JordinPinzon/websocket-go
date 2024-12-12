# WebSocket Go Server

This is a simple WebSocket server implemented in Go. The server listens on port 8080 and allows bidirectional real-time communication with WebSocket clients.

## Requirements

Before running the application, make sure you have the following:

- Go 1.20 or higher
- Docker (if you want to run the app in a container)

## Installation and Running

### Option 1: Run Locally

1. Clone this repository to your local machine:

    ```bash
    git clone https://github.com/JordinPinzon/websocket-go.git
    cd websocket-go
    ```

2. Install dependencies:

    ```bash
    go mod tidy
    ```

3. Build the application:

    ```bash
    go build -o websocket-server .
    ```

4. Run the application:

    ```bash
    ./websocket-server
    ```

    The server will be available at `ws://localhost:8080/ws`.

5. To connect to the server, you can use a WebSocket client like Postman or any library that supports WebSockets.

    **Using Postman**:

    - Open Postman and select the **WebSocket** method.
    - In the URL field, enter `ws://localhost:8080/ws`.
    - Click **Connect** and you will be able to send and receive messages.

### Option 2: Run with Docker

If you prefer to run the app inside a Docker container, follow these steps:

1. Build the Docker image:

    ```bash
    docker build -t websocket-go .
    ```

2. Run the container:

    ```bash
    docker run -p 8080:8080 websocket-go
    ```

    The server will be available at `ws://localhost:8080/ws`.

3. Just like when running locally, you can connect to the WebSocket server using Postman or another WebSocket client.

## Important Files

- `main.go`: The main file of the Go WebSocket server.
- `go.mod`: Go module file to manage dependencies.
- `go.sum`: Go checksum file for dependencies.
- `Dockerfile`: Contains the instructions to build the Docker container.
- `README.md`: This file.

## Contributing

If you wish to contribute to this project, please fork the repository, make your changes, and then open a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
