package main

import (
	"fmt"

	"52_np/np"
)

func main() {
	// arr1 = np.array([2, 3, 4])
	//
	//var2 = 2
	//var3 = 2
	//var4 = 2
	//
	//print(2 - 3 * arr1 - 4 * var2 + 5 * var3 + 6 * var4)
	arr1 := np.Arr{2, 3, 4}
	var res np.Arr
	res = arr1

	var var2 float64 = 2
	var var3 float64 = 2
	var var4 float64 = 2

	res.MultiNum(3)
	res.SubNumL(2)
	res.SubNumR(4 * var2)
	res.AddNum(5*var3 + 6*var4)

	fmt.Println(res)
}
