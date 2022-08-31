package monster

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

// 綁定一個 *monster
func (m *Monster) Store() bool {
	// 先序列化
	data, err := json.Marshal(m)
	if err != nil {
		fmt.Println("marshal err = ", err)
		return false
	}
	// 保存文件

	filePath := "./monster.ser"
	err = ioutil.WriteFile(filePath, data, 0600)
	if err != nil {
		fmt.Println("writer file err = ", err)
		return false
	}

	return true
}

func (m *Monster) ReStore() bool {
	filePath := "monster.ser"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("reader file err")
		return false
	}
	err = json.Unmarshal(data, m)
	if err != nil {
		fmt.Println("Unmarshal false")
		return false
	}
	return true
}
