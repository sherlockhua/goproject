//计算两个大数求和
package main

import (
	"fmt"
	"strconv"
)

func judge(m string, n string) (string, string) {
	//judge方法通过比较m字符串和n字符串的长度大小，从而将长字符串都赋值给max,短的字符串都赋值给min
	var max, min string
	if len(m) > len(n) {
		max = m
		min = n
	} else {
		max = n
		min = m
	}
	return max, min
}

func sum(max string, min string) {
	//这里定义sum 用于表示最后的总和
	//extra用于每次当前位相加需要进一位的数字例如9+5 这个时候需要进1，extra用于存这个1
	var sum string
	var extra int
	for i := 1; i <= len(max); i++ {
		//res 和res2 分别存放两个数当前位的数字
		res := int(max[len(max)-i] - '0')
		var res2 int
		//因为当计算段字符串的位数之后，直接把res2 = 0,这里是根据i进行判断，
		// 如果i 已经大于len(i)就表示已经不需要计算res2了，可以直接赋值为0
		if i > len(min) {
			res2 = 0
		} else {
			res2 = int(min[len(min)-i] - '0')
		}
		//每一位的计算都是通过将res1+res2+extra求和对10求余数获取
		sum = strconv.Itoa((res+res2+extra)%10) + sum
		extra = (res2 + res + extra) / 10
	}
	//下面的情况是容易被遗漏的
	//这里考虑的情况是当出现最后一位也进一位的时候，
	// 例如99+1这个时候结果是三位，即把最后一次循环的时候的extra加在sum的前面即可
	if extra != 0 {
		fmt.Println(strconv.Itoa(extra) + sum)
	} else {
		fmt.Println(sum)
	}

}

func main() {
	var m, n string
	fmt.Scanf("%s %s", &m, &n)
	max, min := judge(m, n)
	sum(max, min)
}