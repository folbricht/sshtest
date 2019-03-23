package sshtest_test

import (
	"fmt"

	"github.com/folbricht/sshtest"
	"golang.org/x/crypto/ssh"
)

func ExampleNewUnstartedServer() {
	hostKey := sshtest.KeyFromFile("testdata/ssh-host-key", "")
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
}
