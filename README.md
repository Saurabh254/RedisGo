# RedisGo

A lightweight, educational Redis server and client implementation written in Go.

> [!NOTE]
> Built for educational purposes to explore network programming, concurrent TCP connection handling, and the Redis Serialization Protocol (RESP).

---

## Features

- **RESP Protocol Parser**: Decodes and serializes Redis commands and data types (simple strings, errors, integers, bulk strings, and arrays) compatible with RESP specifications.
- **Concurrent TCP Server**: Asynchronous peer connection handling using Go channels and goroutines.
- **In-Memory Key-Value Store**: Thread-safe storage supporting core key-value operations (`GET`, `SET`).
- **Go Client Library**: Built-in native client implementation located in `client/` for connecting and executing commands programmatically.

---

## Project Structure

```
.
├── Makefile            # Build and run commands
├── client/             # Go client implementation
├── keyval.go           # In-memory key-value store
├── main.go             # Server initialization & event loop
├── peer.go             # Peer connection management & command dispatcher
├── proto.go            # RESP protocol parser and serializer
└── proto_test.go       # Protocol test suite
```

---

## Getting Started

### Prerequisites

- Go 1.22+

### Build & Run Server

```bash
# Clone the repository
git clone https://github.com/Saurabh254/RedisGo.git
cd RedisGo

# Build the server binary
make build

# Run the server (listens on :5001 by default)
make run
```

### Running Tests

Run protocol parser unit tests:

```bash
go test -v ./...
```

### Using the Client

You can interact with RedisGo using the included client library:

```go
package main

import (
    "context"
    "fmt"
    "goredis/client"
)

func main() {
    c, err := client.New("localhost:5001")
    if err != nil {
        panic(err)
    }
    defer c.Close()

    _ = c.Set(context.TODO(), "greeting", "hello world")
    val, _ := c.Get(context.TODO(), "greeting")
    fmt.Println("Value:", val)
}
```

---

## License

Educational / Personal Use.
