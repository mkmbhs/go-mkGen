# Go Microservice Generator

## Table of Contents
- [Overview](#overview)
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [Open Source](#open-source)
- [Contributing](#contributing)
- [Contributors](#contributors)
- [License](#license)


## Overview
Go Microservice Generator is a command-line tool that simplifies the creation of microservices project structures in Go. It prompts users for project-specific details and automatically generates a comprehensive directory and file structure, including Docker and Kubernetes configurations, adhering to best practices in microservice architecture.

## Features
- Interactive CLI for custom microservices project generation.
- Generates Go module initialization and tidy up for each service.
- Supports generating multiple services with predefined structures.
- Creates Dockerfiles and Kubernetes deployment files.
- Open-source and easily extendable for additional features.

## Prerequisites
- Go 1.15 or later
- Git
- Access to the command line/terminal

## Installation

To install Go Microservice Generator, run the following command in your terminal:

```bash
curl -s https://github.com/mkmbhs/go-mkGen/edit/master/install-mkGen.sh | bash
```

This script will:
- Check if Go is installed on your machine.
- Download the latest version of the Go Microservice Generator script.
- Compile and move the binary to a location in your system's PATH for global access.

**Note:** We highly recommend inspecting the script before executing it for security purposes.

## Usage

After installation, you can start generating your microservices project by running:

```
mkGen
```

When you run the command, the generator will prompt you to enter the following details:
- If you want to generate a new project or add a new service to an existing project.
- The name of your project.
- The number of services you want to generate.
- The names of the services you want to generate.
- And in case you want to add a new service to an existing project, you will be prompted to route to the project directory, or run in the current directory.

Follow the interactive prompts to specify your project name and the number of services you wish to generate, including their names. The generator will create a new directory with your project name and set up the service structures within it.

## Open Source

This project is open source and welcomes contributions from the community. Whether it's adding new features, fixing bugs, or improving documentation, your help is appreciated.

## Contributing

We welcome contributions from the community, including feature requests, bug reports, and pull requests. Please visit our [GitHub Issues](https://github.com/mkmbhs/mkgen/issues) page to report issues or suggest enhancements.

Before submitting a pull request, please:
- Fork the repository and create your branch from `master`.
- Ensure the test suite passes (if applicable).
- Update the README.md with details of changes, including new environment variables, exposed ports, useful file locations, and container parameters.

## Contributors

We thank all the contributors who have helped make this project better. Your contributions are greatly appreciated!

## Contributors

Thanks goes to these wonderful people:

<table>
    <tr>
        <td align="center" style="padding:10px; border:1px solid #ccc;">
            <a href="https://github.com/mkmbhs">
                <img src="https://github.com/mkmbhs.png?size=70" width="70" height="70" style="border-radius:50%;">
                <br>
                <b>Mohamed Kamal</b>
            </a>
            <br>
            <sub><b>Initial work</b></sub>
        </td>
    </tr>
</table>

> **Note:** You can update this list as your project grows.

## License

This project is open-sourced under the MIT License. See the [LICENSE](LICENSE) file for details.

---

For more information and updates, please visit our [GitHub repository](https://github.com/mkmbhs/go-mkgen).
