package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elasticbeanstalk"
	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/plugin"
)

var (
	build     string
	buildDate string
)

func main() {
	fmt.Printf("Drone Elastic Beanstalk Plugin built at %s\n", buildDate)

	repo := drone.Repo{}
	build := drone.Build{}
	vargs := Params{}

	plugin.Param("repo", &repo)
	plugin.Param("build", &build)
	plugin.Param("vargs", &vargs)
	plugin.MustParse()

	if vargs.Application == "" {
		vargs.Application = repo.Name
	}

	if vargs.AccessKey == "" {
		fmt.Println("Please provide an access key id")
		os.Exit(1)
	}

	if vargs.SecretKey == "" {
		fmt.Println("Please provide a secret access key")
		os.Exit(1)
	}

	if vargs.Region == "" {
		fmt.Println("Please provide a region")
		os.Exit(1)
	}

	if vargs.VersionLabel == "" {
		fmt.Println("Please provide a version label")
		os.Exit(1)
	}

	if vargs.BucketName == "" {
		fmt.Println("Please provide a bucket name")
		os.Exit(1)
	}

	if vargs.BucketKey == "" {
		fmt.Println("Please provide a bucket key")
		os.Exit(1)
	}

	svc := elasticbeanstalk.New(
		session.New(&aws.Config{
			Region: aws.String(vargs.Region),
			Credentials: credentials.NewStaticCredentials(
				vargs.AccessKey,
				vargs.SecretKey,
				"",
			),
		}),
	)

	_, err := svc.CreateApplicationVersion(
		&elasticbeanstalk.CreateApplicationVersionInput{
			VersionLabel:          aws.String(vargs.VersionLabel),
			ApplicationName:       aws.String(vargs.Application),
			Description:           aws.String(vargs.Description),
			AutoCreateApplication: aws.Bool(vargs.AutoCreate),
			Process:               aws.Bool(vargs.Process),
			SourceBundle: &elasticbeanstalk.S3Location{
				S3Bucket: aws.String(vargs.BucketName),
				S3Key:    aws.String(vargs.BucketKey),
			},
		},
	)

	if vargs.EnvironmentUpdate == true {

		if vargs.EnvironmentName == "" {
			fmt.Println("Can't update environment without environment name")
			os.Exit(1)
		}

		_, err = svc.UpdateEnvironment(
			&elasticbeanstalk.UpdateEnvironmentInput{
				VersionLabel:    aws.String(vargs.VersionLabel),
				ApplicationName: aws.String(vargs.Application),
				Description:     aws.String(vargs.Description),
				EnvironmentName: aws.String(vargs.EnvironmentName),
			},
		)
	}

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Successfully deployed")
}
