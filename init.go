package components

import "github.com/galaxy-kit/components-go/helloworld"

func init() {
	HelloWorldPt.Register(helloworld.HelloWorld{}, "Hello World示例组件。")
}
