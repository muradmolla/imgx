package main

import (
	"fmt"

	"github.com/muradmolla/imagetingx"
)

func execute() error {
	fmt.Printf("Processing image with values; gamma: %f, input: %s, output: %s\n", setup.gamma, setup.input, setup.output)
	img, err := imagetingx.New(setup.input)
	if err != nil {
		return fmt.Errorf("cannot init image. error :%s", err)
	}

	if setup.gamma != 1 {
		img.Gamma(setup.gamma)
	}

	img.Save(setup.output)

	return nil
}
