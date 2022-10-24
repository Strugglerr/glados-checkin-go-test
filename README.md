# Glados-CheckIn

* GLADOS Automatic Check-in program ( Golang Version )
* Version: 1.2
* Update time：2022.10.25
* Author: [xiabee](https://github.com/xiabee)



## 更新日志

* 暂时关闭 `Git Action` 自动签到功能，防止仓库被封禁（如需打开请查看 `签到功能` 板块）
* 增加 [pushplus](https://pushplus.plus/) 自动推送功能



## 签到功能

- 基于 [Github Actions](https://github.com/features/actions)

* 在 `push` 和 `pr` 的时候自动执行
* 基于`Github Actions` 的签到功能有被`Github` 封禁仓库的风险，因此暂时关闭了每日自动签到
* 如需开启每日自动执行，将`.github/workflows/checkin.yml` 中的 `schedule` 字段取消注释

```yml
on:
  push:
    branches: [ "main" ]
  pull_request:
    branches: [ "main" ]
  schedule:
    - cron: 01 16 * * *
    # UTC: 16:01
    # CST: 00:01
    # Check in every day
```





## 关于 Glados

* 项目地址：https://github.com/glados-network/GLaDOS
* 控制台地址：https://glados.rocks/console
* 注册教程：（待完善）



## 使用方式

### 1. Fork 本仓库

* 直接找到右上角的 `Fork`，然后就会在个人仓库里出现一个 `Fork` 版

![image.png](https://tva1.sinaimg.cn/large/0084b03xgy1h7f37ytaiyj310m0fa48o.jpg)



### 2. 设置 fork 仓库的 secret

*  `Settings` ➡️ `Security` ➡️ `Secrets` ➡️ `Actions` ➡️`New repository secret`
* 创建一个名为 `COOKIE` 的`secret`，用于登录`Glados`，其值为 [Glados 控制台](https://glados.rocks/console) 登录后的 cookie；具体获取方式见[附录](#Cookie 获取方式)
*  创建一个名为 `TOKEN` 的 `secret`，用于登录 `pushplus`，获取方法为进入其[官网](https://pushplus.plus/push1.html)，关注公众号后获取
   *  注：不需要推送功能的旁友可以不设置 `TOKEN` ，但是必须要设置 `COOKIE`，否则无法签到


<img src="https://tva1.sinaimg.cn/large/0084b03xgy1h7guslltcmj31ri13sqm8.jpg" alt="image.png" style="zoom:67%;" />



<img src="https://tva1.sinaimg.cn/large/0084b03xgy1h7f46828goj31960ocafl.jpg" alt="image.png" style="zoom:67%;" />



### 3. 开启 Git Action

<img src="https://tva1.sinaimg.cn/large/0084b03xgy1h7f47zq532j322e0xchdt.jpg" alt="image.png" style="zoom:67%;" />

<img src="https://tva1.sinaimg.cn/large/0084b03xgy1h7f48qvsvlj32na10ah15.jpg" alt="image.png" style="zoom: 67%;" />

* 在 `push` 和 `pull request`的时候会出发执行
* 开启每日自动执行：将`.github/workflows/checkin.yml` 中的 `schedule` 字段取消注释

```yml
on:
  push:
    branches: [ "main" ]
  pull_request:
    branches: [ "main" ]
  schedule:
    - cron: 01 16 * * *
    # UTC: 16:01
    # CST: 00:01
    # Check in every day
```





## 附录

### Cookie 获取方式

* 登录[控制台](https://glados.rocks/console)之后，`F12`打开浏览器的开发者工具
* 找到 `Network` / `网络`，选中 `console` 那个包
* 在请求头中找到 `cookie` 字段

<img src="https://tva1.sinaimg.cn/large/0084b03xgy1h7f4k2lqajj31vo0zk1iq.jpg" alt="image.png" style="zoom:67%;" />



### TOKEN 获取方式

进入[官网](https://pushplus.plus/)，绑定公众号查看

<img src="https://tva1.sinaimg.cn/large/0084b03xgy1h7guuctbdvj31sa0gyal6.jpg" alt="image.png" style="zoom:67%;" />



### 执行结果

在 `Actions` 中查看

<img src="https://tva1.sinaimg.cn/large/0084b03xgy1h7f4mom4orj31o816u48v.jpg" alt="image.png" style="zoom:67%;" />



推送结果在微信公众号中查看（如果关注 [pushplus](https://www.pushplus.plus/) 公众号且配置正确的话）

![image.png](https://tva1.sinaimg.cn/large/0084b03xgy1h7guw7x1n8j30xg0jo78b.jpg)



### BUG 处理

* 待更新
