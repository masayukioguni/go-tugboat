package command

import (
	"bufio"
	"fmt"
	"github.com/masayukioguni/go-tugboat/config"
	"os"
	"path/filepath"
)

type AuthorizeCommand struct {
}

func (c *AuthorizeCommand) Help() string {
	return ""
}

func (c *AuthorizeCommand) ask(text string, defaultText string) string {
	fmt.Printf(text)
	var scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()

	input := scanner.Text()
	if input != "" {
		return input
	}
	return defaultText

}
func (c *AuthorizeCommand) Run(args []string) int {
	apikey := c.ask("Entser your API key:", "input apikey")
	fmt.Println(`To retrieve region, image, size and key ID's, you can use the corresponding go-tugboat command,
such as go-tugboat images. Defaults can be changed at any time in your ~/.go-tugboat/config.yaml configuration file.\n`)

	region := c.ask("Enter your default region (optional, defaults to nyc3 (New York)):", "nyc3")
	image := c.ask("Enter your default image id(optional, defaults to 6918990(Ubuntu 14.04 x64):", "6918990")
	size := c.ask("Enter your default size (optional, defaults to 512MB):", "512mb")
	key := c.ask("Enter your default ssh key ID (optional, defaults to none):", "")

	env := &config.Config{}
	env.Authentication.APIKey = apikey
	env.Defaluts.Image = image
	env.Defaluts.Region = region
	env.Defaluts.Size = size
	env.Defaluts.Key = key

	home := os.Getenv("HOME")
	if home == "" {
		fmt.Errorf("Error Getenv $HOME not found")
		return 1
	}

	saveDirectory := filepath.Join(home, config.GetDefaultDirectory())
	_, err := os.Stat(saveDirectory)
	if err != nil {
		if err = os.Mkdir(saveDirectory, 0755); err != nil {
			fmt.Errorf("Error mkdir %s", saveDirectory)
			return 0
		}
	}
	savePath := filepath.Join(saveDirectory, config.GetDefaultConfigName())

	err = config.SaveConfig(savePath, env)

	if err != nil {
		fmt.Errorf("Error SaveConfig %s", err)
		return 1
	}

	fmt.Println("Authentication with DigitalOcean was successful!")

	return 0
}

func (c *AuthorizeCommand) Synopsis() string {
	return fmt.Sprintf("Authorize a DigitalOcean account with go-tugboat")
}
