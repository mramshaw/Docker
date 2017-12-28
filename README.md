# Docker

In terms of the __Cloud__, learning _Docker_ is a great first step.

All of the necessary concepts are present in Docker; cloud providers offer the same functionality (albeit in somewhat different forms) but the basics are all present in Docker.

Each cloud provider (AWS, Azure, GCP) has multiple services, including VMs (Virtual Machines). They also provide (generally as a premium offering) Container services.

Various __orchestration__ services are available, but all of the cloud providers listed offer __Kubernetes__ (which is not to say that all of the offerings are identical - some lag behind).

## Best Practices

Create a __.dockerfile__ (essentially the same thing as a __.gitignore__ file) and list everything that Docker doesn't need to see in it (for instance __.git__ and __.gitignore__ as well as any __passwords__ or __secrets__). This can also save some transfer time in certain cases, such as when there is a lot of test data or source files.

In order to run an apples-to-apples comparison when testing, do __NOT__ do either an __apt-get update__ or an __apk --update__. The reason for this is that these will make subtle changes to the operating systems that may well invalidate any test comparisons.

Consider specifying __USER nobody__ unless __root__ access is absolutely required.

## ADD versus COPY

ADD will uncompress certain files (.tar, .tar.gz, .tar.bz2) if only a destination directory is specified. It will __not__ uncompress certain other files (.zip). If the intent is to include a compressed file, make sure to specify a destination __name__ and directory.

COPY will simply copy the file, without uncompressing it (usually this is what's wanted).

If unsure, use COPY.

## Useful Shortcuts

[The usual practice on OS/X is to substitute the Option key for the Ctrl key; these are terminal commands, so don't.]

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

#### Processes:

	$ docker ps

	$ docker ps -qa

	$ docker rm ...

#### Images:

	$ docker images

	$ docker rmi ...

#### Search Images:

This can be significantly faster than searching DockerHub (or any other repository) manually:

	$ docker search rediswebserver

Official builds will show first.

#### Tag Image:

Docker images are generally tagged __latest__ by default.

As each succesive __latest__ image is built, it usually leaves behind an unnamed orphan image. Use __rmi__ (see above) to delete these.

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
