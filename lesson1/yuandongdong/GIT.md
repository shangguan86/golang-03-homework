# Git 学习  
## 使用Github  
1.安装配置git  
从官网下载git，安装，进行初始化设置  
>    git config --global user.name "My Name"  
>    git config --global user.email myEmail@example.com  

2.从github clone项目
新建文件夹保存项目文件，这里创建C:\Users\ydd\Desktop\go\src\golang-03-homework，cmd切换到该目录下，进行初始化  
>    git init  
>    git clone git@github.com:51reboot/golang-03-homework.git  
现在项目已经从github克隆到本地了。这是第一次拉取项目，下次可以直接执行git pull从github同步。  

3.提交作业  
创建C:\Users\ydd\Desktop\go\src\golang-03-homework\yuandongdong文件夹，在该目录下创建一个作业文件，比如test.go，保存代码，提交到github  
>    git add yuandongdong/test.go  
>    git commit -m "first commit"  
>    git push   
每次提交文件都要执行上面这三步。

4.查看github
提交作业之后，访问[golang-03-homework](https://github.com/51reboot/golang-03-homework)查看提交的代码。

