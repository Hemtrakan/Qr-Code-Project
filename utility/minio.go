package utility

import (
"fmt"
"strings"
)

func MinioCreateObjectName(project string, moduleName string, fileName string) (string, string) {
	return fmt.Sprintf("%s/%s/%s", strings.ToLower(project), strings.ToLower(moduleName), fileName), strings.Split(fileName, ".")[len(strings.Split(fileName, "."))-1]
}
