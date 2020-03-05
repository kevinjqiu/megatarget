package cmd

import (
	"github.com/kevinjqiu/megatarget/mt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"net"
)

const serverBind = ":9999"

func newServeCommand() *cobra.Command {
	return &cobra.Command{
		Use: "serve",
		Run: func(cmd *cobra.Command, args []string) {
			lis, err := net.Listen("tcp", serverBind)
			if err != nil {
				logrus.Fatal(err)
			}
			server := grpc.NewServer()
			mt.RegisterTargetsServer(server, mt.NewServer())
			logrus.Info("Serving mt on ", serverBind)
			if err := server.Serve(lis); err != nil {
				logrus.Fatal(err)
			}
		},
	}
}
