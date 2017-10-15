Talking 01
====

Business scene
----

* Public CI platform, missing `docker build` utilities.
* `docker build` image building only need rootfs namespace without namespace and cgroup utilities.

User interface, the product functions and how to use
----

* Convert Dockerfile to bash script
* Image mount to a folder, container bundle, blind mounting union fs
* `chroot` into it and run some commands

* (temporarily deprecated) Local image management, according to [containerd project](https://github.com/containerd/containerd)

* Pull image from docker registry
* Tar image layers to tar file
* (optional) Push image to docker registry

* (optional) Using [skopeo](https://github.com/projectatomic/skopeo) to convert image format, and we ignore this function

Decision
----

* push pull image from registry
* `chroot` into folder as rootfs
* Mounting docker image (blinding)

blabla
----

* docker steps: pull image to container, container run Dockerfile, commit, tar image

* review `docker create` code
	* what does it actually do, mount `/dev`, `/proc` or doing something others
	* what does `--privileged` mode actually do

* structure picture of the whole system

* a plan list

repos:
----

buildah + containerd
