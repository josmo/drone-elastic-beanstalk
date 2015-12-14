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
    }
}
EOF
```
