package components

import "github.com/galaxy-kit/galaxy-go/define"

// HelloWorld HelloWorld组件接口
type HelloWorld interface {
	HelloWorld()
}

// HelloWorldPt HelloWorld组件原型
var HelloWorldPt = define.DefineComponentInterface[HelloWorld]().ComponentInterface()
