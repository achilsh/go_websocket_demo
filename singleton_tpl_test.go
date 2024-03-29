package utils

import (
	"sync"
	"testing"
)

type demoSingletonGeneral struct {
	x int
	y string
	z float64
}

func (d *demoSingletonGeneral) setX(x int) {
	d.x = x
}
func (d *demoSingletonGeneral) getX() int {
	return d.x
}
func (d *demoSingletonGeneral) setY(y string) {
	d.y = y
}

// 一般定义一个全局对象，表示该资源的单例创建函数，后面就可以使用该全局变量来创建一个单例对象。
// 如果想要通过单例工厂函数获取对象，只需像 GetDemoHandle()即可
// var GetDemoHandle = SingletonFactoryImpl[demoSingletonGeneral](SingletonConstructor[demoSingletonGeneral], nil)
var GetDemoHandle = SingletonFactory[demoSingletonGeneral](nil)

//	var getDemoH2 = SingletonFactoryImpl[demoSingletonGeneral](SingletonConstructor[demoSingletonGeneral], func(v *demoSingletonGeneral) {
//		v.x = 1000
//	})
var getDemoH2 = SingletonFactory(func(v *demoSingletonGeneral) {
	v.x = 1000
})

var getDemoH3 = SingletonFactory(func(v *demoSingletonGeneral) {
	v.y = "this is demo."
})

func TestSingletonGeneral(t *testing.T) {
	{
		var wg sync.WaitGroup
		var n int = 1000
		wg.Add(n)
		var objPtrArr []*demoSingletonGeneral = make([]*demoSingletonGeneral, 0, n)
		var lk sync.Mutex

		for i := 0; i < n; i++ {
			go func() {
				defer wg.Done()
				ptr := GetDemoHandle()
				{
					lk.Lock()
					defer lk.Unlock()
					ptr.setX(10)
					objPtrArr = append(objPtrArr, ptr)
				}
			}()
		}
		wg.Wait()

		var cmpPtr *demoSingletonGeneral = nil
		for i, ptr := range objPtrArr {
			if ptr.getX() != 10 {
				t.Errorf("value is not 10")
			}

			if i == 0 {
				cmpPtr = ptr
				continue
			}
			if cmpPtr != ptr {
				t.Errorf("fail, singleton design error.")
			}
		}
	}
	var ptr1 *demoSingletonGeneral = nil
	{
		ptr1 = GetDemoHandle()
		t.Logf("other scope demo handle ptr: %p", ptr1)
	}

	var ptr2 *demoSingletonGeneral = nil
	{
		ptr2 = GetDemoHandle()
		t.Logf("other scope demo handle ptr: %p", ptr2)
	}
}

func TestSingletonGeneralWithInitFunc(t *testing.T) {
	h := getDemoH2()
	t.Logf("x: %v", h.getX())

}
