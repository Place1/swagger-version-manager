package commands

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/user"
	"strings"

	"gopkg.in/cheggaaa/pb.v1"
)

func Use(version string) error {
	output := outputPath(version)
	url := downloadPath(version)
	err := downloadFile(url, output)
	if err != nil {
		return err
	}
	setSwaggerBinFile(version)
	fmt.Printf("using swagger-codegen %s\n", version)
	return nil
}

func downloadPath(version string) string {
	return fmt.Sprintf("http://search.maven.org/remotecontent?filepath=io/swagger/swagger-codegen-cli/%s/swagger-codegen-cli-%s.jar", version, version)
}

func outputPath(version string) string {
	currentUser, _ := user.Current()
	homeDir := currentUser.HomeDir
	os.MkdirAll(fmt.Sprintf("%v/.swagger-version-manager", homeDir), os.ModePerm)
	return fmt.Sprintf("%s/.swagger-version-manager/swagger-codegen.%v.jar", homeDir, version)
}

func trimLines(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	output := make([]string, len(lines))
	for index, line := range lines {
		output[index] = strings.TrimSpace(line)
	}
	return strings.Join(output, "\n")
}

func setSwaggerBinFile(version string) error {
	filePath := "/usr/local/bin/swagger-codegen"
	scriptContent := fmt.Sprintf(
		trimLines(`
			#!/bin/bash
			set -e
			java -jar "%s" $@;
		`),
		outputPath(version),
	)
	fmt.Printf("updated %s\n", filePath)
	err := ioutil.WriteFile(filePath, []byte(scriptContent), os.ModePerm)
	return err
}

func downloadFile(url string, output string) error {
	_, err := os.Stat(output)
	if err == nil {
		fmt.Printf("using cached swagger-codegen-cli.jar at: %s\n", output)
		return nil
	}

	out, err := os.Create(output)
	if err != nil {
		return errors.New("failed to create file on local filesystem")
	}
	defer out.Close()

	fmt.Printf("downloading %s\n", url)
	resp, err := http.Get(url)
	if err != nil {
		return errors.New("failed to request swagger-codegen-cli.jar from maven")
	}
	defer resp.Body.Close()

	bar := pb.New(int(resp.ContentLength)).SetUnits(pb.U_BYTES)
	bar.Format("[=> ]")
	bar.Start()
	_, err = io.Copy(out, bar.NewProxyReader(resp.Body))
	if err != nil {
		return errors.New("failed to download swagger-codegen-cli.jar from maven")
	}
	bar.Finish()
	return nil
}
