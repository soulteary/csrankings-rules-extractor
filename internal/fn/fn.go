package fn

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func MakeJSON(data any) string {
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Failed to convert JSON:", err)
		return ""
	}
	return string(jsonData)
}

func ReleaseJSON(filename, data string) error {
	return os.WriteFile(filepath.Join("release", filename), []byte(data), os.ModePerm)
}
