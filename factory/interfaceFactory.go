package main

import "fmt"

/*
   抽象工厂模式
   定义一个创建父类的接口，让子类实现自己的创建接口
*/

// 定义创建父类接口
type jobFactory interface {
	CreateJob() executor
}

// 定义父类接口
type executor interface {
	Run()
}

type once struct{}

// 子类实现父类接口和抽象工厂创建方法

func (o *once) Run() {
	fmt.Println("once run...")
}

func (o *once) CreateJob() executor {
	return o
}

type while struct{}

func (w *while) Run() {
	fmt.Println("while run...")
}

func (w *while) CreateJob() executor {
	return w
}

func main() {

	onceJob := new(once)
	onceJob.CreateJob().Run()

	whileJob := new(while)
	whileJob.CreateJob().Run()

}
