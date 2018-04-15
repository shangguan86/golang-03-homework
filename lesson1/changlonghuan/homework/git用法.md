git用法
===
基础
---
//设置自己的用户名和emmail
//--global参数是本地所有仓库都使用这个相同的配置，可以每个仓库不同。
git config --global user.name "xxx"
git config --global user.email "xxx@xxx"

//clone一个仓库到本地工作区
git clone github.com/xxx

//把文件添加到暂存区
git add xxx

//把文件提交到工作区
//-m参数是提交的说明
git commit -m "Descriptions"

//删除git仓库的文件
git rm xxx

//显示仓库当前的状态
git status

//更新当前的工作区
git pull
//提交更新到仓库
git push


进阶
---
//恢复工作区的文件
git checkout xxx
//创建分支
git checkout -b xxxx
等价于：
  git branch xxxx
  git checkout xxxx
//合并xxxx到当前分支
git merge xxxx
