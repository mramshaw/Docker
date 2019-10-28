# Docker

![Docker icon](images/Docker-R-Logo-08-2018-Monochomatic-RGB_Moby-x1.png)

In terms of the __Cloud__, learning [Docker](https://www.docker.com/) is a
great first step.

All of the necessary concepts are present in Docker; cloud providers offer
the same functionality (albeit in somewhat different forms) but the basics
are all present in Docker.

In addition, [DockerHub](http://hub.docker.com/u/mramshaw4docs) is a great
way to share packaged software.

Each cloud provider (AWS, Azure, GCP, Digital Ocean) has multiple services,
including VMs (Virtual Machines). They also provide (generally as a premium
offering) Container services.

Various __orchestration__ services are available, but all of the cloud providers
listed offer __Kubernetes__ (which is not to say that all of the offerings are
identical - some lag behind).

As a prelude to learning an orchestration service (such as Kubernetes), much
can be learned by starting with [docker-compose](https://docs.docker.com/compose/).
However, the limitations of this approach quickly become apparent, but for
prototyping purposes it is a great tool.

Even so, Docker functions as an introduction to modern methodology such as
the [Twelve-factor App](http://12factor.net/processes).

> __Twelve-factor processes are stateless and share-nothing.__

[This accords well with [Microservices](http://github.com/mramshaw/microservices) thinking.]

__Nota bene__:

> Sticky sessions are a violation of twelve-factor and should never be used or relied upon.

## Contents

The contents are as follows:

* [Terminology](#terminology)
* [Process](#process)
* [Best Practices](#best-practices)
    * [.dockerignore](#dockerignore)
    * [FROM scratch](#from-scratch)
    * [apt-get update / apk --update](#apt-get-update--apk---update)
    * [sort dependencies](#sort-dependencies)
    * [root access](#root-access)
    * [Tag all images](#tag-all-images)
* [ADD versus COPY](#add-versus-copy)
* [Useful Shortcuts](#useful-shortcuts)
    * [Detach from a container (but leave it running)](#detach-from-a-container-but-leave-it-running)
    * [Exit a container](#exit-a-container)
* [Useful Commands](#useful-commands)
    * [General information about Docker and the Docker runtime](#general-information-about-docker-and-the-docker-runtime)
    * [Current information about Docker runtime](#current-information-about-docker-runtime)
    * [Processes](#processes)
    * [Visibility into containers](#visibility-into-containers)
    * [Images](#images)
    * [Search Images](#search-images)
    * [Tag Image](#tag-image)
    * [Layer information about the Docker image](#layer-information-about-the-docker-image)
    * [Docker logs](#docker-logs)
    * [Docker volumes](#docker-volumes)

## Terminology

Docker itself is not too hard to learn and use, but in case you are ever asked
the difference between a __container__ and an __image__, here is my definition:

> A Docker __container__ wraps one or more Docker __images__ into a process
> with everything needed to run an application.

[This is from an ___operating system___ point of view; from an ___application___
 point of view, it is of course your responsibility to add all of the software
 components needed to run your application.]

In general, the base image will correspond to a \*nix operating system (such
as CentOS) or a language (such as Node.js). In the case of a language image,
all of the language components will be layered on top of a (probably) linux
base image (make a note of which one - it will be important).

[Note that __DOCKER__ itself will handle the operating system issues, which
 means that \*nix Docker images can be run on any hardware (such as a MacBook
 or Windows desktop) for which a version of Docker can be installed.]

## Process

My personal Docker process is to start by selecting a ___Base image___ (all
subsequent software layers will then be layered on top of this image). This
base image might be __Alpine__, __CentOS__, __Debian__, __Ubuntu__ or even
a language image such as __Node.js__ or __Python__ or __Golang__ (these are
simply my usual choices and of course are not the only options - in fact the
available choices for a base Docker image are pretty much endless).

Generally, you should pick whichever linux distro you know best, unless you
decide to go with a language option - in which case the choice will probably
be obvious. If unsure which to pick, __CentOS__ is probably a safe choice.

[Docker base images are generally linux-based, as the technology is based
 upon Linux Containers - which were themselves based upon
 [chroot jails](http://en.wikipedia.org/wiki/Chroot).]

In very rare situations I may decide to build [FROM scratch](#from-scratch),
for instance if I am simply distributing a pre-built binary.

And I record the base image decision in a __Dockerfile__ (my preference is to
always use a __Dockerfile__ as they allow for simple and repeatable builds).

Then I add my application's dependencies.

And then I add my application.

I will normally ___build___ my application in Docker (perhaps using a
___buildbot___) and also ___test it___ in Docker. If these succeed, then
I leave my application up and running.

There are variations on this theme, but that's generally my process.

## Best Practices

The following are some suggested best practices:

#### .dockerignore

Create a __.dockerignore__ file (essentially the same thing as a __.gitignore__ file)
and list everything that Docker doesn't need to see in it (for instance __README.md__,
__.git/__ and __.gitignore__ as well as any __passwords__ or __secrets__). This can also
save some transfer time in certain cases, such as when there is a lot of test data or
source files.

Interestingly, Docker doesn't normally need to see the __.dockerignore__ file, so the
first entry in this file should be:

    .dockerignore

Create recursive wildcard patterns as follows:

    **/*.obj
    **/*.pyc

[In this case, all .obj or .pyc files - in _any_ folder - will be ignored by Docker.]

#### FROM scratch

In some circumstances you may wish to build your container starting from an empty image.

[This is unusual but it is possible to think of many use cases where this would be desirable.]

The syntax for doing this is:

```
FROM scratch
```

You can read the documentation for this option here:

    http://docs.docker.com/samples/library/scratch/

#### apt-get update / apk --update

In order to run an apples-to-apples comparison when testing/benchmarking, do __NOT__ do
either an __apt-get update__ or an __apk --update__. The reason for this is that these
will make subtle changes to the operating systems that may well invalidate any test
comparisons.

Of course, for production use, these are strongly recommended for security reasons.

#### sort dependencies

This is so obvious that it almost goes without saying, but when installing dependencies
with __apt-get__ or __apk__ take the time to sort them alphabetically (or in some other
order if that makes more sense).

Large numbers of dependencies can be hard to scan, so sorting them can make it easier
to see what dependencies are being installed.

#### root access

Consider specifying __USER nobody__ unless __root__ access is absolutely required.
[Hint: it is almost never actually required.]

#### Tag all images

The default tag is __latest__ but it is a really terrible practice to use this. For one thing
using it means being vulnerable to new releases (not generally a good idea) - and can also
have a big impact in terms of having to actually download the latest release.

For consistency reasons, always tag Docker images - and only ever use tagged images. This will
at least give you the option of actually __testing__ the latest release.

## ADD versus COPY

ADD will uncompress certain files (.tar, .tar.gz, .tar.bz2) if only a destination directory
is specified. It will __not__ uncompress certain other files (.zip). If the intent is to
include a compressed file, make sure to specify a destination __name__ and directory.

COPY will simply copy the file, without uncompressing it (usually this is what's wanted).

If unsure, use COPY.

## Useful Shortcuts

[The usual practice on OS/X is to substitute the Option key for the Ctrl key;
 these are terminal commands, so don't.]

#### Detach from a container (but leave it running)

It's not always convenient or possible to set up another terminal:

	Ctrl-P followed by Ctrl-Q

For instance, when running in a Cloud Shell.

To reattach to the running container:

	$ docker attach xxxxxxxxxxxx

#### Exit a container

Slightly faster than typing `exit`:

	Ctrl-D

## Useful Commands

One or two useful Docker commands.

#### General information about Docker and the Docker runtime:

	$ docker info

#### Current information about the Docker runtime:

	$ docker stats

As usual, Ctrl-C to stop.

#### Processes:

	$ docker ps

	$ docker ps -qa

	$ docker rm ...

#### Visibility into containers

See into running containers with <kbd>docker top</kbd>.

Start a container:

```bash
$ docker run --name example -it busybox sh
/ #
```

Detach from this container (but leave it running) with <kbd>Ctrl-P<kbd> followed by <kbd>Ctrl-Q</kbd>.

To see what the container is running:

```bash
$ docker top example
UID                 PID                 PPID                C                   STIME               TTY                 TIME                CMD
root                9795                9777                0                   20:21               pts/0               00:00:00            sh
$
```

And clean up:

```bash
$ docker kill example
example
$ docker rm example
example
$
```

#### Images:

	$ docker images

	$ docker rmi ...

#### Search Images:

This can be significantly faster than searching [DockerHub](https://hub.docker.com/)
(or any other repository) manually:

	$ docker search rediswebserver

[Replace `rediswebserver` with whatever software is of interest.]

Official builds will show first. By default only __25__ images will be shown, but
this is usually more than enough. The __STARS__ column is particularly helpful:

```bash
$ docker search redis --no-trunc
NAME                             DESCRIPTION                                                                            STARS               OFFICIAL            AUTOMATED
redis                            Redis is an open source key-value store that functions as a data structure server.     6978                [OK]
bitnami/redis                    Bitnami Redis Docker Image                                                             113                                     [OK]
sameersbn/redis                                                                                                         75                                      [OK]
grokzen/redis-cluster            Redis cluster 3.0, 3.2, 4.0 & 5.0                                                      48
kubeguide/redis-master           redis-master with "Hello World!"                                                       29
rediscommander/redis-commander   Alpine image for redis-commander - Redis management tool.                              24                                      [OK]
redislabs/redis                  Clustered in-memory database engine compatible with open source Redis by Redis Labs    20
arm32v7/redis                    Redis is an open source key-value store that functions as a data structure server.     15
redislabs/redisearch             Redis With the RedisSearch module pre-loaded. See http://redisearch.io                 15
oliver006/redis_exporter          Prometheus Exporter for Redis Metrics. Supports Redis 2.x, 3.x, 4.x and 5.x           10
webhippie/redis                  Docker images for Redis                                                                10                                      [OK]
s7anley/redis-sentinel-docker    Redis Sentinel                                                                         8                                       [OK]
insready/redis-stat              Docker image for the real-time Redis monitoring tool redis-stat                        7                                       [OK]
arm64v8/redis                    Redis is an open source key-value store that functions as a data structure server.     6
redislabs/redisgraph             A graph database module for Redis                                                      5                                       [OK]
centos/redis-32-centos7          Redis in-memory data structure store, used as database, cache and message broker       4
bitnami/redis-sentinel           Bitnami Docker Image for Redis Sentinel                                                4                                       [OK]
frodenas/redis                   A Docker Image for Redis                                                               2                                       [OK]
circleci/redis                   CircleCI images for Redis                                                              2                                       [OK]
wodby/redis                      Redis container image with orchestration                                               2                                       [OK]
kilsoo75/redis-master            This image is for the redis master of SK CloudZ                                        1
tiredofit/redis                  Redis Server w/ Zabbix monitoring and S6 Overlay based on Alpine                       1                                       [OK]
cflondonservices/redis           Docker image for running redis                                                         0
xetamus/redis-resource           forked redis-resource                                                                  0                                       [OK]
$
```

#### Tag Image:

Docker images are generally tagged __latest__ by default.

As each successive __latest__ image is built, it usually leaves behind an unnamed orphan image.
Use __rmi__ (see above) to delete these.

To tag an image with a version number (in this case 1.1):

	$ docker tag xxx/yyy xxx/yyy:1.1

#### Layer information about the Docker image:

	$ docker history ...

	$ docker history --no-trunc ...

#### Docker logs:

	$ docker logs xxxxxxxxxxxx

#### Docker volumes:

	$ docker volume ls

	$ docker volume prune
