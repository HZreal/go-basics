package main

import (
	"encoding/json"
	"fmt"
)

type UserInfo struct {
	ID     int    `json:"id" bson:"id" yaml:"id" xml:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender int    `json:"gender"`
	Avatar string `json:"avatar"`
}

func JsonMarshall() {
	user := UserInfo{ID: 1, Name: "hh", Age: 22, Gender: 1, Avatar: "/users/1/avatar/uuid.png"}
	byteArr, _ := json.Marshal(user) // 返回 []byte
	fmt.Println("byteArr---------", byteArr, len(byteArr))
	str := string(byteArr) // []byte 转 string
	fmt.Println("str---------", str)

	_map := make(map[string]interface{})
	_map["name"] = "huang"
	_map["age"] = 22
	_map["dict"] = map[string]string{"key1": "value1", "key2": "value2"}
	_map["info"] = []map[string]string{
		{"key1": "value1", "key2": "value2"},
		{"key1": "value1", "key2": "value1"},
	}
	byteArr2, _ := json.Marshal(_map)
	//fmt.Println("byteArr2---------", byteArr2, len(byteArr2))
	str2 := string(byteArr2) // []byte 转 string
	fmt.Println("str2---------", str2)

}

func JsonUnmarshall() {
	jsonStr := `{"ID": 1, "Name": "hh", "Age": 22, "Gender": 1, "Avatar": "/users/1/avatar/uuid.png"}`
	byteArr := []byte(jsonStr)
	fmt.Println("byteArr----------", byteArr)

	user := UserInfo{}
	_ = json.Unmarshal(byteArr, &user) // JSON 转换为 结构体
	fmt.Println("user ------------", user)

	_map := make(map[string]interface{})
	_ = json.Unmarshal(byteArr, &_map) // JSON 转换为 map
	fmt.Println("_map ------------", _map)

	// 假设不知道jsonStr中键、值的类型，可以将他解析到interface{}
	var p interface{}
	_ = json.Unmarshal(byteArr, &p)
	v, ok := p.(map[string]interface{}) // 断言p为类型map[string]interface{}  返回：v为断言成功的对应类型值，ok表示是否断言成功
	if ok {
		fmt.Println("断言成功的数据为", v) // 断言成功，输出值
	} else {
		fmt.Println("断言失败")
	}
	// 另一种处理
	m := p.(map[string]interface{})
	fmt.Println("断言成功的数据-----", m)
	for _, v := range m {
		switch v.(type) { // 继续断言value的类型
		case string:
			fmt.Println("\ntype is string, value is ", v.(string))
		case int:
			fmt.Println("\ntype is int, value is ", v.(int))
		case int64:
			fmt.Println("\ntype is int64, value is ", v.(int64))
		case float64:
			fmt.Println("\ntype is float64, value is ", v.(float64))
		case bool:
			fmt.Println("\ntype is bool, value is ", v.(bool))
		case []byte:
			fmt.Println("\ntype is bool, value is ", v.([]byte))
			fmt.Println("[]byte to string ---", string(v.([]byte)))
		default:
			fmt.Printf("\ntype is unknow, print it as %v", v)
		}
	}

}

//struct实现Unmarshaler接口, 便可以实现解析JSON的过程
// UnmarshalJSON 自定义解析
func (p *UserInfo) UnmarshalJSON(data []byte) error {
	// 示例代码使用jsonitor代为解析
	p.ID = 2
	p.Age = 24
	p.Name = "my_test_name"
	return nil
}

// MarshalJSON 自定义编码
func (p *UserInfo) MarshalJSON() ([]byte, error) {
	// 自己编码json
	return []byte(`{"test":"name_test"}`), nil
}

func main() {
	//JsonMarshall()
	JsonUnmarshall()
}
