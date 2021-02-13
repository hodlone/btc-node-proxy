package cmd

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/NodeHodl/btc-node-proxy/cmd/blocksync"
	"github.com/NodeHodl/btc-node-proxy/cmd/grpcserver"
	"github.com/NodeHodl/btc-node-proxy/cmd/txsync.go"
	"github.com/NodeHodl/btc-node-proxy/cmd/zmq"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "btc-node-proxy",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: HandleOpts,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().Bool("--no-grpc", false, "Disables grpcserver")
	rootCmd.PersistentFlags().Bool("--no-zmq", false, "Disables zmq listener")
}

// HandleOpts ...
func HandleOpts(cmd *cobra.Command, args []string) {
	var enabled bool
	var err error
	wg := new(sync.WaitGroup)

	// grpcserver ...
	if enabled, err = cmd.Flags().GetBool("--no-grpc"); err != nil {
		log.Panic(err)
	}

	if enabled {
		wg.Add(1)
		go grpcserver.Start()
	}

	// zmq ...
	if enabled, err = cmd.Flags().GetBool("--no-zmq"); err != nil {
		log.Panic(err)
	}

	if enabled {
		wg.Add(1)
		go zmq.Start()
	}

	// blocksync ...
	wg.Add(1)
	go blocksync.Start(3)

	// txsync ...
	wg.Add(1)
	go txsync.Start()

	wg.Wait()
}
