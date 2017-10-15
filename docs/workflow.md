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
