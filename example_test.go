package sshtest_test

import (
	"fmt"

	"github.com/folbricht/sshtest"
)

func Example() {
	hostKey := sshtest.KeyFromFile("testdata/ssh-host-key", "")
	server := sshtest.NewServer(hostKey)
	defer server.Close()

	fmt.Println("SSH server started on:", server.Endpoint)
}
