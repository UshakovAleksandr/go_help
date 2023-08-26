package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unicode/utf8"
)

func Validate(in interface{}) error {
	var elem reflect.Value
	// проверка на указатель входящей струтуры
	if reflect.ValueOf(in).Kind() == reflect.Ptr {
		elem = reflect.ValueOf(in).Elem()
	} else {
		elem = reflect.ValueOf(in)
	}

	for i := 0; i < elem.NumField(); i++ {
		// правило тэга
		tagRules := elem.Type().Field(i).Tag.Get(tagName)
		for _, tagRule := range TagRulesParse(tagRules) {
			fieldName := elem.Type().Field(i).Name
			// проверка на поддержку правила тэга
			if err := ValidTagNameCheck(tagRule[0], fieldName); err != nil {
				return err
			}
			// пропускаем тэги
			if tagRule[0] == "" {
				continue
			}
			// проверка правила "required"
			if err := RequireRuleCheck(elem.Field(i), tagRule[0], fieldName); err != nil {
				return err
			}

			if tagRule[0] != "required" {
				// проверка правила "minValue"
				if err := MinValueRuleCheck(elem.Field(i), tagRule, fieldName); err != nil {
					return err
				}
				// проверка правила "maxValue"
				if err := MaxValueRuleCheck(elem.Field(i), tagRule, fieldName); err != nil {
					return err
				}
				// проверка правила "minLen"
				if err := MinLenRuleCheck(elem.Field(i), tagRule, fieldName); err != nil {
					return err
				}
				// проверка правила "maxLen"
				if err := MaxLenRuleCheck(elem.Field(i), tagRule, fieldName); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// ValidTagNameCheck проверка на поддержку правила тэга.
func ValidTagNameCheck(tagRule, fieldName string) error {
	switch tagRule {
	case "":
		return nil
	case "required":
		return nil
	case "minValue":
		return nil
	case "maxValue":
		return nil
	case "minLen":
		return nil
	case "maxLen":
		return nil
	default:
		return fmt.Errorf("unsupported tagRule:'%s' on field:'%s'", tagRule, fieldName)
	}
}

// TagRulesParse парсинг правил.
func TagRulesParse(tagRule string) [][]string {
	tempSl := strings.Split(tagRule, ",")
	slOfTags := make([][]string, len(tempSl))
	for index, v := range tempSl {
		slOfTags[index] = append(slOfTags[index], strings.Split(v, "=")...)
	}

	return slOfTags
}

// RequireRuleCheck проверка правила "required"
func RequireRuleCheck(fieldValue reflect.Value, tagRule string, fieldName string) error {
	switch fieldValue.Kind() {
	case reflect.String:
		if tagRule == "required" && fieldValue.Len() == 0 {
			return fmt.Errorf("field '%s' is required", fieldName)
		}
	case reflect.Int:
		if tagRule == "required" && fieldValue.Int() == 0 {
			return fmt.Errorf("field '%s' is required", fieldName)
		}
	case reflect.Float32:
		if tagRule == "required" && fieldValue.Float() == 0 {
			return fmt.Errorf("field '%s' is required", fieldName)
		}
	case reflect.Float64:
		if tagRule == "required" && fieldValue.Float() == 0 {
			return fmt.Errorf("field '%s' is required", fieldName)
		}
	}

	return nil
}

// MinValueRuleCheck проверка правила "minValue"
func MinValueRuleCheck(fieldValue reflect.Value, tagRule []string, fieldName string) error {
	switch fieldValue.Kind() {
	case reflect.Int:
		value, err := ToIntConvert(tagRule[1])
		if err != nil {
			return err
		}
		if tagRule[0] == "minValue" && int(fieldValue.Int()) < value {
			return fmt.Errorf("field '%s' less than '%d'", fieldName, value)
		}
	case reflect.Float32:
		value, err := ToFloat32Convert(tagRule[1])
		if err != nil {
			return err
		}
		if tagRule[0] == "minValue" && float32(fieldValue.Float()) < value {
			return fmt.Errorf("field '%s' less than '%s'", fieldName, fmt.Sprintf("%.2f", value))
		}
	case reflect.Float64:
		value, err := ToFloat64Convert(tagRule[1])
		if err != nil {
			return err
		}
		if tagRule[0] == "minValue" && fieldValue.Float() < value {
			return fmt.Errorf("field '%s' less than '%s'", fieldName, fmt.Sprintf("%.2f", value))
		}
	}

	return nil
}

// MaxValueRuleCheck проверка правила "maxValue"
func MaxValueRuleCheck(fieldValue reflect.Value, tagRule []string, fieldName string) error {
	switch fieldValue.Kind() {
	case reflect.Int:
		value, err := ToIntConvert(tagRule[1])
		if err != nil {
			return err
		}
		if tagRule[0] == "maxValue" && int(fieldValue.Int()) > value {
			return fmt.Errorf("field '%s' greater than '%d'", fieldName, value)
		}
	case reflect.Float32:
		value, err := ToFloat32Convert(tagRule[1])
		if err != nil {
			return err
		}
		if tagRule[0] == "maxValue" && float32(fieldValue.Float()) > value {
			return fmt.Errorf("field '%s' greater than '%s'", fieldName, fmt.Sprintf("%.2f", value))
		}
	case reflect.Float64:
		value, err := ToFloat64Convert(tagRule[1])
		if err != nil {
			return err
		}
		if tagRule[0] == "maxValue" && fieldValue.Float() > value {
			return fmt.Errorf("field '%s' greater than '%s'", fieldName, fmt.Sprintf("%.2f", value))
		}
	}

	return nil
}

// MinLenRuleCheck проверка правила "minLen"
func MinLenRuleCheck(fieldValue reflect.Value, tagRule []string, fieldName string) error {
	if fieldValue.Kind() == reflect.String {
		value, err := ToIntConvert(tagRule[1])
		if err != nil {
			return err
		}
		if tagRule[0] == "minLen" && utf8.RuneCountInString(fieldValue.String()) < value {
			return fmt.Errorf("len of field '%s' less than '%d'", fieldName, value)
		}
	}
	return nil
}

// MaxLenRuleCheck проверка правила "maxLen"
func MaxLenRuleCheck(fieldValue reflect.Value, tagRule []string, fieldName string) error {
	if fieldValue.Kind() == reflect.String {
		value, err := ToIntConvert(tagRule[1])
		if err != nil {
			return err
		}
		if tagRule[0] == "maxLen" && utf8.RuneCountInString(fieldValue.String()) > value {
			return fmt.Errorf("len of field '%s' greater than '%d'", fieldName, value)
		}
	}

	return nil
}

func ToIntConvert(str string) (int, error) {
	res, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}

	return res, nil
}

func ToFloat32Convert(str string) (float32, error) {
	res, err := strconv.ParseFloat(str, 32)
	if err != nil {
		return 0, err
	}

	return float32(res), nil
}

func ToFloat64Convert(str string) (float64, error) {
	res, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, err
	}

	return res, nil
}
