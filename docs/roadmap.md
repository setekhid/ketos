Roadmap
====

The core:

* Research how does docker privileged mode work, if it is possible to use chroot without privileged flag in a container?
	* Command chroot does need privileged mode to mount /dev folder, keeping researching command [fakechroot](https://github.com/dex4er/fakechroot) and [proot](https://github.com/proot-me/PRoot).
	* It seems that [droot](https://github.com/yuuki/droot) can chroot to a docker image to execute applications.
	* [cgroup](https://en.wikipedia.org/wiki/Cgroups) + [namespace](https://en.wikipedia.org/wiki/Namespace) -> [LXC](https://en.wikipedia.org/wiki/LXC) -> [libcontainer](https://github.com/opencontainers/runc/tree/master/libcontainer)
	* This [article](http://hansenms.github.io/devops/tips/2016/02/22/chroot-from-docker.html) contains how to tar entire filesystem and chroot into it. Also contains a link to a script to copy some nessary files into chroot root..
	* [Imitate](https://ericchiang.github.io/post/containers-from-scratch/) a docker runtime with chroot
	* Also [fakeroot](http://freecode.com/projects/fakeroot) and [fakeroot-ng](https://fakeroot-ng.lingnu.com/) can pretend we are root user.

* How to mount docker image to a folder as container's rootfs?
	* See below docker image spec.

Download base image and publish derived image to registry:

* How to tar a folder storing docker layers to docker image?
	* There are two version v1 & v2 of docker image format, still no clue why v1 still exists, compares about two is [here](https://www.slideshare.net/Docker/docker-registry-v2).
	* Version 2 manifest file spec is storing [here](https://github.com/docker/distribution/tree/master/docs/spec), with name `manifest-v2-1.md` and `manifest-v2-2.md`. The version 1 spec is [here](https://github.com/moby/moby/tree/master/image/spec).
		* The docker image metadata seems marshal from [code `image.go`](https://github.com/moby/moby/blob/master/image/image.go)
* Push image tar to docker registry.
	* There is a project named [skopeo](https://git.io/vdcw6) can convert between docker image and OCI image, and seems can push to docker registry, see [this issue comment](https://git.io/vdc6g)
		* Project skopeo using library [containers/image](https://github.com/containers/image) to manager registry communication.
		* Project [containers/storage](https://github.com/containers/storage) can manage image layers.

* Parsing Dockerfile, may accord to:
	* [docker/docker/builder/dockerfile/parser](https://github.com/moby/moby/tree/master/builder/dockerfile/parser)
	* [grammarly/rocker/src/parser](https://github.com/grammarly/rocker/tree/master/src/parser)
	* [jlhawn/dockramp/build/parser](https://github.com/jlhawn/dockramp/tree/master/build/parser)

* It seems [vx32](github.com/0intro/vx32) is another chroot solution
	* [here](github.com/majek/vx32example) is an example
	* [paper](https://swtch.com/~rsc/papers/vx32-usenix2008.pdf) and [some talks](https://news.ycombinator.com/item?id=12620205)
