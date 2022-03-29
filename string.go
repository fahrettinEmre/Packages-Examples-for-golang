package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("test", "es"))         //test içerisinde es var mı ?
	fmt.Println(strings.Count("test", "t"))             // test içerisinde kaç tane  t var ?
	fmt.Println(strings.HasPrefix("unit_test", "unit")) // unit kısmı alt çizgiden önce olduğu
	// için prefix olarak adlandırılır. Prefix var mı diye ?
	fmt.Println(strings.HasSuffix("test.rar", "rar")) // dosya uzantısı var mı ?
	fmt.Println(strings.Index("test", "s"))           // s harfi test içinde kaçıncı index ?

}
