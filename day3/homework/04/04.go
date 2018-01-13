package main
import "fmt"
func charactortype(){
	var s2 string = "112aaaaFGG123   *&^%"
	var e,s,d,o int
	for i:=0;i<len(s2);i++{
		switch {
		case 64<s2[i] && s2[i]<91:
			e += 1
		case 96<s2[i] && s2[i]<123:
			e += 1
		case 47<s2[i] && s2[i]<58:
			d +=1
		case s2[i] == 32:
			s += 1
		default:
			o +=1
		}
	
	}
	fmt.Printf("字符串英文字符个数是：%d\n",e)
	fmt.Printf("字符串数字字符个数是：%d\n",d)
	fmt.Printf("字符串空格字符个数是：%d\n",s)
	fmt.Printf("字符串其它字符个数是：%d\n",o)


}
func main(){
	charactortype()
}