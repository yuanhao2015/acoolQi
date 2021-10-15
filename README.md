## 平台简介
* 基于Gin的后台管理系统（权限认证基础脚手架）
* 前端采用Vue、Element UI
* 后端采用GO语言 框架 Gin

## 内置功能
#### 系统管理
1.  用户管理：用户是系统操作者，该功能主要完成系统用户配置。
2.  部门管理：配置系统组织机构（公司、部门、小组），树结构展现支持数据权限。
3.  岗位管理：配置系统用户所属担任职务。
4.  菜单管理：配置系统菜单，操作权限，按钮权限标识等。
5.  角色管理：角色菜单权限分配、设置角色按机构进行数据范围权限划分。
6.  字典管理：对系统中经常使用的一些较为固定的数据进行维护。
7.  参数管理：对系统动态配置常用参数。
8.  日志管理：对系统操作及登录日志进行管理。
#### 系统监控
1.  在线用户：在线用户展示及强退
2.  服务监控：监控服务器状态信息
3.  缓存监控：监控redis状态信息
#### 系统工具
1.  表单构建：构建基础vue
2.  代码生成：根据数据库表生成golang基础代码
3.  系统接口：swagger接口
#### 统计报表
1.  首页：dashboard信息



## 配置
项目数据库文件 /data/acoolqi-admin.sql 创建数据库导入后修改配置/config/config-*.ini


## 运行

本系统是前后端分离

后端:acoolqi-admin
```bash
# 进入后端项目目录
cd acoolqi-admin

# 启动
go run main.go

```
前端:acoolqi-ui
```bash
# 进入前端项目目录
cd acoolqi-ui

# 安装依赖
npm install

# 建议不要直接使用 cnpm 安装依赖，会有各种诡异的 bug。可以通过如下操作解决 npm 下载速度慢的问题
npm install --registry=https://registry.npm.taobao.org

# 启动服务
npm run dev
```
浏览器访问 http://localhost:80

## docker镜像构建
[docker安装使用请参考官方](https://www.docker.com/)
1. 根据情况修改Dockerfile文件
2. 在项目根目录下使用命令docker build -t <你要出的进行名>:<版本号> .


## 感谢(排名不分先后)
> Gin框架 [https://github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)
> 
> RuoYi-Vue [https://gitee.com/y_project/RuoYi-Vue](https://gitee.com/y_project/RuoYi-Vue)
>
>jwt [https://github.com/dgrijalva/jwt-go](https://github.com/dgrijalva/jwt-go)
>
>excelize [https://github.com/qax-os/excelize](https://github.com/qax-os/excelize)
>
>xorm [https://github.com/go-xorm/xorm](https://github.com/go-xorm/xorm)