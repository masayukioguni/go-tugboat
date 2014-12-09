package command

import (
	"github.com/masayukioguni/go-digitalocean/digitalocean"
)

func GetV4IPAddress(droplet digitalocean.Droplet) []string {
	var ipv4s []string

	for _, v4 := range droplet.Networks.V4s {
		ipv4s = append(ipv4s, v4.IPAddress)
	}

	return ipv4s
}
