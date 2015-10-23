# Nats.io websocket interface

This is just a proof of concept in order to have a websocket interface for nats.

## Run the examples

On three different terminals run:


Server is providing the server websocket interface and it's the only one really connected to nats.
```
$ go run server.go
```

Client is listening through websockets on the server.
```
$ go run client.go
```

Publisher sends a post message to the server with the subject and the body to be sent thorugh nats.
```
$ go run publisher.go
```

You can additionally run another example if you have access to nats itself, just publish a message with the subject the client is listening to, and the message should be printed on client output.


