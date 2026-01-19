package output

import (
	"encoding/json"
	"fmt"

	"passguard/internal/analyzer"
)

func PrintJSON(result analyzer.Result) {
	data, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(data))
}
