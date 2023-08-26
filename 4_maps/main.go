package main

import (
	"fmt"
)

func map1() {
	var cache1 map[string]string  // не-инициализированный словарь, nil
	cache2 := map[string]string{} // с помощью литерала, len(cache) == 0
	cache3 := map[string]string{  // литерал с первоначальным значением "one": "один",
		"two":   "два",
		"three": "три"}
	cache4 := make(map[string]string)      // тоже что и map[string]string{}
	cache5 := make(map[string]string, 100) // заранее выделить память // на 100 ключей

	fmt.Println(cache1)
	fmt.Println(cache2)
	fmt.Println(cache3)
	fmt.Println(cache4)
	fmt.Println(cache5)
}

func map2() {
	cache := make(map[string]string)
	cache["key"] = "value1"

	value := cache["key"] // получение значения
	fmt.Println(value)

	value, ok := cache["key"] // получить значение, и флаг того что ключ найден
	if ok {
		fmt.Println("ok")
	}

	_, ok = cache["key11"] // проверить наличие ключа в словаре
	if ok {
		fmt.Println("in")
	} else {
		fmt.Println("not in")
	}

	cache["key2"] = "value2" // записать значение в инициализированный(!) словарь
	fmt.Println(cache)

	delete(cache, "key2") // удалить ключ из словаря, работает всегда
	fmt.Println(cache)
}

func map3() {
	// В Go нет функций, возвращающих списки ключей и значейний словаря.
	cache := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}

	// Получить ключи
	keys := make([]string, 0, len(cache))
	for key, _ := range cache {
		keys = append(keys, key)
	}
	fmt.Println(keys)

	// Получить значения
	values := make([]string, 0, len(cache))
	for _, val := range cache {
		values = append(values, val)
	}
	fmt.Println(values)
}

func map4() {
	// очистка мапы. Через приравнивание к пустой, сборщик ее удалит
	cache := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}

	cache = map[string]string{}

	fmt.Println(cache)
}

func map5() {
	// Ключом может быть любой типа данных, для которого определена операция сравнения == : строки, числовые типы, bool каналы (chan);
	// интерфейсы;
	// указатели;
	// структуры или массивы содержащие сравнимые типы.
	// СЛАЙСЫ НЕ МОГУТ БЫТЬ КЛЮЧОМ
	//type User struct {
	//	Name string
	//	Host string
	//}
	//
	//cache := make(map[User][]Permission)
}

// не надо в цикле удалять или добавлять ключи
//func ex1() {
//	m := make(map[int]string)
//	for i := 0; i < 10; i++ {
//		m[i] = ""
//	}
//	fmt.Println(m)
//
//
//	for i := 0; i < 10; i++ {
//		//dump(m)
//		delete1If5(m)
//		//add12If3(m)
//	}
//}
//
//func dump(m map[int]string) {
//	for k := range m {
//		fmt.Print(k, " ")
//	}
//	fmt.Println()
//}
//
//func delete1If5(m map[int]string) {
//	for k := range m {
//		if k == 5 {
//			delete(m, 1)
//		}
//		fmt.Print(k, " ")
//	}
//	fmt.Println()
//	m[1] = ""
//}

func slic_map() {
	//// Для слайсов и словарей, zero value — это nil .
	//// С таким значением будут работать функции и операции читающие данные, например:
	//var seq []string             // nil
	//var cache map[string]string  // nil
	//l := len(seq)
	//c := cap(seq)
	//l := len(cache)
	//v, ok := cache[key] // "", false
	//
	//// Для слайсов будет так же работать append
	//var seq []string             // nil
	//seq = append(seq, "hello")   // []string{"hello"}
}

func map6() {
	//Вместо
	//hostUsers := make(map[string][]string)
	//for _, user := range users {
	//	if _, ok := hostUsers[user.Host]; !ok {
	//		hostUsers[user.Host] = make([]string)
	//	}
	//	hostUsers[user.Host] = append(hostUsers[user.Host], user.Name)
	//}
	//Можно
	//hostUsers := make(map[string][]string)
	//for _, user := range users {
	//	hostUsers[user.Host] = append(hostUsers[user.Host], user.Name)
	//}
}

// Уникальность
func map7() {
	type User struct {
		Id int
		Name string
	}

	users := []User{
		{
			Id: 1,
			Name: "Vasia",
		},
		{
			Id: 45,
			Name: "Petya",
		},
		{
			Id: 57,
			Name: "John",
		},
		{
			Id: 45,
			Name: "Petya",
		},
	}

	uniqueUsers := make(map[int]string, len(users))

	for _, user := range users {
		if _, ok := uniqueUsers[user.Id]; !ok {
			uniqueUsers[user.Id] = user.Name
		}
	}

	fmt.Println(uniqueUsers)
}

func map8() {
	type User struct {
		Id int
		Name string
	}

	users := []User{
		{
			Id: 1,
			Name: "Vasia",
		},
		{
			Id: 45,
			Name: "Petya",
		},
		{
			Id: 57,
			Name: "John",
		},
		{
			Id: 45,
			Name: "Petya",
		},
	}

	usersMap := make(map[int]User, len(users))

	for _, user := range users {
		if _, ok := usersMap[user.Id]; !ok {
			usersMap[user.Id] = user
		}
	}

	fmt.Println(usersMap)

	// findUser ниже
	if user, ok := usersMap[45]; ok {
		fmt.Println(user)
	}
}

//// поиск пользователя
//func findUser(id int, usersMap map[int]User) * User {
//	if user, ok := usersMap[id]; ok {
//		return &user
//	}
//	return nil
//}


func main() {

}
