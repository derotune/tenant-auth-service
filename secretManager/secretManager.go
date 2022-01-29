package secretManager

import (
	"fmt"
	"strings"

	"github.com/derotune/tenant-auth-service/secretManager/manager/gsm"
)

func GetSecret(key string) string {
	splittedKey := strings.Split(key, ":")

	if len(splittedKey) > 1 {
		switch splittedKey[0] {
		case "gsm":
			key = gsm.Get(splittedKey[1])
		default:
			fmt.Println("No secret manager found. Just use the key")
		}
	}

	return key
}
