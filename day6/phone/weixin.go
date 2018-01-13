package main

import (
	"fmt"
)

type WeChatPay struct {

}


func (w *WeChatPay) pay(user_id int64, money float32) error {
	fmt.Println("1. 连接到微信支付的服务器")
	fmt.Println("2. 找到对应的用户")
	fmt.Println("3. 检查余额是否充足")
	fmt.Println("4. 扣钱")
	fmt.Println("5. 返回支付是否成功")

	return nil
}