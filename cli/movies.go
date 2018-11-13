package main

import (
	"log"
	"os"

	"github.com/MateuszJonak/movies-api/storage"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "moviescli"
	app.Usage = "CLI for movies api"

	app.Commands = []cli.Command{
		{
			Name:    "database",
			Aliases: []string{"db"},
			Usage:   "options for database",
			Subcommands: []cli.Command{
				{
					Name:  "migrate",
					Usage: "migrate the database",
					Action: func(c *cli.Context) error {
						storage.Migrate()
						return nil
					},
				},
				{
					Name:  "rollback",
					Usage: "rollback the last migration",
					Action: func(c *cli.Context) error {
						storage.Rollback()
						return nil
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
