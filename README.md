# Glados-CheckIn

* GLADOS Automatic Check-in program ( Golang Version )
* Version: 1.1
* Update time：2022.10.23
* Author: [xiabee](https://github.com/xiabee)



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
* 创建一个名为 `COOKIE` 的`secret`，其值为 [Glados 控制台](https://glados.rocks/console) 登录后的 cookie；具体获取方式见[附录](#Cookie 获取方式)

![image.png](https://tva1.sinaimg.cn/large/0084b03xgy1h7f3a6lzcvj32g816i7q2.jpg)



![image.png](https://tva1.sinaimg.cn/large/0084b03xgy1h7f46828goj31960ocafl.jpg)



### 3. 开启 Git Action

![image.png](https://tva1.sinaimg.cn/large/0084b03xgy1h7f47zq532j322e0xchdt.jpg)

![image.png](https://tva1.sinaimg.cn/large/0084b03xgy1h7f48qvsvlj32na10ah15.jpg)

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

![image.png](https://tva1.sinaimg.cn/large/0084b03xgy1h7f4k2lqajj31vo0zk1iq.jpg)



### 执行结果

在 `Actions` 中查看

![image.png](https://tva1.sinaimg.cn/large/0084b03xgy1h7f4mom4orj31o816u48v.jpg)



### BUG 处理

* 待更新
