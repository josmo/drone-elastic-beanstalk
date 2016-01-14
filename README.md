# drone-elastic-beanstalk

[![Build Status](http://beta.drone.io/api/badges/drone-plugins/drone-elastic-beanstalk/status.svg)](http://beta.drone.io/drone-plugins/drone-elastic-beanstalk)
[![](https://badge.imagelayers.io/plugins/drone-elastic-beanstalk:latest.svg)](https://imagelayers.io/?images=plugins/drone-elastic-beanstalk:latest 'Get your own badge on imagelayers.io')

Drone plugin for deploying to AWS Elastic Beanstalk

## Usage

```sh
./drone-elastic-beanstalk <<EOF
{
    "repo": {
        "clone_url": "git://github.com/drone/drone",
        "full_name": "drone/drone"
    },
    "build": {
        "event": "push",
        "branch": "master",
        "commit": "436b7a6e2abaddfd35740527353e78a227ddcb2c",
        "ref": "refs/heads/master"
    },
    "workspace": {
        "root": "/drone/src",
        "path": "/drone/src/github.com/drone/drone"
    },
    "vargs": {
        "access_key": "970d28f4dd477bc184fbd10b376de753",
        "secret_key": "9c5785d3ece6a9cdefa42eb99b58986f9095ff1c",
        "region": "us-east-1",
        "version_label": "v1",
        "description": "Deployed with DroneCI",
        "auto_create": true,
        "bucket_name": "my-bucket-name",
        "bucket_key": "970d28f4dd477bc184fbd10b376de753"
    }
}
EOF
```

## Docker

Build the Docker container using `make`:

```
make deps build docker
```

### Example

```sh
docker run -i plugins/drone-elastic-beanstalk <<EOF
{
    "repo": {
        "clone_url": "git://github.com/drone/drone",
        "full_name": "drone/drone"
    },
    "build": {
        "event": "push",
        "branch": "master",
        "commit": "436b7a6e2abaddfd35740527353e78a227ddcb2c",
        "ref": "refs/heads/master"
    },
    "workspace": {
        "root": "/drone/src",
        "path": "/drone/src/github.com/drone/drone"
    },
    "vargs": {
        "access_key": "970d28f4dd477bc184fbd10b376de753",
        "secret_key": "9c5785d3ece6a9cdefa42eb99b58986f9095ff1c",
        "region": "us-east-1",
        "version_label": "v1",
        "description": "Deployed with DroneCI",
        "auto_create": true,
        "bucket_name": "my-bucket-name",
        "bucket_key": "970d28f4dd477bc184fbd10b376de753"
    }
}
EOF
```
