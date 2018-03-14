package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/urfave/cli"
	yaml "gopkg.in/yaml.v2"
)

type Manifest struct {
	Releases []Release `yml: releases`
}
type Release struct {
	Name    string `yml: name`
	Version string `yml: version`
	URL     string `yml: url`
}

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
		source, err := ioutil.ReadFile(manifestFile)
		if err != nil {
			log.Println(err)
		}
		var manifest Manifest
		yamlErr := yaml.Unmarshal(source, &manifest)
		if yamlErr != nil {
			log.Println(yamlErr)
		}
		for i, release := range manifest.Releases {
			if release.Name == releaseName {
				manifest.Releases[i].Version = releaseVersion
				manifest.Releases[i].URL = releaseURL
			}
		}
		var output []byte
		output, err = yaml.Marshal(&manifest)
		fmt.Printf("%s", output)

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
