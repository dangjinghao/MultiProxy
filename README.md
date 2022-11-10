# MultiProxy
使用二级目录反代程序

## 参数配置



本程序支持以下参数

> target [`scheme`]://[`domain/ip:port`]://[`nickname`],[`scheme2`]://[`domain/ip:port`]://[`nickname2`],`...`	 ——反代站点参数
>
> redirect [`bool`]	——是否重定向

### 如何传入

由于此程序仅为插件，所以传入参数的方式略微特殊。

启动程序时使用 `-add target=xxx;redirect=[bool]`启动

### 例

假设我需要使用此项目反代`http://127.0.0.1:2333`和`http://127.0.0.1:2334`，并分别设置通过`/123`和`/456`的方式来访问这两个服务（不这么做的话，默认需要访问`/127.0.0.1:2333`和`/127.0.0.1:2334`来访问服务），那么我们需要这么配置程序启动参数

>  ./MultiProxy -add target=http://127.0.0.1:2333://123,http://127.0.0.1:2334://456



有些现代的SPA应用可能无法在非根目录工作，所以我们提供了是否重定向的选项以更好支撑其工作。

使用方式为

>  ./MultiProxy -add target=http://127.0.0.1:2333://123,http://127.0.0.1:2334://456;**redirect=true**



## 编译与二次开发

本项目实质为`TheresaProxyV2`项目的插件，二次开发和编译过程

- 克隆`TheresaProxyV2`仓库
  - `git clone https://github.com/dangjinghao/TheresaProxyV2 MultiProxy`
- 删除插件文件夹(也可以保留)
  - `rm MultiProxy/plugins/*`
- 克隆本仓库并重命名为插件文件夹
  - `git clone https://github.com/dangjinghao/MultiProxy MultiProxy/plugins ` 
- 开发
- 编译整个项目

## 已知问题

- [ ] 获取`manifest.json`时浏览器不会附加cookie导致无法确定目标站点。
