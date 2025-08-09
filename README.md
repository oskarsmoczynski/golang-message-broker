## golang-message-broker

An educational, minimal message broker written in Go, inspired by RabbitMQ and Kafka. It accepts messages from producers, stores them in in-memory queues, and delivers them to consumers with support for routing modes (direct, fan-out), acknowledgments, retries, and the foundation for durability.

### Goals
- Accept messages from producers (publishers)
- Store messages in queues (initially in-memory)
- Deliver to consumers (subscribers)
- Support routing modes: direct and fan-out
- Provide acknowledgments (ack), retry/redelivery, and durability hooks

### High-level flow
1. Producer sends a message to the broker (e.g., TCP/HTTP/gRPC)
2. Broker stores the message in the appropriate queue
3. Consumer subscribes to a queue/topic
4. Broker delivers messages according to the routing mode
5. Consumer sends an ACK → broker completes/removes the message
6. If no ACK within the timeout → message is returned to the queue (retry/redelivery)

### Project structure
```
golang-message-broker/
  cmd/
    broker/
      main.go                # Entry point (WIP)
  internal/
    broker/
      ack/
        ack_manager.go       # Ack tracking (timeouts, redelivery hooks)
      config/
        config.go            # Broker configuration (ports, timeouts)
      dispatcher/
        dispatcher.go        # Delivery to subscribers according to routing
      models/
        message.go           # Core domain types (Message)
      queue/
        manager.go           # Queue manager (create/get/list)
        queue.go             # Queue interface
      routing/
        direct.go            # Direct routing mode
        fanout.go            # Fan-out routing mode
      server/
        server.go            # Connection handler abstraction
        http/
          http_server.go     # HTTP server stub
        tcp/
          tcp_server.go      # TCP server stub
      storage/
        memory/
          memory_store.go    # In-memory storage implementation (WIP)
  README.md
  go.mod
  .gitignore
  LICENSE
```

### Delivery modes (routing)
- **Direct**: send the message to a single target queue/consumer.
- **Fan-out**: replicate and deliver the message to all subscribers of the target topic.

### Acknowledgment, retry, durability
- **Ack**: Consumer confirms successful processing; broker marks the message as done.
- **Retry**: On timeout or negative ack, the message is returned to the queue for redelivery.
- **Durability**: This scaffold focuses on in-memory. Hooks and structure allow you to add persistent storage later (e.g., file or log-based store) without changing public APIs.

### Getting started
1. Ensure Go 1.22+ is installed.
2. Initialize the module path in `go.mod` if needed (replace the placeholder).
3. Build and run the broker skeleton:

```bash
go build ./cmd/broker
./broker
```

On Windows PowerShell:

```powershell
go build ./cmd/broker
./broker.exe
```
