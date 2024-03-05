package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
)

// main is the entry point of the program.
// It prompts the user to choose an operation type (1 for adding a new project, 2 for adding a microservice to an existing project),
// reads user input, and performs the corresponding operation.
func main() {
	myFigure := figure.NewFigure("MK Gen", "doom", true)
	// myFigure.Blink(3000, 500, -1)
	myFigure.Print()
	color.Cyan("MK's Mircoservices Project Generator\n")
	reader := bufio.NewReader(os.Stdin)

	// Ask for the operation type
	fmt.Print("Enter 1 to add a new project, or 2 to add a microservice to an existing project: ")
	var operationType int
	_, err := fmt.Scan(&operationType)
	if err != nil {
		color.Red("Error reading operation type: %v", err)
	}

	switch operationType {
	case 1:
		// Existing behavior for creating a new project
		fmt.Print("Enter the name of the project: ")
		projectName, _ := reader.ReadString('\n')
		projectName = strings.TrimSpace(projectName)

		fmt.Print("How many services does your project have? ")
		var serviceCount int
		_, err := fmt.Scan(&serviceCount)
		if err != nil {
			color.Red("Error reading service count: %v", err)
		}

		serviceNames := make([]string, serviceCount)
		for i := 0; i < serviceCount; i++ {
			fmt.Printf("Enter the name of service #%d: ", i+1)
			serviceName, _ := reader.ReadString('\n')
			serviceNames[i] = strings.TrimSpace(serviceName)
		}

		generateProjectStructure(projectName, serviceNames)
	case 2:
		// New behavior for adding a microservice to an existing project
		fmt.Print("Enter the name of the microservice: ")
		serviceName, _ := reader.ReadString('\n')
		serviceName = strings.TrimSpace(serviceName)

		fmt.Print("Enter 1 to create the microservice in the current directory, or 2 to specify a different directory: ")
		var directoryType int
		_, err := fmt.Scan(&directoryType)
		if err != nil {
			color.Red("Error reading directory type: %v", err)
		}

		var basePath string
		switch directoryType {
		case 1:
			basePath = "."
		case 2:
			fmt.Print("Enter the path to the directory where the microservice should be created: ")
			basePath, _ = reader.ReadString('\n')
			basePath = strings.TrimSpace(basePath)
		default:
			color.Red("Invalid directory type: %d", directoryType)
		}

		generateProjectStructure(basePath, []string{serviceName})
	default:
		color.Red("Invalid operation type: %d", operationType)
	}
}

func generateProjectStructure(projectName string, serviceNames []string) error {
	basePath := filepath.Join(".", projectName)

	// Common directories and files
	commonStructure := []string{
		"pkg/auth/jwt.go",
		"pkg/auth/middleware.go",
		"pkg/database/mongodb.go",
		"pkg/database/s3.go",
		"pkg/grpc/setup.go",
		"pkg/model/",
		"pkg/repository/",
		"pkg/service/",
		"pkg/handler/",
		"deployments/kubernetes/",
		"protobuf/",
		".env.example",
		"README.md",
	}

	// Create common structure
	for _, path := range commonStructure {
		if err := createPath(basePath, path); err != nil {
			return err
		}
	}

	// Create service-specific structure
	for _, serviceName := range serviceNames {
		serviceStructure := []string{
			fmt.Sprintf("cmd/%s/main.go", "service-"+serviceName),
			fmt.Sprintf("cmd/%s/Dockerfile", "service-"+serviceName),
			fmt.Sprintf("deployments/docker/%s/Dockerfile", "service-"+serviceName),
			fmt.Sprintf("protobuf/%s.proto", "service-"+serviceName),
		}

		for _, path := range serviceStructure {
			if err := createPath(basePath, path); err != nil {
				return err
			}
		}
	}

	return nil
}

func createPath(basePath, path string) error {
	fullPath := filepath.Join(basePath, path)
	if filepath.Ext(path) == "" { // If no file extension, assume directory
		if err := os.MkdirAll(fullPath, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %v", fullPath, err)
		}
	} else { // Otherwise, create file
		if err := os.MkdirAll(filepath.Dir(fullPath), 0755); err != nil {
			return fmt.Errorf("failed to create directory for file %s: %v", fullPath, err)
		}
		file, err := os.Create(fullPath)
		if err != nil {
			return fmt.Errorf("failed to create file %s: %v", fullPath, err)
		}
		defer file.Close()
	}
	return nil
}

