# drone-elastic-beanstalk

[![Build Status](http://beta.drone.io/api/badges/drone-plugins/drone-elastic-beanstalk/status.svg)](http://beta.drone.io/drone-plugins/drone-elastic-beanstalk)
[![Go Doc](https://godoc.org/github.com/drone-plugins/drone-elastic-beanstalk?status.svg)](http://godoc.org/github.com/drone-plugins/drone-elastic-beanstalk)
[![Go Report](https://goreportcard.com/badge/github.com/drone-plugins/drone-elastic-beanstalk)](https://goreportcard.com/report/github.com/drone-plugins/drone-elastic-beanstalk)
[![Join the chat at https://gitter.im/drone/drone](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/drone/drone)

Drone plugin to deploy or update a project on AWS Elastic Beanstalk. For the
usage information and a listing of the available options please take a look at
[the docs](DOCS.md).

## Build

Build the binary with the following commands:

```
go build
go test
```

## Docker

Build the docker image with the following commands:

```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo
docker build --rm=true -t plugins/elastic-beanstalk .
```

Please note incorrectly building the image for the correct x64 linux and with
GCO disabled will result in an error when running the Docker image:

```
docker: Error response from daemon: Container command
'/bin/drone-elastic-beanstalk' not found or does not exist..
```

## Usage

Execute from the working directory:

```sh
docker run --rm \
  -e PLUGIN_BUCKET=<bucket> \
  -e AWS_ACCESS_KEY_ID=<token> \
  -e AWS_SECRET_ACCESS_KEY=<secret> \
  -e PLUGIN_APPLICATION_NAME=<app> \
  -e PLUGIN_ENVIRONMENT_NAME=<env> \
  -e PLUGING_BUCKET_KEY=<bucketkey> \
  -v $(pwd):$(pwd) \
  -w $(pwd) \
  plugins/elastic-beanstalk
```
