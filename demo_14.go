package main

import (
	"fmt"
	//"reflect"
)

func main() {

	//var m map[int]string
	//m := make(map[int]string)
	//m[0] = "东京热"
	//m[1] = "一本道"
	//m[2] = "加勒比"
	//println(m)
	//fmt.Println(m)
	//fmt.Println(len(m))
	//fmt.Println(cap(m))//  cap方法不能传map

	//n := make(map[string]string)
	//n["JiangXi"] = "NanChang"
	//n["GuangDong"] = "ShenZhen"
	//n["HuNan"] = "ChangSha"
	//n["AnHui"] = "HeFei"
	//n["XiZang"] = "XinJiang"
	//n := map[string]string{"JiangXi": "NanChang", "GuangDong": "ShenZhen", "HuNan": "ChangSha", "AnHui": "HeFei", "XiZang": "XinJiang"}

	//fmt.Println(n)
	//println(n)
	//
	//for key := range n {
	//	fmt.Println(key)
	//	//fmt.Println(value)
	//	fmt.Println(n[key])
	//}

	//nv := n["xxx"]
	//fmt.Println(nv)
	//println(nv)
	//fmt.Println("nv:", reflect.TypeOf(nv))

	//ne, ok := n["AnHui"]
	//fmt.Println(ne)
	//fmt.Println(ok)
	//fmt.Println("ne:", reflect.TypeOf(ok))
	//fmt.Println(n)
	//delete(n, "JiangXi")
	//fmt.Println(n)

	ll := map[int]string{0: "j", 8: "h", 3: "c", 9: "i", 5: "e", 7: "g", 6: "f", 4: "d", 2: "b", 1: "a"}
	kv := make([]int, 0)
	vv := make([]string, 0)
	for key, value := range ll {
		kv = append(kv, key)
		vv = append(vv, value)
	}
	//fmt.Println(kv)
	//fmt.Println(vv)
	rll := make(map[string]int, len(ll))
	for i, vvvalue := range vv {
		rll[vvvalue] = kv[i]
	}
	fmt.Println(ll)
	fmt.Println(rll)

}
