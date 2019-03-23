# sshtest

sshtest is a set of utilities for testing SSH features in Go. It currently supports quick setup and start of SSH servers to test a client against. To minimize the need for error handling in tests, most functions panic on error.

## Examples

### Start a simple SSH server with host key

```go
hostKey := sshtest.KeyFromFile("ssh-host-key", "")
server := sshtest.NewServer(hostKey)
defer server.Close()
```

### Start an SSH server with host key and certificate

```go
hostKey := sshtest.KeyFromFile("ssh-host-key", "ssh-host-key-cert.pub")
server := sshtest.NewServer(hostKey)
defer server.Close()
```

### SSH server with custom server config and handler

```go
hostKey := sshtest.KeyFromFile("ssh-host-key", "")

server := sshtest.NewUnstartedServer()
server.Config = &ssh.ServerConfig{NoClientAuth: true}
server.Config.AddHostKey(hostKey)
server.Handler = func(ch ssh.Channel, in <-chan *ssh.Request) {
  defer ch.Close()

  // Read a request from the client
  req, ok := <-in
  if !ok {
    return
  }
  fmt.Printf("Received '%s' request from client", req.Type)

  // Reply with a string
  req.Reply(true, []byte("Hello client"))

  // Let the client know the command completed successfuly (status=0)
  sshtest.SendStatus(ch, 0)
}

server.Start()
defer server.Close()
```

## Links

- Go SSH library - [https://godoc.org/golang.org/x/crypto/ssh](https://godoc.org/golang.org/x/crypto/ssh)
- GoDoc for sshtest - [https://godoc.org/github.com/folbricht/sshtest](https://godoc.org/github.com/folbricht/sshtest)
