package core_a

import "fmt"

func Version() string {
	return fmt.Sprintf("Core 1 version: %d", CORE_A_VERSION)
}
