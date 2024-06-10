package main

import (
	"golang.org/x/mod/module"
	"golang.org/x/mod/zip"
	"log"
	"os"
)

func main() {
	// Set the module version information
	modVersion := module.Version{
		Path:    "github.com/aarshkshah1992/prebuilt-ffi-darwin-arm64",
		Version: "v0.0.4",
	}

	// Directory containing the Go source code
	dir := "/Users/aarshshah/src/filecoin/prebuilt-ffi-poc/prebuilt-ffi-darwin-arm64"

	// Output file where the zip will be written
	outputFile, err := os.Create("module.zip")
	if err != nil {
		log.Fatalf("Failed to create output file: %v", err)
	}
	defer outputFile.Close()

	// Create the zip file from the directory
	if err := zip.CreateFromDir(outputFile, modVersion, dir); err != nil {
		log.Fatalf("Failed to create zip file: %v", err)
	}

	log.Println("Module zip file created successfully")
}
