## init
处理各服务的初始化

- router/init_router.go ： 服务路由定义（基于Hertz）
- init_config ： 定义配置文件的相关数据结构，解析配置文件
- init_db ： 初始化数据库
- init_oss ： 初始化 oss client 和 bucket