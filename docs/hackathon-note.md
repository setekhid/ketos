Ketos 笔记 -- 记 Go Hackathon 2017
====

2017/11/05 by Huitse Tai

> Ketos 原名 Cross Container Builder (xcbuild)，后来想了想不够霸气，比赛结束后索性改了。

[TOC]

纯故事 (无技术细节，可略过)
----

第一次参加 Hackathon，是被 Ace Tang 抓来背锅的。比赛前填表格敲定参赛项目和团队。Ace Tang 在阿里忙双 11，就由我列了一下，几乎全是日常生活中的小工具。一波 idea 筛选就是凭借 github trending 来的，docker 排第一，Ace Tang 也在负责 docker 开发，于是我们就决定写一个 `docker build` 在 CI 平台中的替代工具。

参赛申请提交后，与组织人员短暂寒暄，掐指一算，三周技术调研及坑位填补时间。因为团队不来自一个公司，只能借助 github 同步调研进度和文档。最初项目名字叫做 xcbuild，意味 Cross Container Building，正是想让他无痛运行在 docker container 中。中间与 Tang 每隔几天都会电话同步下或者正式会晤并交换意见，确保学习方向正确。我在公司的工作主要落在调度层，Tang 在 docker 实现上比我了解得更多，中间更多次帮我审阅文档及给出指导性意见。

比赛是在上海举行的，周五兴奋了一宿，周六早上 5点爬起赶火车。时间安排大约是，周六上午废话加寒暄，顺便打听一下伙食情况。下午开始打瞌睡，晚上通宵乱嗨，周日上午准备吹牛皮，下午开始吹牛皮，晚上打电话叫人抬我回去。粗略统计，48 小时我睡了不足 6 小时，只想说“代码使我兴奋”。

需求整理
----

Ketos 原名 xcbuild，因为在技术调研过程中，我们发现要解决我们的需求，我们的工作要深入到 `docker build` 的底层上，并将整个构建过程拆解成若干工具，由用户自由组合。Ketos 一词作为希腊神话中的海中巨怪，以及鲸鱼、鲸鱼座的词根，被用来命名该组工具套件。相对于可爱友好的 docker “蠢萌海豚”，Ketos 更具侵略性、原始性和粗犷感。

该项目 idea 来源于我最近的工作，目前在公司负责 CI 平台开发。使用 dind 的 CI 平台，对于编译任务的调度需要做一些额外的工作。并且 dind 服务作为 Job 任务运行会导致 `docker build` 命令完全无缓存。这个问题在国外的各大 CI 平台暴露并不严重，譬如 Go 的依赖下载，github 等源码仓库是有 CDN 通道的，而 CI 平台的使用场景通常是与工作异步进行的，略高的耗时用户完全无法感知。但若网络仅比 56K 猫快一点的话，编译耗时问题，甚至编译因网络 timeout 而失败等问题就比较严重了。而项目依赖对于有经验的开发者而言，通常首次提交依赖描述后，常年不做改动。因此如果能将每次发起编译所需踩的坑挪到发起项目时全部踩完，虽不完美，至少用户能省点心。

> 调研过 docker registry 的实现后，不得不说 registry 的 content addressable 设计给出了可能性及不小的启发。

再来说说 geek 们的特殊癖好。如果你写过 C，不知道 C 相对于 Go 在编译时极强的定制性是否至今仍令你回味？！

总之，在一系列文档检索之后，我们发现，`docker build` 命令是一个设计完美的小白用户命令，定制性差，并且扼杀了对于 geek 们的无限的想象空间。

more is [here](http://blog.func.xyz/ketos/prez/)

技术调研
----

整个技术调研成本约为 2周 * 2人。因为我与 Tang 都对 `docker build` 的过程不甚了解，并且 48小时所能完成的核心功能不能过于复杂。所以项目目标确定后，先实行技术学习及调研，确定我们能做些什么。

更详细的 `docker build` 过程，请参考 docker 源码。逻辑上大约分为三步：

* `FROM` 语句下载 docker 镜像到本地。docker 进程的镜像管理功能提供镜像管理服务。
* `ADD` 、 `COPY` 以及 `RUN` 等命令生成一个新的镜像 layer。
* 组合所有 layer，生成新的镜像描述文件，并返回。若谈到完整的 CI 流程，还应该有一步推送新镜像到 docker registry。

总结一下，约有两个问题需要搞清楚：如何与 docker registry 通信，上传、下载并管理 docker 镜像？粗略调研了一下镜像 layer 的结构后，`RUN` 命令成为第二个难点。docker 自身有生成 container layer 并隔离诸多 namespace 运行用户代码的能力，也就是 docker 的核心功能，而我们什么都没有也没有时间去重写一个 runc。最重要的，根据我们的目标需求，我们要考虑如何能够给用户最高自由，并能帮助用户生成这层新的 layer。

docker registry 方面的调研确定了我们的工具将会以一个什么样的形式呈现给用户。runc 替代方案则确定了这整个项目的可行性及存在价值。由于 idea 是我提出来的，万一调研路上项目搁浅，锅要背好。于是由我负责后一项调研工作，Tang 负责前一项。

从确定项目目标并提交参赛申请，到提交申请后的两周多的时间里，我其实心里是完全没底的。因为 runc 替代方案的调研没有完成，我并不知道圈内是否已经有了该类工具，第二，我并不知道项目可行性。所有一切仅仅是凭借着我对于 CI 平台上应用 `docker build` 的理解驱动的。期间有无数次胆战心惊地考虑过能够令我们的项目搁浅的发现。或者 appc 到 runc 的周边已经有工具开始着手提供该类功能，或者替代 Dockerfile 中的 `RUN` 指令不可行，也就是意味着完成该项目需要对 docker 甚至 linux 内核提交 patch（意味着投入与产出失衡，人们会更愿意直接使用 `docker build` 完成工作）。

more is [here](http://blog.func.xyz/ketos/roadmap)

我很幸运。

需求分析
----

两周的技术学习后，对于项目的可行性及难度我心里已有了些底气。中间还帮助 Tang 看了下 registry 的 api 文档。我们非常幸运，虽然 registry 的几个 go client library 有诸多深坑巨壑，但是 registry 本身的设计却让人大呼 genius。content driving 简单到不能再简单。而我们也成功地找到了一个 client 包，还是由小有名气的 heroku 提供的。

结合着新学习到的技术，重新分析下我们的需求：

registry 的设计是：

```
registry-url + image-name = repository
repository + tag/reference = image
```

image 在 registry v2 以后仅由一个到两个 manifest.json 描述， 

```
manifest.json = image-layers + image-config
```

仅仅如此，这就是 docker registry 管理的所有内容。所以我们要做的就是与 registry 同步并管理 layers 和 manifest.json。

对于 `RUN` 语句，

仔细想想，我们的工具适用的场景是用户的本地 PC 编译以及基于 docker 的 CI 平台。在编译时，用户能够看到源码，甚至源码就是用户所写，或者编译过程已然运行在 container 里面，所以其实 namespace 并没有发挥多大作用。而编译工作是一个 one-off task，并非长时服务，所以我们认为 cgroup 也不是必须的。而我们对于 `RUN` 的最大需求其实只是这条语句让一段脚本运行在了一个虚假的 rootfs 里面并得到了一堆正确路径的静态文件，也即 chroot 所做的工作。

所以我们需要做的第二件事情其实是基于下载的 image 挂载或者模拟挂载为一个 union-fs，执行用户提供的命令，并找到一个技术方案让用户的命令感知到的 rootfs 为我们定义的目录。

* 最后一步是对于不论用户使用何种方式生成的新的 layer 目录，打包成一个 layer，并更新 image 的 manifest.json 包含该层 layer。

UI 及架构设计
----

临近比赛前两周，我感冒了。所以进行到这一步的时候，也就是比赛前一两天。这一步也是我们在比赛中最大的失误，直接导致了周六第一天编码过程中，工作未划清。

> UI 并非是指 GUI。

由于一层 image layer 的制作极其简单，因此我们的工具核心是 xcb 命令，即 xcbuild 的缩写，该命令包含一系列子命令帮助用户从 registry 同步 image，并生成新 image 推回 registry。具体参考 `xcb help`。

xcb 命令在这 48 小时中，主要实现 pull（拉取 image）、push（推送 image）和 commit（提交 layer 生成 image），以达到完整演示的目的。更多命令例如 cat-manifest（查看 image 的 manifest 描述）、put-manifest（推送 image 的 manifest）、pull-layer（拉取一层 layer 并解包）、push-layer（推送一层 layer）、has-layer 等等，则需要更多开发时间，为赛后添加。

在生成一层 image layer 过程中，被执行最多的，也即用户目前最依赖的指令是 `RUN` 指令。所以我们提供一个 chr 命令， 即 chroot 的缩写，该命令帮助处理 union-fs 的 layers，非侵入修改用户命令操作的文件系统根目录 rootfs 地址。具体参考 `chr help`。

more is [here](https://github.com/setekhid/ketos/blob/master/README.md)

比赛
----

比赛理论上是 48小时，然而正式开始时间是周六早 9点，提交截止时间是周日 13点，即 28小时。14点开始项目汇报评审，其实是一个每组 9分钟的作品发布会。

赛前预估过我的编码能力，约为 8小时 500行 Go 代码。理想情况下，比赛中我们团队的编码能力约为 `28 / 8 * 500 * 2 = 3500`行。预估 3000行代码也湛湛实现了我们的 prototype。所以到演示时项目完成度可能并不为 100%。我们遇到的第一个难题是判断在适时时裁剪代码以及违背一些编码原则加速实现。

若是开发工作不顺利（chr 这种偏近系统底层的开发工作），问题可能更多，例如最严重的踩到了之前未意料到的深坑巨壑。我们组两人在周六下午同时在自己负责的部分踩到了坑，消耗近 5小时 * 2人，几近灾难。周六晚加之身体疲劳已萌生退意，直至 23点才有显著进展。

评审过程的宣讲，语言组织能力也是一则考验。9分钟，要介绍出需求驱动、项目的价值或意义、项目的技术概述以及 UI。重中之重是这个项目的价值，答辩中简单的技术问题对于参赛的 geek 们几乎不具任何难度。而项目价值想要讲清楚，就要在赛前做好充分的调研工作，至少要网罗到评委们 50% 的技术视野。提前确定并共识一个清晰的价值也能够驱使我们团队在 28小时中保持热情。

整场 Hackathon，从入场开始，见到身边高手如云，心态不好很容易出事情。28小时的编码时间，感受下来很像一场头脑 Marathon。但是比赛进入评审阶段，大家互相交流技术，拓展视野时，实实地一场技术狂欢。庆幸我们坚持到最后。

值得一提，赛场上，我们拼了一桌 4人两队，分别取得了第一和第三。用陈的话说，“风水宝地”啊！

我们的作品：Ketos 原 Cross Container Builder (xcbuild)，在 [https://github.com/setekhid/ketos](https://github.com/setekhid/ketos)。

我们的团队：上海两日游，Ace Tang 和 Huitse Tai。
