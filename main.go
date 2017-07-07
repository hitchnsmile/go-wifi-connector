package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
	"os/exec"
)

func main() {

	app := cli.NewApp()
	app.Name = "Go Wifi Connector"
	app.Version = "1.0.0"
	app.UsageText = "gofi connect [SSID] [PASSWORD]"

	_, err := exec.Command("nmcli", "n", "on").Output()

	if err != nil {
		fmt.Println(err)
	}

	app.Commands = []cli.Command{
		{
			Name:    "list",
			Aliases: []string{"ls"},
			Usage:   "list available wifi `list`",
			Action: func(c *cli.Context) error {
				out, err := exec.Command("nmcli", "d", "wifi", "list").Output()
				if err != nil {
					fmt.Println(err)
					return err
				}
				fmt.Println(string(out[:len(out)-1]))
				return nil
			},
		},
		{
			Name:    "connect",
			Aliases: []string{"c"},
			Usage:   "connect to selected wifi `connect SSID password`",
			Action: func(c *cli.Context) error {

				if c.NArg() > 1 {
					_, err := exec.Command("nmcli", "d", "wifi", "connect", c.Args()[0], "password", c.Args()[1]).Output()
					fmt.Println(c.Args()[0], " ", c.Args()[1])
					if err != nil {
						fmt.Println("Failed to connect to ", c.Args()[0], " with password ", c.Args()[1])
						fmt.Println(err)
						return err
					}
				} else if c.NArg() > 0 {
					_, err := exec.Command("nmcli", "d", "wifi", "connect", c.Args()[0]).Output()
					fmt.Println(c.Args()[0])
					if err != nil {
						fmt.Println("Failed to connect to ", c.Args().First())
						fmt.Println(err)
						return err
					}
				}

				fmt.Println("Successfully connected to ", c.Args().First())

				return nil
			},
		},
	}

	app.Run(os.Args)
}
