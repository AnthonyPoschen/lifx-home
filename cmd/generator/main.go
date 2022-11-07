package main

import (
	"bytes"
	_ "embed"
	"flag"
	"fmt"
	"go/format"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"text/template"

	yaml "gopkg.in/yaml.v2"
)

var outputDir string
var packageName string
var version string

//go:embed generator.tmpl
var templateData string

func init() {
	flag.StringVar(&outputDir, "output", "../../pkg/protocol", "Output directory to put generated code")
	flag.StringVar(&packageName, "package", "protocol", "package name of the code generated")
	flag.StringVar(&version, "version", "main", "git version of the repository to generate from, tag or branch name")
}

type templateVars struct {
	PackageName string
	Data        map[string]map[string]interface{}
}

func StripBrackets(value string) string {
	res := strings.Join(strings.Split(value, "<"), "")
	return strings.Join(strings.Split(res, ">"), "")
}

func main() {
	resp, err := http.Get(fmt.Sprintf("https://raw.githubusercontent.com/LIFX/public-protocol/%s/protocol.yml", version))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Failed to read payload: %s", err)
	}
	var config map[string]map[string]interface{}
	yaml.Unmarshal(data, &config)
	t, err := template.New("t").Funcs(template.FuncMap{"StripBrackets": StripBrackets}).Parse(templateData)
	if err != nil {
		fmt.Printf("Failed to load template: %s", err)
		os.Exit(1)
	}
	var processed bytes.Buffer
	err = t.Execute(&processed, templateVars{
		PackageName: packageName,
		Data:        config,
	})
	if err != nil {
		fmt.Printf("failed to execute template %s", err)
		os.Exit(1)
	}
	formatted, err := format.Source(processed.Bytes())
	if err != nil {
		fmt.Printf("%v", string(processed.Bytes()))
		fmt.Printf("Failed to format go code: %s", err)
		os.Exit(1)
	}
	fmt.Println(string(formatted))
}
