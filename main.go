package main

import (
	"fmt"
	"reflect"
)

func main() {
	i := 13
	var f float64 = 12.9
	c1 := 'z'
	c2 := '김'
	fmt.Printf("value i : %d, value f : %f\n", i, f)
	fmt.Printf("%d * %f : %f\n", i, f, float64(i)*f) // conversion 묵시적 형변환을 지원하지 않음
	fmt.Println(reflect.TypeOf(f), reflect.TypeOf(i), reflect.TypeOf(c1), reflect.TypeOf(c2))
	fmt.Println(c1, c2)
	fmt.Printf("%x\n", c2) // '김'에 대한 16진수 출력 유니코드, home.unicode.org
}
