package main

import (
	"fmt"
	"os"
)


func main() {
	// Parse flags
	templateDir, outputDir := parseFlags()

	if templateDir == "" || outputDir == "" {
		fmt.Println("Usage: mconfig -templateDir <template> -outputDir <output>")
		os.Exit(1)
	}
  
	// Get templates paths
	paths := getTemplatesPaths(templateDir)

	for _, path := range paths {
		if os.Getenv("MCONFIG_DEBUG") == "true" {
			fmt.Printf("Redering %s \n", path)
		}

		// Get template
		templateContent := getTemplate(path)

		// Cut template dir
		outputPath := cutTemplateDirFromOutputPath(path, templateDir)

		// Create output subfolder if not exits
		createSubDirsInOutput(outputPath, outputDir)
		
		// Write to file
		renderConfig(fmt.Sprintf("%s/%s", outputDir, outputPath), templateContent)

		if os.Getenv("MCONFIG_DEBUG") == "true" {
			fmt.Printf("Rendered %s \n", path)
		}
	}
}
