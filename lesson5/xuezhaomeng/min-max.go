package main

import "fmt"

func sum(args ...int) (mn, mx int) {

	mx = args[0]
	mn = args[0]
	for i := 0; i < len(args); i++ {
		if args[i] < mn {
			mn = args[i]

		}
		if args[i] > mx {

			mx = args[i]

		}

	}

	return
}

func main() {

	fmt.Println(sum(1, 2, 3))

}
