package cmd

import (
	"context"
	"fmt"
	"github.com/kevinjqiu/megatarget/mt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

func callClient(f func(client *mt.Client) error) {
	conn, err := grpc.Dial(serverBind, grpc.WithInsecure())
	if err != nil {
		logrus.Fatal(err)
	}
	defer conn.Close()
	client := mt.NewClient(conn)

	err = f(client)
	if err != nil {
		logrus.Fatal(err)
	}
}

func newTargetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     "target",
		Aliases: []string{"t"},
	}

	cmd.AddCommand(&cobra.Command{
		Use: "ls",
		Run: func(cmd *cobra.Command, args []string) {
			callClient(func(client *mt.Client) error {
				resp, err := client.List(context.Background(), &mt.ListTargetParams{})
				if err != nil {
					return err
				}
				for _, t := range resp.Targets {
					fmt.Println(t)
				}
				return nil
			})
		},
	})

	cmd.AddCommand(&cobra.Command{
		Use: "new",
		Run: func(cmd *cobra.Command, args []string) {
			callClient(func(client *mt.Client) error {
				for _, addr := range args {
					req := mt.NewTargetParams{
						Addr: addr,
					}
					_, err := client.New(context.Background(), &req)
					if err != nil {
						logrus.Warn(err)
					}
					logrus.WithField("target", addr).Info("Target added")
				}
				return nil
			})
		},
	})
	return &cmd
}
