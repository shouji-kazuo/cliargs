package main

import (
	"fmt"
	"os"

	"github.com/shouji-kazuo/cliargs/cliargs"
	cli "gopkg.in/urfave/cli.v2"
)

// テスト用
func main() {
	tApp := &cli.App{
		Name:      "test",
		Usage:     "test",
		ArgsUsage: " ",
		Version:   "v1.0.0",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "username",
				Aliases: []string{"u"},
				Usage:   "set dummy username",
			},
			&cli.StringFlag{
				Name:    "password",
				Aliases: []string{"p"},
				Usage:   "set dummy password",
			},
		},
		Action: func(aContext *cli.Context) error {
			args, err := cliargs.Wrap(aContext.Args(), cliargs.DefaultFuncWhenSingleHyphen)
			if err != nil {
				return err
			}
			fmt.Println("args = ", args)
			return nil
		},
	}

	if tError := tApp.Run(os.Args); tError != nil {
		fmt.Fprintln(os.Stderr, tError)
		os.Exit(1)
	}
}
