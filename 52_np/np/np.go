package np

import (
	"math"
)

type (
	Arr  []float64
	Bool []bool
)

func (n Arr) ExpPositive() {
	for i := 0; i < len(n); i++ {
		n[i] = math.Exp(n[i])
	}
}

func (n Arr) ExpNegative() {
	for i := 0; i < len(n); i++ {
		n[i] = math.Exp(-(n[i]))
	}
}

func (n Arr) AddNum(num float64) {
	for i := 0; i < len(n); i++ {
		n[i] = n[i] + num
	}
}

func (n Arr) SubNumL(num float64) {
	for i := 0; i < len(n); i++ {
		n[i] = num - n[i]
	}
}

func (n Arr) SubNumR(num float64) {
	for i := 0; i < len(n); i++ {
		n[i] = n[i] - num
	}
}

func (n Arr) MultiNum(num float64) {
	for i := 0; i < len(n); i++ {
		n[i] = n[i] * num
	}
}

func (n Arr) DivNumL(num float64) {
	for i := 0; i < len(n); i++ {
		n[i] = num / n[i]
	}
}

func (n Arr) AddSl(np Arr) {
	for i := 0; i < len(np); i++ {
		n[i] = n[i] + np[i]
	}
}

func (n Arr) SubSl(np Arr) {
	for i := 0; i < len(np); i++ {
		n[i] = n[i] - np[i]
	}
}

func (n Arr) MultiSl(np Arr) {
	for i := 0; i < len(np); i++ {
		n[i] = n[i] * np[i]
	}
}

func (n Arr) Less(num float64) Bool {
	res := make([]bool, len(n), cap(n))

	for i := 0; i < len(n); i++ {
		res[i] = n[i] < num
	}

	return res
}

func (n Arr) WhereLess(compare, fill float64) Arr {
	res := make([]float64, len(n), cap(n))

	for idx, val := range n {
		if val < compare {
			res[idx] = fill
		} else {
			res[idx] = val
		}
	}

	return res
}

func (n Arr) WhereGte(compare, fill float64) Arr {
	res := make([]float64, len(n), cap(n))

	for idx, val := range n {
		if val > compare {
			res[idx] = fill
		} else {
			res[idx] = val
		}
	}

	return res
}

//
//	return res
//}
//
//func main() {
//	arr1 := Arr{2, 3, 4}
//	arr2 := Arr{5, 6, 7}
//	arr3 := Arr{8, 9, 10}
//
//	var res Arr
//
//	arr1.MultiSl(arr2)
//	arr1.MultiSl(arr3)
//	res = arr1
//
//	fmt.Println(res)
//}

//func main() {
//
//	arr1 := NpArr{2, 3, 4}
//	var res NpArr
//	res = arr1
//
//	var var2 float64 = 2
//	var var3 float64 = 2
//	var var4 float64 = 2
//	var var5 float64 = 26
//	var var6 float64 = 28
//	var var7 float64 = 2
//	var var8 float64 = 20
//	var var9 float64 = 2
//	var var10 float64 = 2
//
//	res.MultiNum(0.90154096)
//	res.SubNumL(-4.98221115)
//	res.SubNumR(0.39401196 * var2)
//	res.AddNum(0.06084002 * var3)
//	res.AddNum(0.21093217 * var4)
//	res.AddNum(0.34868395 * var5)
//	res.AddNum(0.11406123 * var6)
//	res.SubNumR(0.51714443 * var7)
//	res.AddNum(0.89918551 * var8)
//	res.AddNum(1.13220829 * var9)
//	res.AddNum(0.25638490 * var10)
//
//	fmt.Println(res)
//}
//
////func main() {
////	// print(2 - 3 * arr1 - 4 * var2 + 5 * var3)
////	arr1 := NpArr{2, 3, 4}
////	var res NpArr
////	res = arr1
////
////	var var2 float64 = 2
////	var var3 float64 = 2
////
////	res.MultiNum(3)
////	res.SubNumL(2)
////	res.SubNumR(4 * var2)
////	res.AddNum(5 * var3)
////
////	fmt.Println(res)
////}
//
////func main() {
////	// import numpy as np
////	//
////	// arr1 = np.array([2, 3, 4])
////	// arr2 = np.array([5, 6, 7])
////	// arr3 = np.array([8, 9, 10])
////	//
////	// print(-1 - 2 * arr1 - 3 * arr2 + 4 * arr3)
////
////	arr1 := NpArr{2, 3, 4}
////	arr2 := NpArr{5, 6, 7}
////	arr3 := NpArr{8, 9, 10}
////
////	arr3.MultiNum(4)
////	arr2.MultiNum(3)
////	arr1.MultiNum(2)
////
////	arr1.SubNum(-1)
////	arr1.SubSl(arr2)
////	arr1.AddSl(arr3)
////
////	fmt.Println(arr1)
////
////}
//
///////////////////////////////////
//////approved = np.array([2, 3, 4])
//////approved1 = np.array([5, 6, 7])
//////
//////print(-1 + np.exp(approved) - np.exp(-approved1))
////
////approved := NpArr{2, 3, 4}
////approved.ExpPositive()
////fmt.Println(approved)
////
////approved1 := NpArr{5, 6, 7}
////approved1.ExpNegative()
////fmt.Println(approved1)
////
////approved.AddNum(-1)
////fmt.Println(approved)
////
////approved.SubSl(approved1)
////fmt.Println(approved)
/////////////////////////////////////
