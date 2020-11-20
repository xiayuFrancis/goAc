package main

import (
	"fmt"
	"strconv"
)

type Good interface {
	settleAccount() int
	orderInfo() string
}

type Phone struct {
	name     string
	quantity int
	price    int
}

func (phone Phone) settleAccount() int {
	return phone.price * phone.quantity
}

func (phone Phone) orderInfo() string {
	return "您要购买" + strconv.Itoa(phone.quantity) +"个"+phone.name +"计: " +strconv.Itoa(phone.settleAccount())+
		"元"
}

type FreeGift struct {
	name string
	quantity int
	price int
}

func (gift FreeGift) settleAccount() int {
	return 0
}
func (gift FreeGift) orderInfo() string{
	return "您要购买" + strconv.Itoa(gift.quantity)+ "个" +
		gift.name + "计：" + strconv.Itoa(gift.settleAccount()) + "元"
}

func calculateAllPrice(goods []Good) int {
	var allPrice int
	for _, good := range goods{
		fmt.Println(good.orderInfo())
		allPrice += good.settleAccount()
	}
	return  allPrice
}

func main() {
	ip := Phone{
		name:     "iphone",
		quantity: 1,
		price:    8000,
	}

	earphones := FreeGift{
		name:     "ear",
		quantity: 1,
		price:    200,
	}

	goods := []Good{ip, earphones}
	allPrice := calculateAllPrice(goods)
	fmt.Printf("该订单需支付%d元", allPrice)
}
