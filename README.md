# Mastodon/长毛象 点赞狂战士

可以自动给主页嘟嘟点赞。


#### 关注作者

[@meina](https://mastodon.social/web/accounts/849118)



#### 灵感

我有时会看到有些人接连发好几条嘟嘟但都没有人点赞...太苦逼🥺了，所以我一般看到都会点个赞，尽管我有时也觉得内容不十分有趣:sweat_smile:。

所以我决定开发这个程序，来给首页的朋友们**捧场！**



#### 特色

1. 自定义程序运行的间隔时间（e.g. 每 60s 获取一次首页数据（只能获取到最新的20条））：`interval`。
2. 支持所有实例：通过自行配置`homeURL`、`starURL`来实现。
3. 控制给哪些人点赞，默认给首页所有人点赞，但如果配置了`users`，将只会给部分用户点赞。



#### 如何使用

1. 在[release](https://github.com/falkaa/mastodon_star_bot/releases)下载对应平台的可执行文件。

2. 在同目录下创建配置文件`config.json`并进行自定义配置。配置文件示例参考[这个](https://github.com/falkaa/mastodon_star_bot/blob/master/config.json)。

3. > 配置文件说明：
   >
   > 1. token 如何获取：通过 Mastodon/长毛象 的 「首选项」— 「开发」—「创建新应用」— 名称随便填 — 然后选择「read/write」— 「提交」— 点击创建好的应用名称 — 「你的访问令牌」即是 token
   > 2. starURL/homeURL：如果你是`https://mastodon.social`实例的用户，那么不需要这两个字段，删掉即可，或者和示例配置保持一致也行。
   > 3. users：是你想特殊关注的用户，这些用户发的每条嘟嘟都会被点赞。由他们的`username`组成。P.S. 不是`displayName`哦。比如我的`displayName`是“CyberChina‧翠”，但我的`username`是`meina`。
   > 4. interval：程序每次拉取首页数据的间隔，每次只能获取到最新的20条数据。默认60s。

4. 运行。

   1. Windows 平台双击来运行。
   2. Linux/macOS 平台在终端（Terminal）`cd`到该目录后，通过`./starbot-osx`或者`./starbot-linux`运行。

#### 常见问题

1. Windows 平台下直接闪退
   1. 请检查是否在同目录下创建`config.json`的配置文件，并且配置文件是否有问题（参考示例配置文件）
   2. 如果你所谓实例被墙的话...那么你需要想办法让这个程序走VPN（后面有可能推出设置代理功能哦～猴年马月了...）
2. Linux、macOS 平台报错
   1. 参考抛出的错误消息，偶尔程序请求接口失败后，会直接推出程序，暂时还不知道 Golang 如何在程序中忽略这种错误，来让程序继续执行。
   2. 同上，如果实例被墙的话，那么需要配置下使命令行走代理，参考[文章](https://zhuanlan.zhihu.com/p/46973701)
3. 程序出错退出后，需要手动重新执行。（P.S. 你也考虑使用[pm2](https://pm2.keymetrics.io/)这种工具来管理你的程序，这样它将会自动重启。）
   
P.S. 不是上面的问题？那么，请到[issues](https://github.com/falkaa/mastodon_star_bot/issues)反馈，或者直接来长毛象找我[@meina](https://mastodon.social/web/accounts/849118)
