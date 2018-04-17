1、配置Git
安装完git,首要任务是配置我们的信息，最重要的是用户名及邮箱，打开终端，执行以下命令。


#[root@VM_0_14_centos homework]# git config --global user.name "wgngoo"    #配置git用户名	
#[root@VM_0_14_centos homework]# git config --global user.email wg305803@163.com #配置git的邮箱


2、克隆远程版本仓库到本地
#[root@VM_0_14_centos homework]# git clone https://github.com/51reboot/golang-03-homework  #将51reboot golang3期内容同步到本地

 
3、 查看远程仓库的信息
[root@VM_0_14_centos homework]# git remote -v   
origin	https://github.com/51reboot/golang-03-homework (fetch)
origin	https://github.com/51reboot/golang-03-homework (push)

4、在wangning/homework目录下创建一个git_use_notes.go文件
#[root@VM_0_14_centos homework]# pwd
	/usr/local/gopath/golang-03-homework/lesson1/wangning/homework
#[root@VM_0_14_centos homework]# ls
		git_use_notes.go  math_pow.go

5、检查状态
[root@VM_0_14_centos homework]# git status
# On branch master
# Untracked files:
#   (use "git add <file>..." to include in what will be committed)
#
#	git_use_notes.go			#表示这个文件是新的,还没有被git跟踪,下一步需要git add 
nothing added to commit but untracked files present (use "git add" to track)



6、暂存

[root@VM_0_14_centos homework]# git add git_use_notes.go 


7、重新查看状态

# On branch master
# Changes to be committed:
#   (use "git reset HEAD <file>..." to unstage)
#
#	new file:   git_use_notes.go  #表示已经添加到git的临时库
#

8、创建一次提交

git commit -m "upload git use notes file"
这就创建了一次提交，-m "upload git use notes file"表示对这次提交的描述，建议使用有意义的描述性信息。

[root@VM_0_14_centos homework]# git commit -m "upload git use notes file"
[master c9ed19e] upload git use notes file
 1 file changed, 0 insertions(+), 0 deletions(-)
 create mode 100644 lesson1/wangning/homework/git_use_notes.go

9、上传前先从远程拉取一次最新版本代码git pull，防止远程代码已经被别人更新过


[root@VM_0_14_centos homework]# git pull 
remote: Counting objects: 45, done.
remote: Compressing objects: 100% (24/24), done.
remote: Total 45 (delta 20), reused 38 (delta 13), pack-reused 0
Unpacking objects: 100% (45/45), done.
From https://github.com/51reboot/golang-03-homework
   5ea312d..b83f73e  master     -> origin/master
Merge made by the 'recursive' strategy.
 lesson1/turi/HomeWork/README.org | 8 ++++++++
 1 file changed, 8 insertions(+)
 create mode 100644 lesson1/turi/HomeWork/README.org


10、上传本地最新代码到git远程服务器


[root@VM_0_14_centos homework]# git push origin master
Username for 'https://github.com': wgngoo
Password for 'https://wgngoo@github.com': 
Counting objects: 14, done.
Delta compression using up to 4 threads.
Compressing objects: 100% (8/8), done.
Writing objects: 100% (8/8), 784 bytes | 0 bytes/s, done.
Total 8 (delta 5), reused 0 (delta 0)
remote: Resolving deltas: 100% (5/5), completed with 3 local objects.
To https://github.com/51reboot/golang-03-homework
   b83f73e..c84424b  master -> master
