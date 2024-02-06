# readme

这是一个yao应用的golang语言模板

可以使用golang编写插件程序来扩展yao应用的功能。

同时还可以在插件中调用yao应用的其它处理器，

## 如何在yao的go插件中调用yao的处理器

首先需要在插件项目中引用yao的源代码。

```go
replace github.com/yaoapp/yao => ../../wwsheng009/yao

replace github.com/yaoapp/kun => ../../wwsheng009/kun // kun local

replace github.com/yaoapp/xun => ../../wwsheng009/xun // xun local

replace github.com/yaoapp/gou => ../../wwsheng009/gou // gou local

replace rogchap.com/v8go => ../../wwsheng009/v8go

```

然后增加一个自定义的加载器，比如这里的`load.go文件`。在加载器中，可以按需要加载yao应用的配置文件。


这里的原理是：

当yao框架启动插件程序时，程序的运行目录是跟yao应用是在同一个目录，所以在插件中也能读取.env环境变量设置文件。

在插件中没有启动api http服务器，而是启动了一个grpc服务，外部程度或是yao宿主程序,通过grpc协议来调用在插件中的功能。



## 注意

需要注意,在插件中不要加载其它yao插件，因为程序本身就是一个插件，如果再加载插件，会造成递归调用。
