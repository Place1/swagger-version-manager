// +build linux darwin netbsd openbsd freebsd

package commands

import (
	"fmt"
	"strings"
)

func getSwaggerExecutablePath() string {
	return "/usr/local/bin/swagger-codegen"
}

func getSwaggerExecutableContent(swaggerJar string) string {
	return fmt.Sprintf(
		trimLines(`
			#!/bin/sh
			set -e
			java -jar "%s" $@;
		`),
		swaggerJar,
	)
}

func trimLines(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	output := make([]string, len(lines))
	for index, line := range lines {
		output[index] = strings.TrimSpace(line)
	}
	return strings.Join(output, "\n")
}
