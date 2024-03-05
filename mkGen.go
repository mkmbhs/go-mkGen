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

func main() {
	myFigure := figure.NewFigure("MK Gen", "doom", true)
	// myFigure.Blink(3000, 500, -1)
	myFigure.Print()
	color.Cyan("MK's Mircoservices Project Generator\n")

	reader := bufio.NewReader(os.Stdin)

	// Ask for the project name
	fmt.Print("Enter the name of the project: ")
	projectName, err := reader.ReadString('\n')
	if err != nil {
		color.Red("Error reading project name: %v", err)
		return
	}
	projectName = strings.TrimSpace(projectName)

	// Ask for the number of services
	fmt.Print("How many services does your project have? ")
	var serviceCount int
	_, err = fmt.Scan(&serviceCount)
	if err != nil {
		color.Red("Error reading service count: %v", err)
		return
	}

	// Ask for the names of the services
	serviceNames := make([]string, serviceCount)
	for i := 0; i < serviceCount; i++ {
		fmt.Printf("Enter the name of service #%d: ", i+1)
		serviceName, err := reader.ReadString('\n')
		if err != nil {
			color.Red("Error reading service name: %v", err)
			return
		}
		serviceNames[i] = strings.TrimSpace(serviceName)
	}

	// Generate the project structure
	if err := generateProjectStructure(projectName, serviceNames); err != nil {
		color.Red("Error generating project structure: %v", err)
		return
	}

	color.Green("Project %s created successfully.", projectName)
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
