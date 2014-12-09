package command

import (
	"errors"
	"flag"
	"fmt"
	"github.com/masayukioguni/go-digitalocean/digitalocean"
	"strings"
)

type DestroyCommand struct {
	Client *digitalocean.Client
}

type destroyFlags struct {
	ID int
}

func (c *DestroyCommand) Help() string {
	helpText := `
Usage: go-tugboat destory [Require Options]

  delete a Droplet

Require Options:
  -id=[droplet id] 
`
	return strings.TrimSpace(helpText)
}

func (c *DestroyCommand) parse(args []string) (*destroyFlags, error) {
	flags := new(destroyFlags)

	cmdFlags := flag.NewFlagSet("build", flag.ContinueOnError)

	cmdFlags.IntVar(&flags.ID, "id", 0, "Droplet id")

	err := cmdFlags.Parse(args)

	if err != nil {
		return nil, err
	}

	if flags.ID == 0 {
		return nil, errors.New("invalid id")
	}

	return flags, nil
}

func (c *DestroyCommand) Run(args []string) int {

	flags, err := c.parse(args)

	if err != nil {
		fmt.Printf("comannd parse error : %v\n\n", err)
		return 0
	}

	hr, err := c.Client.DropletsService.Destroy(flags.ID)

	if err != nil {
		fmt.Printf("error: %+v\n\n", err)
		return 1
	}

	if hr.StatusCode != 204 {
		fmt.Printf("http response error: %+v\n\n", hr)
		return 1
	}

	fmt.Printf("Queuing destroy for %d ...done\n", flags.ID)

	return 0
}

func (c *DestroyCommand) Synopsis() string {
	return fmt.Sprintf("Destroy a droplet")
}
