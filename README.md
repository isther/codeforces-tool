# codeforces-tool

**This is a tool for codeforces contest.**

<div align="center">

<a href="https://github.com/isther/codeforces-tool"><img src="https://img.shields.io/github/repo-size/isther/codeforces-tool?style=flat-square&label=Repo" alt="GitHub repository size"/></a>
<a href="https://github.com/isther/codeforces-tool/blob/master/LICENSE"><img src="https://img.shields.io/github/license/isther/codeforces-tool?style=flat-square&logo=GNU&label=License" alt="License" /></a>
[![Go Report Card](https://goreportcard.com/badge/github.com/28251536/codeforces-tool)](https://goreportcard.com/report/github.com/28251536/codeforces-tool)

</div>

![use](./image/use.gif)
[中文文档](https://github.com/isther/codeforces-tool/blob/master/README_cn.md)
## Installation

- Download the latest release from the [releases](https://github.com/isther/codeforces-tool/releases/) page.
- Or download the source code to compile.

## Features

- Support template

  - Add template
  - Remove template
  - Set default template
  - Generate source files based on templates

- Test sample

  - Download the problem sample to the local
  - Test local samples
  - Add sample

- Submit problem

- Get personal current contest status

- Get the schedule of the contest

Good luck in the codeforces contest!

## Usage

### Login

![login](./image/login.png)

### Template

First of all, you should create a template file, you can insert some placeholders into your template code. When generate a problem from the template, cf will replace all placeholders by following rules:

`$%U%$`: Handle (e.g. ther)

`$%Y%$`: Year (e.g. 2021)

`$%M%$`: Month (e.g. 12)

`$%D%$`: Day (e.g. 11)

`$%h%$`: Hour (e.g. 13)

`$%m%$`: Minute (e.g. 30)

`$%s%$`: Second (e.g. 05)

### Script

Template will run 3 scripts in sequence when you run "cf test":

- before_script: Mainly compile the source code, if you run a scripting language, you can skip it.

- script: Run your program, this step must have.

- after_script: Delete the compiled program, or not delete it.

There are also placeholders here:

`$%full%$`: Full name of source file (e.g. "a.cpp")

`$%file%$`: Name of source file (Excluding suffix, e.g. "a")

`$%rand%$`: Random string with 8 character (including "a-z" "0-9")

**Please configure the script carefully**

### Get contest info

- `cf race [contest]`: Initialize the contest and get the sample of the contest.

- `cf list`: Get a list of current matches and problem status

- `cf skd`: Get the schedule of the contest

### Test problem

- `cf test`: Test your code according to the sample of the current problem.

### Submit problem

- `cf submit`: Submit the source code of the current problem
