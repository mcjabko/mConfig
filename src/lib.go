package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/joho/godotenv"
)

func parseFlags() (string, string, string) {
	templateFile := flag.String("templateDir", "", "<template>")
	outputFile := flag.String("outputDir", "", "<output>")
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

func getTemplatesPaths(templateDir string) []string {
	paths := []string{}

	err := filepath.Walk(templateDir, func(path string, info fs.FileInfo, err error) error {
		matched, _ := regexp.Match(`\.yml$|\.yaml$|\.properties|\.json|\.txt`, []byte(path))

		if !info.IsDir() && matched {
			paths = append(paths, path)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return paths
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

func cutTemplateDirFromOutputPath(path string, templateDir string) string {
	regex := regexp.MustCompile(fmt.Sprintf("%s/", templateDir))
	return regex.ReplaceAllString(path, "")
}

func createSubDirsInOutput(path string, outputDir string) {
	_, err := os.Stat(fmt.Sprintf("%s/%s", outputDir, path))
	if err != nil {
		splitedPath := strings.Split(path, "/")
		err := os.MkdirAll(fmt.Sprintf("%s/%s", outputDir, strings.Join(splitedPath[:len(splitedPath)-1], "")), os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func renderConfig(outputFile string, templateContent []byte) {
	err := os.WriteFile(outputFile, parseEnvToTemplate(templateContent), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
