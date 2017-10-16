Workflow
====

Example:

```bash
xcb pull alpine:3.6 ./alpine-3.6
cd ./alpine-3.6
cp ~/asset.txt ./
xcb chroot sh -c 'echo hello >> /hello.txt'
xcb commit
xcb push --image-name setekhid/alpine:3.6
```

First pull down a docker image from registry to a directory, and extract it into `.ketos` folder.

Then do some stuff, like copy some assets file to the working directory. You may wanna change rootfs into it to execute some program like the 4th line.

Subcommand `chroot`, will change rootfs to working directory and combine the image pulled before as the lower layer.

And then commit as a new layer.

Finally, push back to docker registry. Each steps you are working on are only modifying the working directory and basing the image which you pulled at first and storing at `.ketos`.


## 步骤

1. parse dockerfile ，将每一条解析成一个cmd
2. 使用linux的一些命令，主要是namespace，chroot 简单组装出一个容器环境，这个环境要可以运行第一步中解析出的cmd
3. 将working container commit成一个镜像， 镜像应该是可以用`docker images`读取的
