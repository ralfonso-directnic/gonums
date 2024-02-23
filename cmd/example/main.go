package main

import(
	"fmt"
	"github.com/ralfonso-directnic/gonums"
	"os"
	"path/filepath"
)

func main() {

	fmt.Println("Enum Generator - Building Enums from Config")
	config, err := ParseInput()
	if err != nil {
		return
	}
	g := gonums.New(config)
	err = g.Generate()
	if err != nil {
		fmt.Println(err)
	}
}

func ParseInput() (gonums.Config, error) {

	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	cfgPath := filepath.Join(path, "./input.json")
	cfg, err := gonums.ReadConfig(cfgPath)
	if err != nil {
		fmt.Println("Enum Build Error", err)
		return gonums.Config{}, err
	}
	cfg.OutputPath = filepath.Join(path, "./")
	fmt.Println("Enum Generator - Enums Loaded in", cfg.OutputPath)
	return cfg, err
}
