# Fraise 
## 一个简单的 HTTP/HTTPS 负载均衡器（加权轮询 & 加权随机）

***

### 使用方法
* 填写配置文件 `application.yaml`
* 填写本地服务节点列表文件 `nodeList.json`
* 启动项目

***

### 实现功能
#### 负载均衡转发
* 将发送至该负载均衡器的HTTP/HTTPS请求按照指定策略，转发给本地服务节点列表 `nodeList.json` 文件中的一个节点。
* 例：负载均衡器服务地址及端口为 `127.0.0.1:443` ，前端访问 `https://127.0.0.1:443/login` ，该访问请求将被转发到 `http://127.0.0.1:8X/login`。
* `nodeList.json` 文件说明
```json
[
	{
		"url": "http://127.0.0.1:80",
		"weight": 1
	},
	{
		"url": "http://127.0.0.1:81",
		"weight": 2
	},
	{
		"url": "http://127.0.0.1:82",
		"weight": 3
	}
]
```
* url: 节点服务访问路径
* weight：节点权重（权重越大，分配给该节点的请求就越多）

#### 动态更新负载均衡器的本地节点列表
* 通过控制台接口更新，见下文


***

### 控制台接口

#### 更新负载均衡器的本地服务节点列表

##### 接口URL（接口URL可在配置文件 `application.yaml` 中更改）
> https://127.0.0.1/console/api/set

##### 请求方式
> POST

##### Content-Type
> form-data

##### 请求Body参数
| 参数名      | 示例值                                                                                                                                      | 参数类型 | 是否必填 | 参数描述              |
|----------|------------------------------------------------------------------------------------------------------------------------------------------|------|------|-------------------|
| nodeList | [{"url": "http://127.0.0.1:80","weight": 1},	{"url": "http://127.0.0.1:81",	"weight": 2	},	{"url": "http://127.0.0.1:82",	"weight": 3	}] | Text | 是    | 服务节点列表（JSON格式字符串） |

#### 获取负载均衡器的本地服务节点列表

##### 接口URL（接口URL可在配置文件 `application.yaml` 中更改）
> https://127.0.0.1/console/api/get

##### 请求方式
> POST

##### Content-Type
> form-data

***

### 打包方式
* 终端运行 `go build`
* 打包成Linux可执行文件 
```
//终端运行
set GOARCH=amd64
set GOOS=linux
go build main.go
```

***

### 部署方式

* 部署文件夹
  * fraise 或 fraise.exe `打包好的执行文件`
  * application.yaml `配置文件`
  * nodeList.json `服务节点列表文件`
  * xxx.key `SSL证书文件（nginx），开启https时使用`
  * xxx.pem `SSL证书文件（nginx），开启https时使用`

Linux后台运行项目：cd到部署文件夹，输入 `setsid ./fraise`
