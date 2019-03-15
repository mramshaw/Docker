# Docker

In terms of the __Cloud__, learning [Docker](https://www.docker.com/) is a
great first step.

All of the necessary concepts are present in Docker; cloud providers offer
the same functionality (albeit in somewhat different forms) but the basics
are all present in Docker.

Each cloud provider (AWS, Azure, GCP) has multiple services, including VMs
(Virtual Machines). They also provide (generally as a premium offering)
Container services.

Various __orchestration__ services are available, but all of the cloud providers
listed offer __Kubernetes__ (which is not to say that all of the offerings are
identical - some lag behind).

As a prelude to learning an orchestration service (such as Kubernetes), much
can be learned by starting with [docker-compose](https://docs.docker.com/compose/).
However, the limitations of this approach quickly become apparent, but for
prototyping purposes it is a great tool.

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
 or Windows desktop) for which a version of Docker can been installed.]

## Process

My personal Docker process is to start by selecting a ___Base image___ (all
subsequent software layers will then be layered on top of this image). This
base image might be __Alpine__, __CentOS__, __Debian__, __Ubuntu__ or even
a language image such as __Node.js__ or __Python__ or __Golang__.

Generally, you should pick whichever linux distro you know best, unless you
decide to go with a language option - in which case the choice will probably
be obvious. If unsure which to pick, __CentOS__ is probably a safe choice.

[Docker base images are generally linux-based, as the technology is based
 upon Linux Containers - which were themselves based upon
 [chroot jails](http://en.wikipedia.org/wiki/Chroot).]

And I record this decision in a __Dockerfile__ (my preference is to always
use a __Dockerfile__ as they allow for simple and repeatable builds).

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

Create recursive wildcard patterns as follows:

    **/*.obj
    **/*.pyc

[In this case, all .obj or .pyc files - in _any_ folder - will be ignored by Docker.]

#### apt-get update / apk --update

In order to run an apples-to-apples comparison when testing/benchmarking, do __NOT__ do
either an __apt-get update__ or an __apk --update__. The reason for this is that these
will make subtle changes to the operating systems that may well invalidate any test
comparisons.

Of course, for production use, these are strongly recommended for security reasons.

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

#### Current information about Docker runtime:

	$ docker stats

As usual, Ctrl-C to stop.

#### Processes:

	$ docker ps

	$ docker ps -qa

	$ docker rm ...

#### Images:

	$ docker images

	$ docker rmi ...

#### Search Images:

This can be significantly faster than searching [DockerHub](https://hub.docker.com/)
(or any other repository) manually:

	$ docker search rediswebserver

[Replace `rediswebserver` with whatever software is of interest.]

Official builds will show first.

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
