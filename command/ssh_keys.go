package command

import (
	"fmt"
	"github.com/masayukioguni/go-digitalocean/digitalocean"
)

type SSHKeyCommand struct {
	Client *digitalocean.Client
}

func (c *SSHKeyCommand) Help() string {
	return ""
}

func (c *SSHKeyCommand) Run(args []string) int {

	SSHKeys, hr, err := c.Client.SSHKeysService.List()

	if err != nil {
		fmt.Printf("error: %v\n\n", err)
		return 1
	}

	if hr.StatusCode != 200 {
		fmt.Printf("http response error: %+v %+v \n\n", hr, err)
		return 1
	}

	for _, key := range SSHKeys {
		fmt.Printf("id:%d name:%s\n", key.ID, key.Name)
	}

	return 0
}

func (c *SSHKeyCommand) Synopsis() string {
	return fmt.Sprintf("Show available SSH keys")
}
