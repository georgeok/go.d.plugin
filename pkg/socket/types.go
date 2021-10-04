package socket

import (
	"crypto/tls"
	"time"
)

type Socket interface {
	Connect() error
	Disconnect() error
	Command(command string, process Processor) error
}

type Processor func([]byte) bool

type Config struct {
	Network Network
	Address string
	Timeout time.Duration
	TLSConf *tls.Config
}

type Network string

const (
	NetworkIP   Network = "ip"
	NetworkTCP  Network = "tcp"
	NetworkUDP  Network = "udp"
	NetworkUnix Network = "unix"
)
