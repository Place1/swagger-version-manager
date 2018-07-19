package commands

import (
	"path/filepath"
	"fmt"
)

func getSwaggerExecutablePath() string {
	return filepath.Join("C:/", "Windows", "System32", "swagger-codegen.bat")
}

func getSwaggerExecutableContent(swaggerJar string) string {
	return fmt.Sprintf("java -jar \"%s\" %%*", swaggerJar)
}
