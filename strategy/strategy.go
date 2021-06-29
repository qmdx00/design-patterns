package strategy

import "fmt"

type PaymentStrategy interface {
	Pay(*PaymentContext)
}

type PaymentContext struct {
	Name  string
	Money int
}

type Payment struct {
	context  *PaymentContext
	strategy PaymentStrategy
}

func NewPayment(name string, money int, strategy PaymentStrategy) *Payment {
	return &Payment{
		context: &PaymentContext{
			Name:  name,
			Money: money,
		},
		strategy: strategy,
	}
}

func (p *Payment) Pay() {
	p.strategy.Pay(p.context)
}

type Wechat struct{}

func (Wechat) Pay(ctx *PaymentContext) {
	fmt.Printf("Pay $%d to %s by wechat\n", ctx.Money, ctx.Name)
}

type Alipay struct{}

func (*Alipay) Pay(ctx *PaymentContext) {
	fmt.Printf("Pay $%d to %s by Alipay\n", ctx.Money, ctx.Name)
}
