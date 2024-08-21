package main

import "fmt"

/*
 装饰器模式:
  装饰器模式允许向一个现有的对象添加新的功能，同时又不改变其结构。
  这种模式通过创建一个装饰类，用来包装原有的类，并在保持原类方法签名完整性的前提下，提供了额外的功能
*/

/*
   例子：父类仅运行一次，我们需要在运行一次之后执行一下附加逻辑。
*/

type executor interface {
	Run()
}

type once struct{}

func (o *once) Run() {
	fmt.Println("once run...")
}

// 组合继承
type onceWrappers struct {
	e executor
}

//装饰器实现，我们需要基于父类已实现的功能添加新功能

func (o *onceWrappers) Run() {
	o.e.Run()
	fmt.Println("exec wrappers run...")
}

func main() {
	//这里简化创建流程，直接定义出来
	
	onceRun := &once{}
	wrappers := onceWrappers{e: onceRun}
	wrappers.Run()
}
