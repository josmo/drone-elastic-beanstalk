package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"os"
)

var version string

func main() {
	app := cli.NewApp()
	app.Name = "Beanstalk deployment plugin"
	app.Usage = "beanstalk deployment plugin"
	app.Action = run
	app.Version = version
	app.Flags = []cli.Flag{

		cli.StringFlag{
			Name:   "access-key",
			Usage:  "aws access key",
			EnvVar: "PLUGIN_ACCESS_KEY,AWS_ACCESS_KEY_ID",
		},
		cli.StringFlag{
			Name:   "secret-key",
			Usage:  "aws secret key",
			EnvVar: "PLUGIN_SECRET_KEY,AWS_SECRET_ACCESS_KEY",
		},
		cli.StringFlag{
			Name:   "bucket",
			Usage:  "aws bucket",
			Value:  "us-east-1",
			EnvVar: "PLUGIN_BUCKET",
		},
		cli.StringFlag{
			Name:   "region",
			Usage:  "aws region",
			Value:  "us-east-1",
			EnvVar: "PLUGIN_REGION",
		},

		cli.StringFlag{
			Name:   "bucket-key",
			Usage:  "upload files from source folder",
			EnvVar: "PLUGIN_BUCKET_KEY",
		},
		cli.StringFlag{
			Name:   "application",
			Usage:  "application name for beanstalk",
			EnvVar: "PLUGIN_APPLICATION",
		},
		cli.StringFlag{
			Name:   "environment-name",
			Usage:  "environment name in the app to update",
			EnvVar: "PLUGIN_ENVIRONMENT_NAME",
		},
		cli.StringFlag{
			Name:   "version-label",
			Usage:  "version label for the app",
			EnvVar: "PLUGIN_VERSION_LABEL",
		},
		cli.StringFlag{
			Name:   "description",
			Usage:  "description for the app version",
			EnvVar: "PLUGIN_DESCRIPTION",
		},
		cli.StringFlag{
			Name:   "auto-create",
			Usage:  "auto create app if it doesn't exist",
			EnvVar: "PLUGIN_AUTO_CREATE",
		},
		cli.StringFlag{
			Name:   "process",
			Usage:  "Preprocess and validate manifest",
			EnvVar: "PLUGIN_PROCESS",
		},
		cli.StringFlag{
			Name:   "environment-update",
			Usage:  "update the environment",
			EnvVar: "PLUGIN_ENVIRONMENT_UPDATE",
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
func run(c *cli.Context) error {
	plugin := Plugin{
		Key:               c.String("access-key"),
		Secret:            c.String("secret-key"),
		Bucket:            c.String("bucket"),
		BucketKey:         c.String("bucket-key"),
		Application:       c.String("application"),
		EnvironmentName:   c.String("environment-name"),
		VersionLabel:      c.String("version-label"),
		Description:       c.String("description"),
		AutoCreate:        c.Bool("auto-create"),
		Process:           c.Bool("process"),
		EnvironmentUpdate: c.Bool("environment-update"),
		Region:            c.String("region"),
	}

	return plugin.Exec()
}
