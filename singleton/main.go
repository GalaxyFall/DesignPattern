package singleton

import "sync"

/*
   单例模式
   一个类只能生成一个实例，提供一个全局访问点供外部获取该实例
*/

type Client struct{}

var instance *Client
var clientOnce sync.Once

func GetClient() *Client {
	clientOnce.Do(func() {
		instance = &Client{}
	})
	return instance
}
