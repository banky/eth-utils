// Package anvil provides a Go client for interacting
// with the Anvil Ethereum development environment. Anvil
// reference: https://getfoundry.sh/anvil/reference/
package anvil

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

type Anvil struct {
	url       string
	wsUrl     string
	cmd       *exec.Cmd
	rpcClient *rpc.Client
	ethClient *ethclient.Client
}

// New creates a new Anvil instance with default configuration.
func New() (Anvil, error) {
	c := NewConfig()
	return NewWithConfig(c)
}

// NewWithConfig creates a new Anvil instance with custom configuration.
func NewWithConfig(config *Config) (Anvil, error) {
	args := getArgs(config)

	cmd := exec.Command("anvil", args...)
	if config.showLogs {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	err := cmd.Start()
	if err != nil {
		return Anvil{}, fmt.Errorf("Failed to start Anvil: %v", err)
	}

	var port uint = 8545 // Default port
	if config.port != 0 {
		port = config.port
	}

	url := fmt.Sprintf("http://localhost:%d", port)
	wsUrl := fmt.Sprintf("ws://localhost:%d", port)

	rpcUrl := url
	if strings.HasPrefix(config.forkURL, "ws") {
		rpcUrl = wsUrl
	}

	rpcClient, err := rpc.Dial(rpcUrl)
	if err != nil {
		return Anvil{}, fmt.Errorf("error connecting to RPC: %v", err)
	}

	ethClient := ethclient.NewClient(rpcClient)

	// Wait until client is ready
	startedSuccessfully := false
	for range 1000 {
		_, err := ethClient.BlockNumber(
			context.Background(),
		)
		if err != nil {
			time.Sleep(100 * time.Millisecond)
			continue
		}

		startedSuccessfully = true
		break
	}

	if !startedSuccessfully {
		return Anvil{}, fmt.Errorf("Anvil did not start in time")
	}

	return Anvil{
		url:       url,
		wsUrl:     wsUrl,
		cmd:       cmd,
		rpcClient: rpcClient,
		ethClient: ethClient,
	}, nil
}

// EthClient returns an instance of *ethclient.Client
// connected to the running Anvil instance
func (a *Anvil) EthClient() *ethclient.Client {
	return a.ethClient
}

// HttpUrl returns the HTTP URL of the running Anvil instance
func (a *Anvil) HttpUrl() string {
	return a.url
}

// WsUrl returns the WebSocket URL of the running Anvil instance
func (a *Anvil) WsUrl() string {
	return a.wsUrl
}

// Close closes the running Anvil instance
func (a *Anvil) Close() error {
	err := a.cmd.Process.Signal(syscall.SIGTERM)
	if err != nil {
		return fmt.Errorf("error closing Anvil: %v", err)
	}

	a.rpcClient.Close()
	a.ethClient.Close()

	return nil
}
