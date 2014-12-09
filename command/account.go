package command

import (
	"fmt"
	"github.com/masayukioguni/go-digitalocean/digitalocean"
)

type AccountCommand struct {
	Client *digitalocean.Client
}

func (c *AccountCommand) Help() string {
	return ""
}

func (c *AccountCommand) Run(args []string) int {

	account, hr, err := c.Client.AccountService.GetUserInformation()

	if err != nil {
		fmt.Printf("error: %v\n\n", err)
		return 1
	}

	if hr.StatusCode != 200 {
		fmt.Printf("http response error: %+v %+v \n\n", hr, err)
		return 1
	}

	fmt.Printf("%s(%s) verified:%t limit:%d\n",
		account.Email, account.UUID, account.EmailVerified, account.DropletLimit)

	return 0
}

func (c *AccountCommand) Synopsis() string {
	return fmt.Sprintf("Show account information")
}
