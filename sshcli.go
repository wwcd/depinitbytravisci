package sshcli

import (
	"fmt"
	"time"

	"golang.org/x/crypto/ssh"
)

// SSHClient simple ssh client
type SSHClient struct {
	Host     string // host:port
	User     string // Username
	Password string // Password

	workdir string // Workdir
}

// NewSSHClient make ssh cli instance
func NewSSHClient(Host, Username, Password string) *SSHClient {
	cli := new(SSHClient)
	cli.Host = Host
	cli.User = Username
	cli.Password = Password

	return cli
}

// CD change workdir
func (cli *SSHClient) CD(d string) *SSHClient {
	cli.workdir = d
	return cli
}

// Run excute cmd
func (cli *SSHClient) Run(cmd string, timeout time.Duration) ([]byte, error) {
	conf := &ssh.ClientConfig{
		User: cli.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(cli.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         timeout,
	}

	conn, err := ssh.Dial("tcp", cli.Host, conf)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	if cli.workdir != "" {
		cmd = fmt.Sprintf("cd %s && %s", cli.workdir, cmd)
	}

	return session.CombinedOutput(cmd)
}
