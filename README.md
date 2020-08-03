# lcffilemanager(lcff's filemanager)
##### 2020.8.3
## lcffilemanager是我(lcff)用golang(gin)写的一个个人文件只读查看器，相当于一个文件服务器的样子（像apache的文件服务）
### 虽然叫manager但其实只读（hhh

* 运行
  #### 直接run main.go
  #### 在浏览器里访问[host]/index/?h=[盘符]:/
* 搜索
  #### 输入框中搜索当前文件夹下所有 包括子目录
* 问题
  
  1. 搜索范围太大会栈溢出
  2. 不能操作文件
* 改进建议
  
  1. server push搜索列表+分页
  2. 进行管理安全验证
* 后续
  
  1. 俺刚敲代码没几年，况且今年才刚上高中，会有很多菜的地方，请多包涵指教。
  2. 可以尝试内网穿透
  3. 实际上，我的用途是看一些河蟹的东西（~~为什么不能在电脑上直接看？？？~~），而百度网盘也有什么净网行动
  4. 同样喜欢 编程/音游/v家/gal/芦苇娘/a岛 的大哥哥~~和小姐姐~~们看我看我！
  #### 我的账号 [acfun](https://www.acfun.cn/u/14402634) [bilibili](https://space.bilibili.com/44067270) [wyy](https://music.163.com/#/user/home?id=1295123060) [个人blog](https://didodip.moe/blog) [twitter](https://twitter.com/Didodip0614)
  ###[osu(来加我好友！！)](https://osu.ppy.sh/users/15475115)
  ###### 企鹅：2967377955
  ###![architecture](doc/1.jpg)
  ##### 是的你没看错，这只是一个扩列的项目