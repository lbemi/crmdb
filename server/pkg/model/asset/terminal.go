package asset

import "golang.org/x/crypto/ssh"

type SSHClient struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	IpAddress string `json:"ipaddress"`
	Port      int    `json:"port"`
	Session   *ssh.Session
	Client    *ssh.Client
	channel   ssh.Channel
}

type PtyRequestMsg struct {
	Term     string
	Columns  uint32
	Rows     uint32
	Width    uint32
	Height   uint32
	Modelist string
}

type TerminalWindow struct {
	Columns uint32 `json:"cols"`
	Rows    uint32 `json:"rows"`
	Width   uint32 `json:"width"`
	Height  uint32 `json:"height"`
}
