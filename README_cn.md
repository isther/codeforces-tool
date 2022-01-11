# codeforces-tool

**这是一个用于 codefrces 比赛的工具.**

<div align="center">
<a href="https://github.com/isther/codeforces-tool"><img src="https://img.shields.io/github/repo-size/isther/codeforces-tool?style=flat-square&label=Repo" alt="GitHub repository size"/></a>
<a href="https://github.com/isther/codeforces-tool/blob/master/LICENSE"><img src="https://img.shields.io/github/license/isther/codeforces-tool?style=flat-square&logo=GNU&label=License" alt="License" /></a>
</div>

![use](./image/use.gif)

## 安装

- 从[releases](https://github.com/isther/codeforces-tool/releases/)下载最新的二进制文件。
- 或者下载源代码自行编译。

## 特性

- 支持模板

  - 添加模板
  - 删除模板
  - 设置默认模板
  - 基于模板生成源文件

- 测试样例

  - 自动下载问题样例
  - 测试本地样例
  - 添加样例

- 提交问题

- 获取个人当前比赛情况

- 获取比赛时间表

祝你好运！

## 使用

### 登录

![login](./image/login.png)

### 模板

首先，你应该创建一个模板文件，你可以在模板文件中插入一些占位符，当生成源文件是，cf-tool 将会替你解析这些占位符，这是解析规则：

`$%U%$`: Handle (e.g. ther)

`$%Y%$`: Year (e.g. 2021)

`$%M%$`: Month (e.g. 12)

`$%D%$`: Day (e.g. 11)

`$%h%$`: Hour (e.g. 13)

`$%m%$`: Minute (e.g. 30)

`$%s%$`: Second (e.g. 05)

### 脚本

当你使用命令`cf test`时，模板将一次运行这三个脚本：

- before_script: 主要是编译源码，如果你使用脚本行语言，这步可以跳过。

- script: 运行你的程序, 这步必须有。

- after_script: 删除编译生成的程序，也可以不删除，留空即可。

There are also placeholders here:

`$%full%$`: 源文件名 (e.g. "a.cpp")

`$%file%$`: 运行文件名 (e.g. "a")

`$%rand%$`: 8 位随机字符串 (including "a-z" "0-9")

**Please configure the script carefully**

### 获取比赛信息

- `cf race [contest]`: 初始化比赛并获取比赛的样例。

- `cf list`: 获取当前比赛所有问题的状态。

- `cf skd`: 获取比赛的时间表。

### 测试问题样例

- `cf test`: 测试当前问题的样例。

### 提交问题代码

- `cf submit`: 提交当前问题。
