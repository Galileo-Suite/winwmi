package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {

	flags1 := []cli.Flag{
		&cli.StringFlag{
			Name:    "lang",
			Usage:   "Set the language",
			Value:   "English",
			Aliases: []string{"l"},
			EnvVars: []string{"LEGACY_COMPAT_LANG", "APP_LANG", "LANG"},
		},
	}

	flags2 := []cli.Flag{
		&cli.StringFlag{
			Name:    "os",
			Usage:   "Set the Operating System",
			Value:   "Windows",
			Aliases: []string{"o"},
			EnvVars: []string{"TOMMY"},
		},
	}

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "add a task to the list",
				Action: func(c *cli.Context) error {
					fmt.Println("added task: ", c.Args().First())
					return nil
				},
				Flags: flags1,
			},
			{
				Name:    "delete",
				Aliases: []string{"a"},
				Usage:   "delete a task from the list",
				Action: func(c *cli.Context) error {
					fmt.Println("added task: ", c.Args().First())
					return nil
				},
				Flags: flags2,
			},
		},
		Flags: flags1,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
