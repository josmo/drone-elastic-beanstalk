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

	workspace := drone.Workspace{}
	repo := drone.Repo{}
	build := drone.Build{}
	vargs := Params{}

	plugin.Param("workspace", &workspace)
	plugin.Param("repo", &repo)
	plugin.Param("build", &build)
	plugin.Param("vargs", &vargs)
	plugin.MustParse()

	if len(vargs.AccessKeyID) == 0 {
		fmt.Println("Please provide an access key")

		os.Exit(1)
		return
	}

	if len(vargs.SecretAccessKey) == 0 {
		fmt.Println("Please provide a secret key")

		os.Exit(1)
		return
	}

	if len(vargs.Region) == 0 {
		fmt.Println("Please provide a region")

		os.Exit(1)
		return
	}

	svc := elasticbeanstalk.New(
		session.New(&aws.Config{
			Region:      aws.String(vargs.Region),
			Credentials: credentials.NewStaticCredentials(vargs.AccessKeyID, vargs.SecretAccessKey, ""),
		}))

	resp, err := svc.CreateApplicationVersion(
		&elasticbeanstalk.CreateApplicationVersionInput{
			ApplicationName:       aws.String("ApplicationName"),
			VersionLabel:          aws.String("VersionLabel"),
			AutoCreateApplication: aws.Bool(true),
			Description:           aws.String("Description"),
			Process:               aws.Bool(true),
			SourceBundle: &elasticbeanstalk.S3Location{
				S3Bucket: aws.String("S3Bucket"),
				S3Key:    aws.String("S3Key"),
			},
		})

	if err != nil {
		fmt.Println(err.Error())

		os.Exit(1)
		return
	}

	fmt.Println(resp)
}
