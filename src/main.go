package main

import "fmt"

func main() {
	// Parse flags
	templateFile, outputFile := parseFlags()

	// Load default env
	loadDefaultEnv()

	// Get template 
	templateContent := getTemplate(templateFile)

	// Write to file
	renderConfig(outputFile, templateContent)

	fmt.Println("Config file generated successfully!")
}
