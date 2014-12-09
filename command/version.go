package command

import (
	"fmt"
)

type VersionCommand struct {
	AppName        string
	Version        string
	LibraryVersion string
}

func (c *VersionCommand) Help() string {
	return ""
}

func (c *VersionCommand) Run(args []string) int {
	fmt.Printf("Version %s/%s\n", c.AppName, c.Version)
	return 1
}

func (c *VersionCommand) Synopsis() string {
	return fmt.Sprintf("Prints the %s version", c.AppName)
}
