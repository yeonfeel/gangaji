package cmd

import (
	"context"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"github.com/yeonfeel/gopkg/er"
	"github.com/yeonfeel/gopkg/grpcserver"
	"github.com/yeonfeel/gopkg/logger"

	"github.com/yeonfeel/gangaji/config"
	"github.com/yeonfeel/gangaji/pkg/pb/v1beta1/common"
	"github.com/yeonfeel/gangaji/pkg/server"
)

var log = logger.Get("cmd")
var cfgFile string

type rootCmd struct {
	conf *config.Config
}

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "bartender",
	Short: "run the bartender server",
	Run: func(cmd *cobra.Command, args []string) {
		conf := config.Load()

		c := &rootCmd{conf}

		c.run()
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Error(er.Error(err, "Can't execute the root command!"))
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(config.ConfigFileName) // name of config file
	viper.AddConfigPath("$HOME")               // adding home directory as first search path
	viper.AutomaticEnv()                       // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Info("Using config file:", viper.ConfigFileUsed())
	}
}

func (c *rootCmd) run() error {
	// run the remote gRPC Server
	grpcDone := make(chan struct{})
	httpDone := make(chan struct{})

	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		if r := recover(); r != nil {
			log.Error("panic occurred: %v\n%s", r, debug.Stack())
		}
		cancel()

		log.Info("wait for graceful stopped")
		<-grpcDone
		<-httpDone
		log.InfoF("stopped!")
	}()

	if err := runGateway(ctx, c.conf, grpcDone, httpDone); err != nil {
		return err
	}
	log.Info("wait for signal")

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP)

	select {
	// case <-done:
	case sigval := <-sigterm:
		log.InfoF("SIGNAL %s received, then shutting down...\n", sigval)
	}

	return nil
}

func runGateway(ctx context.Context, cfg *config.Config, grpcDone chan<- struct{}, httpDone chan<- struct{}) error {
	s := server.NewServer(cfg)

	return grpcserver.RunGateway(ctx, grpcDone, httpDone, cfg.Server.GRPCPort, cfg.Server.HTTPPort,
		cfg.Server.SSL, cfg.Server.SSLCertPath, cfg.Server.SSLKeyPath, cfg.Server.GracefulShutdownTimeout,
		func(grpcServer *grpc.Server) {
			common.RegisterPingTestServer(grpcServer, s)
		},
		map[grpcserver.RegisterServiceURLPattern]grpcserver.RegisterServiceFromEndpoint{
			"/common/": common.RegisterPingTestHandlerFromEndpoint,
		},
	)
}
