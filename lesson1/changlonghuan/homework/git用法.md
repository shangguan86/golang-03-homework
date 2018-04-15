git用法
===
基础
---
//设置自己的用户名和emmil<br>
//--global参数是本地所有仓库都使用这个相同的配置，可以每个仓库不同。<br>
git config --global user.name "xxx"<br>
git config --global user.email "xxx@xxx"<br>
<br>
//clone一个仓库到本地工作区<br>
git clone github.com/xxx<br>
<br>
//把文件添加到暂存区<br>
git add xxx<br>
<br>
//把文件提交到工作区<br>
//-m参数是提交的说明<br>
git commit -m "Descriptions"<br>
<br>
//删除git仓库的文件<br>
git rm xxx<br>
<br>
//显示仓库当前的状态<br>
git status<br>
<br>
//更新当前的工作区<br>
git pull<br>
//提交更新到仓库<br>
git push<br>
<br>

进阶
---
//恢复工作区的文件<br>
git checkout xxx<br>
//创建分支<br>
git checkout -b xxxx<br>
等价于：<br>
  git branch xxxx<br>
  git checkout xxxx<br>
//合并xxxx到当前分支<br>
git merge xxxx<br>
