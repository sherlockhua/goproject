package main

/*
					  daren	hongming
0	0	0	0 	0	0	1	1
0	1	0	0 	0	0	0	0		
0	1	0	0 	0	0	1	1

1. |
0	0	0	0 	0	0	1	1
0	1	0	0 	0	0	0	0	
0   1   0   0   0   0   1   1

2. &
0	0	0	0 	0	0	1	1
0	1	0	0 	0	0	0	0	
0   0   0   0   0   0   0   0

3. ^
0	0	0	0 	0	0	1	1
0	1	0	0 	0	0	0	1
0   1   0   0   0   0   1   0

一、设置其中一位为1
用1 左移 n 与目标数做 或操作

二、判断其中一位是否等于1
用1 左移 n 与目标数做 &操作

三、清除其中一位
用1 左移 n 与目标数做 异或操作

*/
import(
	"fmt"
)

const (
	HongMing = 1 << 0
	DaRen = 1 << 1
	Vip = 1 << 2
)

type User struct {
	name string
	flag uint8
}

func set_flag(user User, isSet bool, flag uint8) User {
	if isSet == true {
		user.flag = user.flag | flag
	} else {
		user.flag = user.flag ^ flag
	}
	return user
}


func is_set_flag(user User, flag uint8) bool {
	result := user.flag & flag
	return result == 1
}

/*
func set_hongming(user User, isSet bool) User {
	if isSet == true {
		user.flag = user.flag | 1
	} else {
		user.flag = user.flag & 0
	}
	return user
}

func is_hongming(user User) bool {
	result := user.flag & 1
	return result == 1
}


func set_daren(user User, isSet bool) User {
	if isSet == true {
		user.flag = user.flag | DaRen
	} else {
		user.flag = user.flag ^ DaRen
	}
	return user
}

func is_daren(user User) bool {
	result := user.flag & DaRen
	return result == DaRen
}

func set_vip(user User, isSet bool) User {
	if isSet == true {
		user.flag = user.flag | Vip
	} else {
		user.flag = user.flag ^ Vip
	}
	return user
}

func is_vip(user User) bool {
	result := user.flag & Vip
	return result == Vip
}
*/

func weibo_test() {
	var user User
	user.name = "test01"
	user.flag = 0

	result := is_set_flag(user, HongMing)
	fmt.Printf("user is hongming:%t\n", result)

	user = set_flag(user, true, HongMing)
	result = is_set_flag(user, HongMing)
	fmt.Printf("user is hongming:%t\n", result)

	user = set_flag(user, false,HongMing )
	result = is_set_flag(user, HongMing)
	fmt.Printf("user is hongming:%t\n", result)
/*
	result = is_daren(user)
	fmt.Printf("user is daren:%t\n", result)
	user = set_daren(user, true)
	result = is_daren(user)
	fmt.Printf("user is daren:%t\n", result)

	user = set_daren(user, false)
	result = is_daren(user)
	fmt.Printf("user is daren:%t\n", result)
	

	fmt.Printf("vip:\n")
	result = is_vip(user)
	fmt.Printf("user is vip:%t\n", result)
	user = set_vip(user, true)
	result = is_vip(user)
	fmt.Printf("user is vip:%t\n", result)

	user = set_vip(user, false)
	result = is_vip(user)
	*/
	fmt.Printf("user is vip:%t\n", result)
	
}