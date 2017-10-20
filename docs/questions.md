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
	* the golang may wanna implement everything in golang, the runtime is in golang since 1.5 see [here](https://www.infoq.com/news/2015/08/go-1-5).
		* dig deeper if the runtime of rust-lang implemented in c or rust, and is rust based syscall or libc.
* Why need manage `/dev` folder when docker create a container.
* A brief comparison between overlayfs and aufs, and the relationship to docker storage driver
	* The differences between docker overlay and overlay2 storage driver
	* The brief history of other unionfs, devicemapper, btrfs and zfs
	* How does docker unionfs read and write, the io costs.
	* [Some old documents](https://git.io/vd17o) about docker storage driver

* Using ptrace to solve golang-skipping-libc problem, and benchmark each solution, ptrace and libc `LD_PRELOAD`.
	* [strace in 60 lines of go](https://hackernoon.com/strace-in-60-lines-of-go-b4b76e3ecd64)
	* [write yourself an strace in 70 lines of code](https://blog.nelhage.com/2010/08/write-yourself-an-strace-in-70-lines-of-code/)
	* [brief comparison between ptrace and `LD_PRELOAD`](https://fakeroot-ng.lingnu.com/index.php/Home_Page#Technical_differences_between_Fakeroot_and_Fakeroot-ng)
	* projects comparison fakeroot-ng vs fakeroot and fakechroot vs proot
	* ptrace syscall leaks some security problems. see CVE-2014-4699 and CVE-2016-5195 in [article](https://docs.docker.com/engine/security/non-events/).
	* here is another [issue of seccomp with ptrace allowed](https://blog.lizzie.io/linux-containers-in-500-loc.html#fn.51)

* How does `LD_PRELOAD` work.
	* Will it work with golang `os/exec` pkg, yes.
	* Is it under POSIX definition, no.
	* more detail in ld.so man page.

* When docker image format will be replaced by OCI image format.
* How does `docker create` mount docker image and create a container.
* Some combination of outer docker daemon storage driver and inner storage driver [doesn't work](https://goo.gl/cjKAUs), research the reason. The storage driver compatible matrix is [here](https://goo.gl/Me7EFF), and other information for choosing storage driver on the same page.

* Some guys said glibc can't be statically linked. dig the reason
	* Statically link libc will cause porting problem, if glibc is too big, uClibc or musl is a replicaing solution for some situation.
