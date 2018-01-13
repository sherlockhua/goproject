//获取1000以内的完数
package main

import (
	"fmt"
)

func getFactor(number int) []int {
	//获取每个数的因数，并把除了自己本身的数追加到slice中
	var allFactor = []int{}
	for i := 1; i <= (number / 2); i++ {
		if isContent(i, allFactor) {
			//因数已经存在于slice中，存在终止此次循环
			continue
		}
		if number%i == 0 {
			if i == 1 || (i == number/i) {
				// 因数1和当存在类似2*2=4这种情况是只将一个因数追加
				allFactor = append(allFactor, i)
			} else {
				// 其他情况会把两个因数同时追加
				allFactor = append(allFactor, i, number/i)
			}

		}
	}
	return allFactor
}

func isContent(num int, slice []int) bool {
	//用于判断num是否在slice中
	for _, v := range slice {
		if v == num {
			return true
		}
	}
	return false
}

func getSum(slice []int) int {
	//获取slice所有元素的和
	var sum int
	for _, v := range slice {
		sum += v
	}
	return sum
}

func getResult() []int {
	//获取1000以内的完数，并追加到slice中
	var result = []int{}
	for i := 1; i <= 1000; i++ {
		if i == 1 {
			//1 不是完数，直接放弃
			continue
		} else {
			allFactor := getFactor(i)
			sum := getSum(allFactor)
			if sum == i {
				result = append(result, i)
			}
		}

	}
	return result
}

func main() {
	result := getResult()
	fmt.Printf("1000以内所有的完数为：%v\n", result)
}
