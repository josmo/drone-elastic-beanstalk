# drone-elastic-beanstalk

[![Build Status](http://beta.drone.io/api/badges/drone-plugins/drone-elastic-beanstalk/status.svg)](http://beta.drone.io/drone-plugins/drone-elastic-beanstalk)
[![Coverage Status](https://aircover.co/badges/drone-plugins/drone-elastic-beanstalk/coverage.svg)](https://aircover.co/drone-plugins/drone-elastic-beanstalk)
[![](https://badge.imagelayers.io/plugins/drone-elastic-beanstalk:latest.svg)](https://imagelayers.io/?images=plugins/drone-elastic-beanstalk:latest 'Get your own badge on imagelayers.io')

Drone plugin to deploy or update a project on AWS Elastic Beanstalk. For the usage information and a listing of the available options please take a look at [the docs](DOCS.md).

## Binary

Build the binary using `make`:

```
make deps build
```



## Docker

Build the container using `make`:

```
make deps docker
```

## Usage

Build and publish from your current working directory:

```
docker run --rm                     \
  -e PLUGIN_BUCKET=<bucket>         \
  -e AWS_ACCESS_KEY_ID=<token>      \
  -e AWS_SECRET_ACCESS_KEY=<secret> \
  -e PLUGIN_APPLICATION_NAME=<app>  \
  -e PLUGIN_ENVIRONMENT_NAME=<env>  \
  -e PLUGING_BUCKET_KEY=<bucketkey> \
  -v $(pwd):$(pwd)                  \
  -w $(pwd)                         \
  plugins/s3 --dry-run
```

