[![Build Status](https://cloud.drone.io/api/badges/josmo/drone-elastic-beanstalk/status.svg)](https://cloud.drone.io/josmo/drone-elastic-beanstalk)
[![Go Doc](https://godoc.org/github.com/josmo/drone-elastic-beanstalk?status.svg)](http://godoc.org/github.com/josmo/drone-elastic-beanstalk)
[![Go Report](https://goreportcard.com/badge/github.com/josmo/drone-elastic-beanstalk)](https://goreportcard.com/report/github.com/josmo/elastic-beanstalk)
[![](https://images.microbadger.com/badges/image/peloton/drone-elastic-beanstalk.svg)](https://microbadger.com/images/peloton/drone-elastic-beanstalk "Get your own image badge on microbadger.com")

# drone-elastic-beanstalk

Drone plugin to deploy or update a project on AWS Elastic Beanstalk. For the
usage information and a listing of the available options please take a look at
[the docs](DOCS.md).

## Build

Build the binary with the following commands:

```
drone exec
```

## Docker

Build the docker image with the following commands:

```
drone exec
docker build --rm=true -t plugins/elastic-beanstalk .
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
  peloton/drone-elastic-beanstalk
```


### Contribution

This repo is setup in a way that if you enable a personal drone server to build your fork it will
 build and publish your image (makes it easier to test PRs and use the image till the contributions get merged)
 
* Build local ```DRONE_REPO_OWNER=josmo DRONE_REPO_NAME=drone-elastic-beanstalk drone exec```
* on your server just make sure you have DOCKER_USERNAME, DOCKER_PASSWORD, and PLUGIN_REPO set as secrets
