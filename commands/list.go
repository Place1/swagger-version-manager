package commands

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type MavenSearch struct {
	Response struct {
		NumFound int `json:"numFound"`
		Start    int `json:"start"`
		Docs     []struct {
			ID        string   `json:"id"`
			Group         string   `json:"g"`
			Artifact         string   `json:"a"`
			Version         string   `json:"v"`
			P         string   `json:"p"`
			Timestamp int64    `json:"timestamp"`
			Ec        []string `json:"ec"`
			Tags      []string `json:"tags"`
		} `json:"docs"`
	} `json:"response"`
}

func List() (err error) {
	resp, err := http.Get("http://search.maven.org/solrsearch/select?q=g:%22io.swagger%22+AND+a:%22swagger-codegen-cli%22&core=gav&wt=json")
	if err != nil {
		return err
	}
	result := MavenSearch{}
	json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return err
	}
	for _, value := range result.Response.Docs {
		fmt.Printf("%v\n", value.Version)
	}
	return nil
}
