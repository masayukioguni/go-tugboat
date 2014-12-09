package command

import (
	"fmt"
	"github.com/masayukioguni/go-digitalocean/digitalocean"
)

type RegionsCommand struct {
	Client *digitalocean.Client
}

func (c *RegionsCommand) Help() string {
	return ""
}

func (c *RegionsCommand) Run(args []string) int {

	regions, hr, err := c.Client.RegionsService.List()

	if err != nil {
		fmt.Printf("error: %v\n\n", err)
		return 1
	}

	if hr.StatusCode != 200 {
		fmt.Printf("http response error: %+v %+v \n\n", hr, err)
		return 1
	}

	for _, region := range regions {
		fmt.Printf("slug: %s name: %s\n", region.Slug, region.Name)
	}

	return 0
}

func (c *RegionsCommand) Synopsis() string {
	return fmt.Sprintf("Show regions")
}
