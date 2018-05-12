package main

import "fmt"

func fib(n int64, caches map[int64]int64) int64 {

	if cache, ok := caches[n]; ok {
		return cache
	}
	if n <= 1 {
		return 1
	} else {
		caches[n] = fib(n-1, caches) + fib(n-2, caches)
		//fmt.Println(caches[n])
		return caches[n]
	}

}

func main() {
	fmt.Println("Cache impore,fib fast calculate.")
	cache := map[int64]int64{0: 0}
	fmt.Println(fib(50, cache))
	//	fmt.Println(fib(20, map[int64]int64{0: 0}))
	fmt.Println(fib(100, cache))
}
