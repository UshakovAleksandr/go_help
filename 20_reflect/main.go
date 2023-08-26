package main

import (
	"fmt"
	"math"
	"reflect"
)

// определение типа
func isInt(value interface{}) bool {
	return reflect.TypeOf(value).Kind() == reflect.Int

	//fmt.Println(isInt(123))
	//fmt.Println(isInt("123"))
	//fmt.Println(isInt(12.3))
}

// проверка на isZero значение
func isEmpty(value interface{}) bool {
	return reflect.ValueOf(value).IsZero()

	//fmt.Println(isEmpty("aaaa"))
	//fmt.Println(isEmpty(""))
	//fmt.Println(isEmpty(111))
	//fmt.Println(isEmpty(0))
}

type User struct {
	FirstName  string
	SecondName string
	Age        int
}

// работа с полями структуры
func allFields(value interface{}) {
	// определения типа структуры
	fmt.Println(reflect.TypeOf(value).Elem().Kind())
	fmt.Println(reflect.TypeOf(value).Elem().Kind() == reflect.Struct)
	// получаем значение рефлекстного типа от указателя переменной для дальнейшей работы с ним через рефлексию
	elem := reflect.ValueOf(value).Elem()
	// значение всей структуры
	fmt.Println(elem)
	// количество полей в структуре
	fmt.Println(elem.NumField())

	for i := 0; i < elem.NumField(); i++ {
		// значения поля структуры
		field := elem.Field(i)
		fmt.Println(field)
		// тип поля структуры
		typeName := field.Type().Name()
		fmt.Println(typeName)
		// навание поля структуры
		f := elem.Type().Field(i).Name
		fmt.Println(f)
	}
}

// проверить наличие и получить значения поля из структуры
func getValue(structure interface{}, key string) (interface{}, error) {
	var result interface{}
	// проверяем структура ли это
	if reflect.TypeOf(structure).Elem().Kind() == reflect.Struct {
		// кладем в elem структуру
		elem := reflect.ValueOf(structure).Elem()
		// достаем поле по ключу (key - название поля)
		field := elem.FieldByName(key)
		// если поле существует
		if field.IsValid() {
			result = field.Interface()
			return result, nil
		}
	}
	// отдаем значение поля
	return nil, fmt.Errorf("поле '%s' отсутствует в структуре '%v'", key, reflect.TypeOf(structure).String())
}

// установить значения поля структуры
func setValue(structure interface{}, key string, value interface{}) error {
	// проверяем структура ли это
	if reflect.TypeOf(structure).Elem().Kind() == reflect.Struct {
		// кладем в elem структуру
		elem := reflect.ValueOf(structure).Elem()
		//fmt.Println(elem) // {Sam Diy 20}
		// достаем поле по ключу (key - название поля)
		field := elem.FieldByName(key)
		//fmt.Println(field) 20
		// проверка приватностия и наличия поля
		if field.CanSet() && field.IsValid() {
			// записываем значение
			field.Set(reflect.ValueOf(value))
			return nil
		}
	}
	return fmt.Errorf("поле '%s' отсутствует в структуре '%v'", key, reflect.TypeOf(structure).String())
}

const tagName = "validate"

func Validate(structure interface{}) {
	elem := reflect.ValueOf(structure).Elem()

	for i := 0; i < elem.NumField(); i++ {
		fieldName := elem.Type().Field(i).Name
		fieldValue := elem.Field(i)
		fieldTag := elem.Type().Field(i).Tag.Get(tagName)
		fmt.Println(fieldName, fieldValue, fieldTag)
	}
}

func main() {
	user := &User{
		FirstName:  "Sam",
		SecondName: "Diy",
		Age:        20,
	}

	res, err := getValue(user, "FirstName")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)

	if err := setValue(user, "Age", 30); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(user)
}

// FloatRoundUpToTwoChars - функция округления float32 и float64 полей структуры до двух знаков.
// Использована встроенная библиотека "reflect", без знания библиотеки КОД НЕ МЕНЯТЬ.
func FloatRoundUpToTwoChars(resultFields interface{}, precision int) error { //nolint:funlen
	if reflect.TypeOf(resultFields).Elem().Kind() == reflect.Struct {
		elem := reflect.ValueOf(resultFields).Elem()

		for i := 0; i < elem.NumField(); i++ {
			field := elem.Field(i)

			switch field.Kind() {
			case reflect.Float64, reflect.Float32:
				if field.CanSet() && field.IsValid() {
					if field.Kind() == reflect.Float64 {
						field.Set(reflect.ValueOf(math.Ceil(field.Float()*(math.Pow10(precision))) / math.Pow10(precision)))
					} else {
						field.Set(reflect.ValueOf(float32(math.Ceil(field.Float()*(math.Pow10(precision))) / math.Pow10(precision))))
					}
				} else {
					return fmt.Errorf("convert error on field '%v'. May be field is private", elem.Type().Field(i).Name)
				}
			case reflect.Array:
				switch field.Index(0).Kind() {
				case reflect.Float64:
					if field.CanSet() && field.IsValid() {
						for j := 0; j < field.Len(); j++ {
							field.Index(j).Set(reflect.ValueOf(math.Ceil(field.Index(j).Float()*(math.Pow10(precision))) / math.Pow10(precision)))
						}
					} else {
						return fmt.Errorf("convert error on field '%v'. May be field is private", elem.Type().Field(i).Name)
					}
				case reflect.Float32:
					if field.CanSet() && field.IsValid() {
						for j := 0; j < field.Len(); j++ {
							field.Index(j).Set(reflect.ValueOf(float32(math.Ceil(field.Index(j).Float()*(math.Pow10(precision))) / math.Pow10(precision))))
						}
					} else {
						return fmt.Errorf("convert error on field '%v'. May be field is private", elem.Type().Field(i).Name)
					}
				}
			case reflect.Slice:
				if field.Len() > 0 {
					switch field.Index(0).Kind() {
					case reflect.Float64:
						if field.CanSet() && field.IsValid() {
							for j := 0; j < field.Len(); j++ {
								field.Index(j).Set(reflect.ValueOf(math.Ceil(field.Index(j).Float()*(math.Pow10(precision))) / math.Pow10(precision)))
							}
						} else {
							return fmt.Errorf("convert error on field '%v'. May be field is private", elem.Type().Field(i).Name)
						}
					case reflect.Float32:
						if field.CanSet() && field.IsValid() {
							for j := 0; j < field.Len(); j++ {
								field.Index(j).Set(reflect.ValueOf(float32(math.Ceil(field.Index(j).Float()*(math.Pow10(precision))) / math.Pow10(precision))))
							}
						} else {
							return fmt.Errorf("convert error on field '%v'. May be field is private", elem.Type().Field(i).Name)
						}
					}
				}
			}
		}
	}

	return nil
}
