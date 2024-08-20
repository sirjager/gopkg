package utils

import (
	"encoding/json"
	"fmt"
)

func PrettyPrint[T any](input T) error {
	prettyJSON, err := json.MarshalIndent(input, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(prettyJSON))
	return nil
}
