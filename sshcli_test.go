package sshcli

import (
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSSHCli(t *testing.T) {
	cases := []struct {
		host     string
		port     string
		user     string
		password string

		workdir  string
		cmd      string
		result   interface{}
		expected string
	}{
		{"10.42.6.66", "22", "root", "zm123", "", "pwd", nil, "/root\n"},
		{"10.42.6.66", "22", "root", "zm123", "/home", "pwd", nil, "/home\n"},
		{"10.42.6.66", "22", "root", "zm123", "/wwcd", "pwd", struct{}{}, ""},
		{"10.42.6.66", "22", "root", "zm111", "/wwcd", "pwd", struct{}{}, ""},
	}

	for _, v := range cases {
		cli := NewSSHClient(net.JoinHostPort(v.host, v.port), v.user, v.password)
		actual, err := cli.CD(v.workdir).Run(v.cmd, 10*time.Second)
		if v.result != nil {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
			assert.Equal(t, v.expected, string(actual))
		}
	}
}
