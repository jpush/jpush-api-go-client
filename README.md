jpush-api-go-library
====================

JPush's officially supported Go-lang client library for accessing JPush APIs. 极光推送官方支持的 Go 语言版本服务器端 SDK。

对应的 REST API 文档：<http://docs.jpush.cn/display/dev/REST+API>

## 安装与使用
1. 配置好GOPATH
2. 执行go get github.com/jpush/jpush-api-go-client
3. 进入%GOPATH%\github.com\jpush\jpush-api-go-client目录
4. 编译example.go

## 最简示例代码

> func push(po *jpush.PushPayload) {
>>	appKey := "7d431e42dfa6a6d693ac2d04"
>>	masterSecret := "5e987ac6d2e04d95a9d8f0d1"

>>	jpc := jpush.NewJPushClient(appKey, masterSecret)

>>	str, err := jpc.Push(po)

>>	fmt.Println("result : ", str, err)

>}

> .......

> basic := jpush.NewPushPayload(jpush.NewPlatform(), jpush.NewAudience()).
>		AddMessage(jpush.NewMessage("hello,jpush"))

>	push(basic)

更多接口请参考示例代码，所以对象定义都遵循<http://docs.jpush.cn/display/dev/REST+API>中的定义
