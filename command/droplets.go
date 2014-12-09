package command

import (
	"fmt"
	"github.com/masayukioguni/go-digitalocean/digitalocean"
)

type DropletsCommand struct {
	Client *digitalocean.Client
}

func (c *DropletsCommand) Help() string {
	return ""
}

func (c *DropletsCommand) Run(args []string) int {

	droplets, hr, err := c.Client.DropletsService.List()

	if err != nil {
		fmt.Printf("error: %v\n\n", err)
		return 1
	}

	if hr.StatusCode != 200 {
		fmt.Printf("http response error: %+v %+v \n\n", hr, err)
		return 1
	}

	for _, droplet := range droplets {
		fmt.Printf("%s (ip: %s, status: %s, region :%s, id: %d)\n",
			droplet.Name, GetV4IPAddress(droplet), droplet.Status, droplet.Region.Slug, droplet.ID)
	}
	return 0
}

func (c *DropletsCommand) Synopsis() string {
	return fmt.Sprintf("Retrieve a droplet list")
}
