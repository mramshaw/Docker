# Docker

In terms of the __Cloud__, learning _Docker_ is a great first step.

All of the necessary concepts are present in Docker; cloud providers offer the same functionality (albeit in somewhat different forms) but the basics are all present in Docker.

Each cloud provider (AWS, Azure, GCP) has multiple services, including VMs (Virtual Machines). They also provide (generally as a premium offering) Container services.

Various __orchestration__ services are available, but all of the cloud providers listed offer __Kubernetes__ (which is not to say that all of the offerings are identical - some lag behind).

## Best Practices

Create a __.dockerfile__ (essentially the same things as a __.gitignore__ file) and list everything that Docker doesn't need to see in it (for instance __.git__ and __.gitignore__ as well as any __passwords__ or __secrets__). This can save some transfer time in certain cases, such as when there is a lot of test data or source files.

In order to run an apples-to-apples comparison when testing, do __NOT__ do either an __apt-get update__ or an __apk --update__. The reason for this is that these will make subtle changes to the operating systems that may well invalidate any test comparisons.

Consider specifying __USER nobody__ unless __root__ access is absolutely required.

## ADD versus COPY

ADD will uncompress certain files (.tar, .tar.gz, .tar.bz2) if only a destination directory is specified. It will __not__ uncompress certain other files (.zip). If the intent is to include a compressed file, make sure to specify a destination __name__ and directory.

COPY will simply copy the file, without uncompressing it (usually this is what's wanted).

If unsure, use COPY.

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

#### Layer information about the Docker image:

	$ docker history ...

	$ docker history --no-trunc ...

#### Docker volumes:

	$ docker volume ls

	$ docker volume prune
