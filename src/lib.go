package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"github.com/joho/godotenv"
)

func parseFlags() (string, string, string) {
	templateFile := flag.String("template", "", "<template.yml>")
	outputFile := flag.String("output", "", "<output.yml>")
	envFile := flag.String("envFile", ".env", "<.env.defaults>")
	flag.Parse()

	return *templateFile, *outputFile, *envFile
}

func loadDefaultEnv(envFile string) {
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatal(err)
	}
}

func getTemplate(templateFile string) []byte {
	if _, err := os.Stat(templateFile); err != nil {
		log.Fatal("Template doesn't exits")
	}

	template, err := os.ReadFile(templateFile)
	if err != nil {
		log.Fatal(err)
	}

	return template
}

func parseEnvToTemplate(templateContent []byte) []byte {
	output := templateContent
	for _, value := range os.Environ() {
		env := strings.SplitN(value, "=", 2)
		
		output = bytes.Replace(output, []byte(fmt.Sprintf("{%s}", env[0])), []byte(env[1]), -1)
	}
	return output
}

func renderConfig(outputFile string, templateContent []byte) {
	err := os.WriteFile(outputFile, parseEnvToTemplate(templateContent), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
