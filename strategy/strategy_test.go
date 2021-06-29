package strategy

import "testing"

func TestAlipay_Pay(t *testing.T) {
	alipay := &Alipay{}
	payment := NewPayment("Tom", 100, alipay)
	payment.Pay()
	// output: Pay $100 to Tom by Alipay
}

func TestWechat_Pay(t *testing.T) {
	wechat := &Wechat{}
	payment := NewPayment("Alice", 99, wechat)
	payment.Pay()
	// output: Pay $99 to Alice by wechat
}
