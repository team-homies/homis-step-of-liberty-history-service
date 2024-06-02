package util

import "encoding/json"

// Struct를 맵으로 변환하는 함수
func StructToMap(data any) (result map[string]any, err error) {
	result = make(map[string]any)

	jsonData, err := json.Marshal(data)
	if err != nil {
		return
	}

	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		return
	}

	return
}
