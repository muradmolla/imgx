package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

type Setup struct {
	input  string
	output string
	gamma  float64
}

var setup *Setup = &Setup{
	input:  "",
	output: "",
	gamma:  1.0,
}

func main() {
	app := &cli.App{
		Name:  "imgx",
		Usage: "a simple image processing cli that using muradmolla/imagetingx module",
		Flags: []cli.Flag{
			&cli.Float64Flag{
				Name:    "gamma",
				Value:   1.0,
				Aliases: []string{"g"},
				Usage:   "Posive number for gamma multiplyer",
			},
		},
		Action: func(cCtx *cli.Context) error {
			setup.gamma = cCtx.Float64("gamma")
			if setup.gamma < 0 {
				return fmt.Errorf("gamma cannot be negative")
			}
			setup.input = cCtx.Args().Get(0)
			setup.output = cCtx.Args().Get(1)
			fmt.Println("executing")
			err := execute()
			if err != nil {
				return err
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
