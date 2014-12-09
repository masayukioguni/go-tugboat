package command

import (
	"flag"
	"fmt"
	"github.com/masayukioguni/go-digitalocean/digitalocean"
	"net/http"
	"strings"
)

type ImagesCommand struct {
	Client *digitalocean.Client
}

func (c *ImagesCommand) Help() string {
	helpText := `
Usage: go-tugboat image [options] 

  Get a list of images that are provided in the digitalocean.

Options:

  -type=dist  Only distribution images
  -type=app   Only application  images

`
	return strings.TrimSpace(helpText)
}

func (c *ImagesCommand) Run(args []string) int {

	var typeFlag string
	cmdFlags := flag.NewFlagSet("build", flag.ContinueOnError)

	cmdFlags.StringVar(&typeFlag, "type", "", "include the type query paramter")
	cmdFlags.StringVar(&typeFlag, "t", "", "include the type query paramter")

	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	var images []digitalocean.Image
	var hr *http.Response
	var err error

	switch typeFlag {
	case "dist":
		images, hr, err = c.Client.ImagesService.ListDistribution()
	case "app":
		images, hr, err = c.Client.ImagesService.ListApplication()
	default:
		images, hr, err = c.Client.ImagesService.ListAll()
	}

	if err != nil {
		fmt.Printf("error: %v\n\n", err)
		return 1
	}

	if hr.StatusCode != 200 {
		fmt.Printf("http response error: %+v %+v \n\n", hr, err)
		return 1
	}

	for _, image := range images {
		fmt.Printf("%s (id: %d, distro: %s)\n",
			image.Name, image.ID, image.Distribution)
	}

	return 0
}

func (c *ImagesCommand) Synopsis() string {
	return fmt.Sprintf("Retrieve a image list")
}
