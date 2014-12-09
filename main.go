package main

import (
	"github.com/masayukioguni/go-digitalocean/digitalocean"
	"github.com/masayukioguni/go-tugboat/command"
	"github.com/masayukioguni/go-tugboat/config"
	"github.com/mitchellh/cli"
	"log"
	"os"
)

var applicationName = "go-tugboat"
var version string

func main() {
	c := cli.NewCLI(applicationName, version)
	path, _ := config.GetConfigPath()
	config, _ := config.LoadConfig(path)

	client, _ := digitalocean.NewClient(&digitalocean.Option{APIKey: config.Authentication.APIKey})
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"version": func() (cli.Command, error) {
			return &command.VersionCommand{
				AppName: applicationName,
				Version: version,
			}, nil
		},
		"sizes": func() (cli.Command, error) {
			return &command.SizesCommand{
				Client: client,
			}, nil
		},
		"regions": func() (cli.Command, error) {
			return &command.RegionsCommand{
				Client: client,
			}, nil
		},
		"images": func() (cli.Command, error) {
			return &command.ImagesCommand{
				Client: client,
			}, nil
		},
		"account": func() (cli.Command, error) {
			return &command.AccountCommand{
				Client: client,
			}, nil
		},
		"keys": func() (cli.Command, error) {
			return &command.SSHKeyCommand{
				Client: client,
			}, nil
		},
		"droplets": func() (cli.Command, error) {
			return &command.DropletsCommand{
				Client: client,
			}, nil
		},
		"create": func() (cli.Command, error) {
			return &command.CreateCommand{
				Client:        client,
				DefalutSize:   config.Defaluts.Size,
				DefalutImage:  config.Defaluts.Image,
				DefalutKey:    config.Defaluts.Key,
				DefalutRegion: config.Defaluts.Region,
			}, nil
		},
		"info": func() (cli.Command, error) {
			return &command.InfoCommand{
				Client: client,
			}, nil
		},
		"destroy": func() (cli.Command, error) {
			return &command.DestroyCommand{
				Client: client,
			}, nil
		},
		"authorize": func() (cli.Command, error) {
			return &command.AuthorizeCommand{}, nil
		},
		"ssh": func() (cli.Command, error) {
			return &command.SSHCommand{Client: client}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
