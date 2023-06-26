package Utils

import "fmt"

func ErrorMessage(message string) string {
	return fmt.Sprintf(`{"message": "%s"}`, message)
}