Roadmap
====

The core:

* Research how does docker privileged mode work, if it is possible to use chroot without privileged flag in a container?
	* Command chroot does need privileged mode to mount /dev folder, keeping researching command fakechroot at [github](https://git.io/vdnm1) or [debian](https://goo.gl/39cX9f), [proot](https://git.io/vdnmQ) or other solution.
	* It seems that [droot](https://git.io/vdnYQ) can chroot to a docker image to execute applications.
	* [cgroup](https://goo.gl/9iALHX) + [namespace](https://goo.gl/ki7A7s) -> [LXC](https://goo.gl/xyGEzu) -> [libcontainer](https://git.io/vdnO4)
	* This [article](https://goo.gl/ngTwR7) contains how to tar entire filesystem and chroot into it. Also contains a link to a script to copy some nessary files into chroot root. [Wikipedia of chroot](https://goo.gl/ksHZgG) and [ubuntu document](https://goo.gl/C8rtvn).
	* [Imitate](https://goo.gl/xYXJgX) a docker runtime with chroot
	* Use [fakeroot](https://goo.gl/GDVota), to pretend we are root user.

* If rkt can fit in this situation or runc?
* When docker image format will be replaced by OCI image format.

* How to mount docker image to a folder as container's rootfs?
	* See how to tar for the spec-s.
	* Some combination of outer docker daemon storage driver and inner storage driver [doesn't work](https://goo.gl/cjKAUs), research the reason. The storage driver compatible matrix is [here](https://goo.gl/Me7EFF), and other information for choosing storage driver on the same page.
	* How does `docker create` mount docker image and create a container.
	* [How to use fuse](https://git.io/vdur0) in docker container

* How to use fakechroot to imitate a union fs mounter

Download base image and publish derived image to registry:

* How to tar a mounted folder to docker image?
	* There are two version v1 & v2 of docker image format, still no clue why v1 still exists, compares about two is [here](https://goo.gl/4bKjL5).
	* Version 2 manifest file spec is storing [here](https://git.io/vdcor), with name `manifest-v2-1.md` and `manifest-v2-2.md`. The version 1 spec is [here](https://git.io/vdcod), we might wanna its golang code from its parent folder.
* Push image tar to docker registry.
	* There is a project named [skopeo](https://git.io/vdcw6) can convert between docker image and OCI image, and seems can push to docker registry, see [this comment](https://git.io/vdc6g) of an issue
		* Project skopeo using library [containers/image](https://github.com/containers/image) to manager registry communication.
		* Project [containers/storage](https://github.com/containers/storage) can manage image layers.
