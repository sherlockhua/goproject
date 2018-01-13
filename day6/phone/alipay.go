package main


import (
	"fmt"
)

type AliPay struct {

}


func (a *AliPay) pay(user_id int64, money float32) error {
	fmt.Println("1. 连接到阿里的服务器")
	fmt.Println("2. 找到对应的用户")
	fmt.Println("3. 扣钱")
	fmt.Println("4. 返回支付是否成功")

	fmt.Println("5. 连接到阿里的服务器")
	fmt.Println("6. 找到对应的用户")
	fmt.Println("7. 扣钱")
	fmt.Println("8. 返回支付是否成功")
	return nil
}