package main

import (
	"github.com/kevinjqiu/megatarget/cmd"
	"github.com/sirupsen/logrus"
)

func main() {
	rootCmd := cmd.NewRootCommand()
	if err := rootCmd.Execute(); err != nil {
		logrus.Fatal(err)
	}
}
