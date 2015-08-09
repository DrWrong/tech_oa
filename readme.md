# 教学辅助系统 by drwrong<yuhangchaney@gmail.com>

# 简介
基于BS的架构的教学辅助系统，提供诸如学生分组， 学生选题 问卷调查，学生评分，统计调查等工能。
server端使用`golang`写成,使用`macaron`和`beego`框架。

# 运行

1.  安装golang的运行环境， 设置GOPATH
2.  `git clone https://github.com/DrWrong/tech_oa.git` 到"$GOPATH/src"" 目录下
3. 安装`beego`，`macaron` 等`go build` 编译安装 执行生成的二进制文件

# 数据库设计

本系统的重点在于flexbile 的数据设计

## 用户系统 - user表
## 项目管理系统  - project表
具体设计可以参考Model层
