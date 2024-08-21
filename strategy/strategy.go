package main

import "fmt"

/*
 策略模式:
  策略模式允许在运行时动态地改变算法或策略的选择。
  通过使用策略模式，可以在运行时根据需要选择不同的算法，而不需要修改客户端代码。
*/

/*
   例子：一个数字统计，有增加和删除两种策略。
*/

// 策略上下文
type execCount struct {
	count    int
	strategy strategy
}

func newExecCount(s strategy) *execCount {
	return &execCount{
		strategy: s,
	}
}

// 设置需要使用的策略抽象
func (m *execCount) setStrategy(s strategy) {
	m.strategy = s
}

// 执行策略，无需实现为Run()

func (m *execCount) Do() {
	m.strategy.control(&m.count)
}

// 策略行为的接口
type strategy interface {
	control(*int)
}

// 减少数量的策略
type decrease struct{}

func (d *decrease) control(c *int) {
	*c--
}

// 增加数量的策略
type increase struct{}

func (i *increase) control(c *int) {
	*c++
}

func main() {
	c := newExecCount(&increase{})
	c.Do()
	fmt.Println("count ", c.count)

	c.Do()
	fmt.Println("count ", c.count)

	c.setStrategy(&decrease{})
	c.Do()
	fmt.Println("count ", c.count)

}
