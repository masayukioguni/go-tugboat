package command

import (
	"errors"
	"flag"
	"fmt"
	"github.com/masayukioguni/go-digitalocean/digitalocean"
	"strings"
)

type CreateCommand struct {
	Client        *digitalocean.Client
	DefalutSize   string
	DefalutImage  string
	DefalutRegion string
	DefalutKey    string
}

type CreateFlags struct {
	Name   string
	Image  string
	Size   string
	Region string
	Keys   string
}

func (c *CreateCommand) Help() string {
	helpText := `
Usage: go-tugboat create [Require Options]

  create a new Droplet

Require Options:
  -name=[slug]  The human-readable string you wish to use when displaying the Droplet name.
  -size=[slug]  The unique slug identifier for the size that you wish to select for this Droplet.
  -image=[image id]  The image ID of a public or private image
  -region=[slug]  The unique slug identifier for the region that you wish to deploy in.
  -keys=[KEYS] A comma separated list of SSH key ids to add to the droplet
`
	return strings.TrimSpace(helpText)
}

func (c *CreateCommand) parse(args []string) (*CreateFlags, error) {
	flags := new(CreateFlags)

	cmdFlags := flag.NewFlagSet("build", flag.ContinueOnError)

	cmdFlags.StringVar(&flags.Name, "name", "", "droplet name")
	cmdFlags.StringVar(&flags.Size, "size", "", "size slug")
	cmdFlags.StringVar(&flags.Image, "image", "", "image ID or slug")
	cmdFlags.StringVar(&flags.Region, "region", "", "region slug")
	cmdFlags.StringVar(&flags.Keys, "keys", "", "SSH key ids")

	err := cmdFlags.Parse(args)

	if err != nil {
		return nil, err
	}

	if flags.Name == "" {
		return nil, errors.New("invalid name")
	}

	if flags.Size == "" {
		flags.Size = c.DefalutSize
	}
	if flags.Image == "" {
		flags.Image = c.DefalutImage
	}

	if flags.Region == "" {
		flags.Region = c.DefalutRegion
	}

	if flags.Keys == "" {
		flags.Keys = c.DefalutKey
	}

	return flags, nil
}

func (c *CreateCommand) Run(args []string) int {

	flags, err := c.parse(args)

	if err != nil {
		fmt.Printf("comannd parse error : %v\n\n", err)
		return 0
	}

	keys := strings.Split(flags.Keys, ",")

	input := &digitalocean.CreateDropletRequest{
		Name:              flags.Name,
		Size:              flags.Size,
		Image:             flags.Image,
		Region:            flags.Region,
		SSHKeys:           keys,
		Backups:           false,
		PrivateNetworking: false}

	droplet, hr, err := c.Client.DropletsService.Create(input)

	if err != nil {
		fmt.Printf("error: %v\n\n", err)
		return 1
	}

	if hr.StatusCode != 202 {
		fmt.Printf("http response error: %+v %+v \n\n", hr, err)
		return 1
	}

	fmt.Printf("Queueing creation of droplet '%s' ...done\n", droplet.Name)

	return 0
}

func (c *CreateCommand) Synopsis() string {
	return fmt.Sprintf("Create a droplet.")
}
