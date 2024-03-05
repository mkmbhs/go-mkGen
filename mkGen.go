package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Ask for the project name
	fmt.Print("Enter the name of the project: ")
	projectName, _ := reader.ReadString('\n')
	projectName = strings.TrimSpace(projectName)

	// Ask for the number of services
	fmt.Print("How many services does your project have? ")
	var serviceCount int
	_, err := fmt.Scan(&serviceCount)
	if err != nil {
		log.Fatalf("Error reading service count: %v", err)
	}

	// Ask for the names of the services
	serviceNames := make([]string, serviceCount)
	for i := 0; i < serviceCount; i++ {
		fmt.Printf("Enter the name of service #%d: ", i+1)
		serviceName, _ := reader.ReadString('\n')
		serviceNames[i] = strings.TrimSpace(serviceName)
	}

	// Generate the project structure
	generateProjectStructure(projectName, serviceNames)
}

func generateProjectStructure(projectName string, serviceNames []string) {
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
		createPath(basePath, path)
	}

	// Create service-specific structure
	for _, serviceName := range serviceNames {
		serviceStructure := []string{
			fmt.Sprintf("cmd/%s/main.go", serviceName),
			fmt.Sprintf("cmd/%s/Dockerfile", serviceName),
			fmt.Sprintf("deployments/docker/%s/Dockerfile", serviceName),
			fmt.Sprintf("protobuf/%s.proto", serviceName),
		}

		for _, path := range serviceStructure {
			createPath(basePath, path)
		}

		// Initialize Go module for each service
		initGoModule(filepath.Join(basePath, "cmd", serviceName), serviceName)
	}
}

func createPath(basePath, path string) {
	fullPath := filepath.Join(basePath, path)
	if filepath.Ext(path) == "" { // If no file extension, assume directory
		log.Printf("Creating directory: %s\n", fullPath)
		if err := os.MkdirAll(fullPath, os.ModePerm); err != nil {
			log.Fatalf("Failed to create directory %s: %v", fullPath, err)
		}
	} else { // Otherwise, create file
		log.Printf("Creating file: %s\n", fullPath)
		if err := os.MkdirAll(filepath.Dir(fullPath), os.ModePerm); err != nil {
			log.Fatalf("Failed to create directory for file %s: %v", fullPath, err)
		}
		file, err := os.Create(fullPath)
		if err != nil {
			log.Fatalf("Failed to create file %s: %v", fullPath, err)
		}
		file.Close()
	}
}

func initGoModule(servicePath, serviceName string) {
	log.Printf("Initializing Go module for service %s\n", serviceName)

	// Change to the service directory
	os.Chdir(servicePath)

	// Run `go mod init`
	cmd := exec.Command("go", "mod", "init", serviceName)
	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to initialize Go module for service %s: %v", serviceName, err)
	}

	// Run `go mod tidy`
	cmd = exec.Command("go", "mod", "tidy")
	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to tidy Go module for service %s: %v", serviceName, err)
	}

	// Change back to the original directory
	os.Chdir("../../..")
}
