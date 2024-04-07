package main

import (
	"fmt"
	"os"
)

func main() {
	// Parse flags
	templateFile, outputFile, envFile := parseFlags()

	// Load default env
	loadDefaultEnv(envFile)

	// Get template 
	templateContent := getTemplate(templateFile)

	// Write to file
	renderConfig(outputFile, templateContent)

	if (os.Getenv("MCONFIG_DEBUG") == "true") {
		fmt.Printf("Rendered %s", templateFile)
	}
}
