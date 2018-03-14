package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/smallfish/simpleyaml"
	"github.com/urfave/cli"
)

func main() {
	var manifestFile string
	var releaseName string
	var releaseSHA1 string
	var releaseURL string
	var releaseVersion string

	app := cli.NewApp()
	app.Name = "manifest-updaterer"
	app.Usage = "todo"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "manifest",
			Value:       "manifest.yml",
			Usage:       "Path to manifest file",
			Destination: &manifestFile,
		},
		cli.StringFlag{
			Name:        "name",
			Usage:       "Release name",
			Destination: &releaseName,
		},
		cli.StringFlag{
			Name:        "sha",
			Usage:       "Release sha1",
			Destination: &releaseSHA1,
		},
		cli.StringFlag{
			Name:        "url",
			Usage:       "Public release download url",
			Destination: &releaseURL,
		},
		cli.StringFlag{
			Name:        "release-version",
			Usage:       "Release Version",
			Destination: &releaseVersion,
		},
	}
	app.Action = func(c *cli.Context) error {
		d, err := ioutil.ReadFile(manifestFile)
		data := string(d)
		if err != nil {
			log.Fatal(err)
		}

		y, err := simpleyaml.NewYaml(data)
		if err != nil {
			log.Fatal(err)
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
