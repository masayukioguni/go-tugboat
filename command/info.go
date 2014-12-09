package command

import (
	"errors"
	"flag"
	"fmt"
	"github.com/masayukioguni/go-digitalocean/digitalocean"
	"strings"
)

type InfoCommand struct {
	Client *digitalocean.Client
}

type InfoFlags struct {
	ID int
}

func (c *InfoCommand) Help() string {
	helpText := `
Usage: go-tugboat info [Require Options]

  Show a droplet's information

Require Options:
  -id=[droplet id] 
`
	return strings.TrimSpace(helpText)
}

func (c *InfoCommand) parse(args []string) (*InfoFlags, error) {
	flags := new(InfoFlags)

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

func (d *InfoCommand) GetV4IPAddress(droplet *digitalocean.Droplet) []string {
	var ipv4s []string

	for _, v4 := range droplet.Networks.V4s {
		ipv4s = append(ipv4s, v4.IPAddress)
	}

	return ipv4s
}

func (c *InfoCommand) Run(args []string) int {

	flags, err := c.parse(args)

	if err != nil {
		fmt.Printf("comannd parse error : %v\n\n", err)
		return 0
	}

	droplet, hr, err := c.Client.DropletsService.Get(flags.ID)

	if err != nil {
		fmt.Printf("error: %+v\n\n", err)
		return 1
	}

	if hr.StatusCode != 200 {
		fmt.Printf("http response error: %+v\n\n", hr)
		return 1
	}

	fmt.Printf("%s (ip: %s, status: %s, region :%s, id: %d, image id:%d size:%s)\n",
		droplet.Name, c.GetV4IPAddress(droplet), droplet.Status, droplet.Region.Slug, droplet.ID, droplet.Image.ID, droplet.Size)

	return 0
}

func (c *InfoCommand) Synopsis() string {
	return fmt.Sprintf("Show a droplet's information")
}
