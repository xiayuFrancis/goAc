package main

import "fmt"

func main() {
	staleSlices()
}

func get() []byte {
	raw := make([]byte, 10000)
	fmt.Println(len(raw), cap(raw), &raw[0])
	res := make([]byte, 3)
	copy(res, raw[:3])
	return res
}

func staleSlices() {
	s1 := []int{1, 2, 3}
	fmt.Println(len(s1), cap(s1), s1) // 输出 3 3 [1 2 3]
	s2 := s1[1:]
	fmt.Println(len(s2), cap(s2), s2) // 输出 2 2 [2 3]
	for i := range s2 {
		s2[i] += 20
	}

	// s2的修改会影响到数组数据，s1输出新数据
	fmt.Println(s1) // 输出 [1 22 23]
	fmt.Println(s2) // 输出 [22 23]

	s2 = append(s2, 4) // append  导致了slice 扩容
	for i := range s2 {
		s2[i] += 10
	}
	// s1 的数据现在是陈旧的老数据，而s2是新数据，他们的底层数组已经不是同一个了。
	fmt.Println(s1) // 输出[1 22 23]
	fmt.Println(s2) // 输出[32 33 14]
}
