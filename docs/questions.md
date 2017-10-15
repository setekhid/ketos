questions need to be dig deeper
----

* Which ACI, OCI or docker image builder projects use `--privileged` mode, how sure, and is there any one doesn't use it?
	* rkt/rkt
	* containers/build
	* projectatomic/buildah
	* sgotti/baci
	* blablacar/dgr
* What's the current relationship and state between ACI, OCI and docker image.
	* https://github.com/appc/spec#-disclaimer-
* When dig in, lots of languages using libc to affect with kernel, but golang not. see [my ask](https://goo.gl/S4KJse), and the talk at [Hacker News](https://goo.gl/bFysCw) and the [relative article](https://goo.gl/1XmwtC)
	* Golang execution [design document](https://goo.gl/UY4vDB)
* A brief comparison between overlayfs and aufs, and the relationship to docker storage driver
	* The differences between docker overlay and overlay2 storage driver
	* The brief history of other unionfs, devicemapper, btrfs and zfs
	* How does docker unionfs read and write, the io costs.
	* [Some old documents](https://git.io/vd17o) about docker storage driver
