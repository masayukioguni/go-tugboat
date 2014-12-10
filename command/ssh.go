package command

import (
	"errors"
	"flag"
	"fmt"
	"github.com/masayukioguni/go-digitalocean/digitalocean"
	"os"
	"os/exec"
	"strings"
)

type SSHCommand struct {
	Client *digitalocean.Client
}

type SSHFlags struct {
	Name    string
	ID      string
	SSHUser string
	SSHPort int
}

func (c *SSHCommand) Help() string {
	helpText := `
Usage: go-tugboat ssh [options] 

  SSH into a droplet

Options:
  -name=[NAME]    The exact name of the droplet

`
	return strings.TrimSpace(helpText)
}

func (c *SSHCommand) getList() ([]digitalocean.Droplet, error) {
	droplets, hr, err := c.Client.DropletsService.List()

	if err != nil {
		return nil, err
	}

	if hr.StatusCode != 200 {
		return nil, err
	}

	return droplets, err

}

func (c *SSHCommand) parse(args []string) (*SSHFlags, error) {
	flags := new(SSHFlags)

	cmdFlags := flag.NewFlagSet("build", flag.ContinueOnError)

	cmdFlags.StringVar(&flags.Name, "name", "", "droplet name")

	err := cmdFlags.Parse(args)

	if err != nil {
		return nil, err
	}

	if flags.Name == "" {
		return nil, errors.New("invalid name")
	}

	return flags, nil
}

func (c *SSHCommand) isMatchName(name string, droplets []digitalocean.Droplet) (*digitalocean.Droplet, error) {
	for _, droplet := range droplets {
		if name == droplet.Name {
			return &droplet, nil
		}
	}
	return nil, errors.New("droplet not found")

}

func (c *SSHCommand) Run(args []string) int {

	flags, err := c.parse(args)

	if err != nil {
		fmt.Printf("comannd parse error : %v\n\n", err)
		return 1
	}

	droplets, err := c.getList()

	if err != nil {
		fmt.Printf("error: %v\n\n", err)
		return 1
	}

	droplet, err := c.isMatchName(flags.Name, droplets)

	if err != nil {
		fmt.Printf("error: %s not found \n\n", flags.Name)
		return 1
	}

	s := fmt.Sprintf("root@%s", GetV4IPAddress(*droplet)[0])

	cmd := exec.Command("ssh", s)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	return 0
}

func (c *SSHCommand) Synopsis() string {
	return fmt.Sprintf("SSH into a droplet")
}
