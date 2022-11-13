package components

import "github.com/galaxy-kit/galaxy/define"

// HelloWorld HelloWorld组件接口
type HelloWorld interface {
	HelloWorld()
}

// HelloWorldPt 定义HelloWorld组件原型
var HelloWorldPt = define.DefineComponentPt[HelloWorld]().ComponentPt()