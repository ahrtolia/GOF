package structural

// Author      : zhaoxiaodong
// Create      : 2022/7/19 17:06
// Description : 装饰模式

//装饰模式使用对象组合的方式动态改变或增加对象行为。
//
//Go语言借助于匿名组合和非入侵式接口可以很方便实现装饰模式。
//
//使用匿名组合，在装饰器中不必显式定义转调原对象方法。

type DecoratorComponent interface {
	Calc() int
}

type ConcreteComponent struct{}

func (*ConcreteComponent) Calc() int {
	return 0
}

type MulDecorator struct {
	DecoratorComponent
	num int
}

func WarpMulDecorator(c DecoratorComponent, num int) DecoratorComponent {
	return &MulDecorator{
		DecoratorComponent: c,
		num:                num,
	}
}

func (d *MulDecorator) Calc() int {
	return d.DecoratorComponent.Calc() * d.num
}

type AddDecorator struct {
	DecoratorComponent
	num int
}

func WarpAddDecorator(c DecoratorComponent, num int) DecoratorComponent {
	return &AddDecorator{
		DecoratorComponent: c,
		num:                num,
	}
}

func (d *AddDecorator) Calc() int {
	return d.DecoratorComponent.Calc() + d.num
}
