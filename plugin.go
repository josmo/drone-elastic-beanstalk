package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elasticbeanstalk"
)

// Plugin defines the beanstalk plugin parameters.
type Plugin struct {
	Key    string
	Secret string
	Bucket string

	// us-east-1
	// us-west-1
	// us-west-2
	// eu-west-1
	// ap-southeast-1
	// ap-southeast-2
	// ap-northeast-1
	// sa-east-1
	Region string

	BucketKey         string
	Application       string
	EnvironmentName   string
	VersionLabel      string
	Description       string
	AutoCreate        bool
	Process           bool
	EnvironmentUpdate bool
}

func (p *Plugin) client() *elasticbeanstalk.ElasticBeanstalk {

	// Use key and secret if provided otherwise fall back to ec2 instance profile
	if p.Key != "" && p.Secret != "" {
		return elasticbeanstalk.New(session.New(), &aws.Config{
			Credentials: credentials.NewStaticCredentials(p.Key, p.Secret, ""),
			Region:      aws.String(p.Region),
		})
	} else {
		return elasticbeanstalk.New(session.New(), &aws.Config{
			Region: aws.String(p.Region),
		})
	}
}

// Exec runs the plugin
func (p *Plugin) Exec() error {
	// create the client

	client := p.client()



	log.WithFields(log.Fields{
		"region":           p.Region,
		"application-name": p.Application,
		"environment":      p.EnvironmentName,
		"bucket":           p.Bucket,
		"bucket-key":       p.BucketKey,
		"versionlabel":     p.VersionLabel,
		"description":      p.Description,
		"env-update":       p.EnvironmentUpdate,
		"auto-create":	    p.AutoCreate,
	}).Info("Attempting to create and update")

	_, err := client.CreateApplicationVersion(
		&elasticbeanstalk.CreateApplicationVersionInput{
			VersionLabel:          aws.String(p.VersionLabel),
			ApplicationName:       aws.String(p.Application),
			Description:           aws.String(p.Description),
			AutoCreateApplication: aws.Bool(p.AutoCreate),
			Process:               aws.Bool(p.Process),
			SourceBundle: &elasticbeanstalk.S3Location{
				S3Bucket: aws.String(p.Bucket),
				S3Key:    aws.String(p.BucketKey),
			},
		},
	)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Problem create application")
		return err
	}
	if p.EnvironmentUpdate == true && err == nil {

		if p.EnvironmentName == "" {
			return fmt.Errorf("Can't update environment without environment name")
		}

		_, err = client.UpdateEnvironment(
			&elasticbeanstalk.UpdateEnvironmentInput{
				VersionLabel:    aws.String(p.VersionLabel),
				ApplicationName: aws.String(p.Application),
				Description:     aws.String(p.Description),
				EnvironmentName: aws.String(p.EnvironmentName),
			},
		)
		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
			}).Error("Problem updating beanstalk")
			return err
		}
	}
	return nil

}
