// Package helloworld 提供HelloWorld示例组件。
package helloworld

import (
	"fmt"
	"github.com/galaxy-kit/galaxy-go/ec"
)

// HelloWorld HelloWorld示例组件，组件Start时，在控制台打印`Hello World`。
type HelloWorld struct {
	ec.ComponentBehavior
}

// Start 开始
func (comp *HelloWorld) Start() {
	comp.HelloWorld()
}

func (comp *HelloWorld) HelloWorld() {
	fmt.Printf("I'm entity[%s:%d:%d], Hello World!\n", comp.GetEntity().GetPrototype(), comp.GetEntity().GetID(), comp.GetEntity().GetSerialNo())
}
