package main

import (
	"fmt"
	"time"
)

/*
   简单工厂模式
   定义一个创建父类的函数或方法，让子类决定实例化哪个类
*/

type Executor interface {
	Run()
}

type syncJob struct {
}

func (*syncJob) Run() {
	fmt.Println("syncJob run...")
}

type asyncJob struct {
}

func (*asyncJob) Run() {
	go func() {
		fmt.Println("asyncJob run...")
	}()
}

func NewExecutor(async bool) Executor {
	if async {
		return &asyncJob{}
	} else {
		return &syncJob{}
	}
}

func main() {

	NewExecutor(false).Run()

	NewExecutor(true).Run()

	time.Sleep(time.Second)
}
