package jsonHelper

import (
	"encoding/json"
	"errors"
	"strings"
)

type JsonObject struct {
	data map[string]interface{}
}

var (
	errNil    = errors.New("error key not exist") // key所对应的值不存在
	errFormat = errors.New("error formart")       // 类型转换错误
)

// NewJsonObject 创建一个json对象
func NewJsonObject(str string) (*JsonObject, error) {
	jsonObject := new(JsonObject)
	err := json.Unmarshal([]byte(str), &jsonObject.data)
	if err != nil {
		return nil, err
	}
	return jsonObject, nil
}

func (o *JsonObject) getValue(key string) (interface{}, error) {
	m := o.data
	keys := strings.Split(key, ".")

	var (
		value interface{}
		ok    = false
	)

	for i := 0; i < len(keys)-1; i++ {
		value, ok = m[keys[i]]
		if !ok {
			return nil, errNil
		}
		m, ok = value.(map[string]interface{})
		if !ok {
			return nil, errFormat
		}
	}

	value, ok = m[keys[len(keys)-1]]
	if !ok {
		return nil, errNil
	}
	return value, nil
}

// GetBool 获取Bool值
func (o *JsonObject) GetBool(key string) (bool, error) {
	value, err := o.getValue(key)
	if err != nil {
		return false, err
	}

	result, ok := value.(bool)
	if !ok {
		return false, errFormat
	}
	return result, nil
}

// GetFloat64 获取float64值
func (o *JsonObject) GetFloat64(key string) (float64, error) {
	value, err := o.getValue(key)
	if err != nil {
		return 0, err
	}

	result, ok := value.(float64)
	if !ok {
		return 0, errFormat
	}
	return result, nil
}

// GetString 获取string值
func (o *JsonObject) GetString(key string) (string, error) {
	value, err := o.getValue(key)
	if err != nil {
		return "", err
	}

	result, ok := value.(string)
	if !ok {
		return "", errFormat
	}
	return result, nil
}

// GetBool 获取Bool值
func GetBool(str string, key string) (bool, error) {
	object, err := NewJsonObject(str)
	if err != nil {
		return false, err
	}
	return object.GetBool(key)
}

// GetFloat64 获取float64值
func GetFloat64(str string, key string) (float64, error) {
	object, err := NewJsonObject(str)
	if err != nil {
		return 0, err
	}
	return object.GetFloat64(key)
}

// GetString 获取string值
func GetString(str string, key string) (string, error) {
	object, err := NewJsonObject(str)
	if err != nil {
		return "", err
	}
	return object.GetString(key)
}
