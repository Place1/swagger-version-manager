package commands

import (
	"fmt"
	"os/exec"
)

func Current() error {
	cmd := exec.Command("swagger-codegen", "version")
	output, err := cmd.Output()
	if err != nil {
		// TODO: handle when swagger-codegen doesn't exist on $PATH
		return err
	}
	fmt.Println(string(output))
	return nil
}
