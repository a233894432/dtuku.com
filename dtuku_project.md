App──app  项目开发目录
        ├─config 配置文件目录
        │  ├─app.config 系统应用配置文件
        │  └─db.config 数据库配置文件
        ├─common 后端公共目录
        │  └─... 如utils等其他
        ├─middleware 后端公共中间件目录
        ├─static 前端公共目录 (url: /static)
        │  ├─tpl 公共tpl模板目录
        │  ├─js 公共js目录 (url: /static/js)
        │  ├─css 公共css目录 (url: /static/css)
        │  ├─img 公共img目录 (url: /static/img)
        │  └─plugin 公共js插件 (url: /static/plugin)
        ├─uploads 默认上传下载目录
        ├─router 源码路由配置
        │  ├─sys_router.go 系统模块路由文件
        │  ├─biz_router.go 业务模块路由文件
        ├─database 默认数据库文件存储目录
        ├─logger 运行日志输出目录
        └─main.go 应用入口文件
github.com
