package main

import "fmt"

func addBinary(a string, b string) string {
	as := []byte(a)
	bs := []byte(b)
	if len(a) < len(b) {
		as, bs = bs, as
	}
	lena := len(as) - 1
	lenb := len(bs) - 1
	flag := 0
	j := lena
	for i := lenb; i >= 0; i-- {
		count := int(bs[i]-'0') + int(as[j]-'0') + flag
		switch count {
		case 0:
		case 1:
			as[j] = '1'
			flag = 0
		case 2:
			as[j] = '0'
			flag = 1
		case 3:
			as[j] = '1'
			flag = 1
		}
		j--
	}
	for flag == 1 {
		if j < 0 {
			return "1" + string(as)
		}
		count := int(as[j]-'0') + flag
		switch count {
		case 1:
			as[j] = '1'
			flag = 0
		case 2:
			as[j] = '0'
			j--
		}
	}
	return string(as)

}

func main() {
	a := "101111"
	b := "10"
	fmt.Println(addBinary(a, b))
}
