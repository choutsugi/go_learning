## MVC模式

Model-View-Controller（模型-视图-控制器）模式。

- Model：模型代表一个存取数据的对象，带有逻辑在数据变化时更新控制器。
- View：视图代表模型包含的数据的可视化。
- Controller：控制器作用于模型和视图上（分离模型与视图），控制数据流向模型对象，并在数据变化时更新视图。

## CLD分层

`VUE/React`=>`Nginx`=>`HTTP/Thrift/gRPC`=>`Controller`=>`Logic`=>`DAO`=>`数据库`

协议处理层：支持各种协议。

Controller：服务入口，处理路由、参数校验、请求转发。

Logic/Service：逻辑（服务）层，负责处理业务逻辑。

DAO/Repository：负责数据与存储相关功能。

