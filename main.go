package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
	"yao-plugin-go/pluginapi"

	"github.com/hashicorp/go-hclog"
	_ "github.com/yaoapp/gou/encoding"
	"github.com/yaoapp/gou/process"
	"github.com/yaoapp/kun/grpc"
	_ "github.com/yaoapp/yao/aigc"
	"github.com/yaoapp/yao/cmd"
	"github.com/yaoapp/yao/config"
	_ "github.com/yaoapp/yao/crypto"
	_ "github.com/yaoapp/yao/helper"
	_ "github.com/yaoapp/yao/openai"
	"github.com/yaoapp/yao/utils"
	_ "github.com/yaoapp/yao/wework"
)

// // 主程序
// func main() {
// 	utils.Init()
// 	cmd.Boot()
// }

// 定义插件类型，包含grpc.Plugin
type DemoPlugin struct{ grpc.Plugin }

// 设置插件日志到单独的文件
func (demo *DemoPlugin) setLogFile() {
	var output io.Writer = os.Stdout
	//开启日志
	logroot := os.Getenv("PLUGIN_HTTPX_CLIENT")
	if logroot == "" {
		logroot = "./logs"
	}
	if logroot != "" {
		logfile, err := os.Create(path.Join(logroot, "plugin.log"))
		if err == nil {
			output = logfile
		}
	}
	demo.Plugin.SetLogger(output, grpc.Trace)
}

// 插件执行需要实现的方法
// 参数name是在调用插件时的方法名，比如调用插件demo的Hello方法是的规则是plugins.demo.Hello时。
//
// 注意：name会自动的变成小写
//
// args参数是一个数组，需要在插件中自行解析。判断它的长度与类型，再转入具体的go类型。
func (demo *DemoPlugin) Exec(name string, args ...interface{}) (*grpc.Response, error) {

	demo.Logger.Log(hclog.Trace, "plugin method called", name)
	demo.Logger.Log(hclog.Trace, "args", args)

	//输出值支持的类型：map/interface/string/integer,int/float,double/array,slice
	var out any
	switch name {
	case "post":
		if len(args) < 1 {
			out = pluginapi.Response{Status: 400, Message: "参数不足，需要一个参数"}
			break
		}
		var process = process.New("http.post", args...)
		res, err := process.Exec()
		if err != nil {
			out = pluginapi.Response{Status: 400, Message: err.Error()}
		} else {
			out = res
		}
	case "script":
		if len(args) < 1 {
			out = pluginapi.Response{Status: 400, Message: "参数不足，需要一个参数"}
			break
		}
		scriptName := ""
		ok := false
		if scriptName, ok = args[0].(string); ok {
			args = args[1:]
		}
		var process = process.New(scriptName, args...)
		res, err := process.Exec()
		if err != nil {
			out = pluginapi.Response{Status: 400, Message: err.Error()}
		} else {
			out = pluginapi.Response{Status: 200, Data: res}
		}
	default:
		out = pluginapi.Response{Status: 400, Message: fmt.Sprintf("%s不支持", name)}
	}

	//所有输出需要转换成bytes
	bytes, err := json.Marshal(out)
	if err != nil {
		return nil, err
	}
	//支持的类型：map/interface/string/integer,int/float,double/array,slice
	return &grpc.Response{Bytes: bytes, Type: "map"}, nil
}

var myPlugin *DemoPlugin

// 生成插件时函数名修改成main
func main() {
	utils.Init()//初始化处理器，同时可以增加自定义的处理器。
	cmd.Boot() //加载配置文件
	//插件初始化
	myPlugin = &DemoPlugin{}
	myPlugin.setLogFile()
	
	// 自定义yao应用的加载过程
	// load the application engine
	CustomLoad(config.Conf, myPlugin.Logger)
	grpc.Serve(myPlugin)
}
