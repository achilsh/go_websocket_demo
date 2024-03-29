package utils

import "sync"

// SingletonFactory 实现了 创建单例对象函数的工厂模板，参数是被创建单例对象的初始化函数。
func SingletonFactory[T any](fn func(*T)) func() *T {
	o := new(sync.Once)
	var data *T = nil

	//多次调用该返回值，可以服用相同的上面变量。
	return func() *T {
		o.Do(func() {
			data = SingletonConstructor(fn)
		})
		return data
	}
}

// SingletonConstructor 构造函数的模板函数
func SingletonConstructor[T any](initFn func(*T)) *T {
	v := new(T)
	if initFn != nil {
		initFn(v)
	}
	return v
}

// SingletonFactoryImpl 封装单例，返回一个创建单例的工厂函数， 参数是创建单例的构造函数和初始化单例对象的函数。
// 业务只需要实例化具体烈性
func SingletonFactoryImpl[T any](constructor func(fn func(*T)) *T, initFn func(*T)) func() *T {
	o := new(sync.Once)
	var data *T = nil

	//多次调用该返回值，可以服用相同的上面变量。
	return func() *T {
		o.Do(func() {
			data = constructor(initFn)
		})
		return data
	}
}
