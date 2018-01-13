package main

 import(
	 "fmt"
 )

 func main(){
	 phone := &Phone{
		 PayMap:make(map[string]Pay, 16),
	 }
	 //phone.OpenWeChatPay()
	 //phone.OpenAliPay()
	 weChat := &WeChatPay{}
	 var tmp interface{} = weChat
	 _, ok := tmp.(Pay)
	 if ok {
		 fmt.Println("weChat is implement Pay interface")
	 	//phone.OpenPay("wechat_pay", weChat)
	 }
	 phone.OpenPay("ali_pay", &AliPay{})

	 err := phone.PayMoney("wechat_pay", 20.32)
	 if err != nil {
		 fmt.Printf("支付失败，失败原因:%v\n", err)
		 fmt.Printf("使用支付宝支付\n")
		 err = phone.PayMoney("ali_pay", 20.32)
		 if err != nil {
			fmt.Printf("支付失败，失败原因:%v\n", err)
			return
		 }
	 }

	 fmt.Println("支付成功，欢迎再次光临！")
 }