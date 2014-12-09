package command

import (
	"fmt"
	"github.com/masayukioguni/go-digitalocean/digitalocean"
)

type SizesCommand struct {
	Client *digitalocean.Client
}

func (c *SizesCommand) Help() string {
	return ""
}

func (c *SizesCommand) Run(args []string) int {

	sizes, hr, err := c.Client.SizesService.List()

	if err != nil {
		fmt.Printf("error: %v\n\n", err)
		return 1
	}

	if hr.StatusCode != 200 {
		fmt.Printf("http response error: %+v %+v \n\n", hr, err)
		return 1
	}

	for _, size := range sizes {
		fmt.Printf("slug:%5s memory:%6dmb vcpus:%2d disk:%3dgb transfer:%1.1ftb monthly:%3.1f$\n",
			size.Slug, size.Memory, size.Vcpus, size.Disk, size.Transfer, size.PriceMonthly)
	}

	return 0
}

func (c *SizesCommand) Synopsis() string {
	return fmt.Sprintf("Show available droplet sizes")
}
