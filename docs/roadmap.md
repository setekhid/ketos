Roadmap
====

The core:

* Research how does docker privileged mode work, if it is possible to use chroot without privileged flag in a container?
	* Command chroot does need privileged mode to mount /dev folder, keeping researching command fakechroot at [github](https://git.io/vdnm1) or [debian](https://goo.gl/39cX9f), [proot](https://git.io/vdnmQ) or other solution.
	* It seems that [droot](https://git.io/vdnYQ) can chroot to a docker image to execute applications.
	* [cgroup](https://goo.gl/9iALHX) + [namespace](https://goo.gl/ki7A7s) -> [LXC](https://goo.gl/xyGEzu) -> [libcontainer](https://git.io/vdnO4)
	* This [article](https://goo.gl/ngTwR7) contains how to tar entire filesystem and chroot into it. Also contains a link to a script to copy some nessary files into chroot root. [Wikipedia of chroot](https://goo.gl/ksHZgG) and [ubuntu document](https://goo.gl/C8rtvn).

* If rkt can fit in this situation or runc?

* When docker image format will be replaced by OCI image format.

Download base image and publish derived image to registry:

* How to untar docker image and mount it as a rootfs?
* How to tar a mounted folder to docker image?
* Push image tar to docker registry.
