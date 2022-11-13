package components

import "github.com/galaxy-kit/components/helloworld"

func init() {
	HelloWorldPt.Register("组件Start时，在控制台打印`Hello World`", helloworld.HelloWorld{})
}
