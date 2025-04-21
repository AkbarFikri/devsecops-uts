# DevSecOps UTS Project

## Overview

This is a simple authentication service application developed as part of a university exam (UTS) for the DevSecOps course. The project demonstrates implementing a complete CI/CD pipeline with security testing using GitHub and Jenkins.

## Features

- Basic authentication service with login functionality
- Unit testing with mocks
- CI/CD pipeline implementation
- Security scanning integration
- Infrastructure as Code (IaC) approach

## Technology Stack

- **Backend**: Go (Golang)
- **Testing**: Go testing package, Testify
- **CI/CD**: GitHub Actions, Jenkins
- **Security Scanning**: SonarQube, OWASP ZAP
- **Containerization**: Docker
- **Orchestration**: Kubernetes (optional)

## Project Structure

```
devsecops-uts/
├── src/                  # Source code
│   ├── handler/          # HTTP Handler Layer
│   ├── repository/       # Data access layer
│   ├── service/          # Business logic
├── test/                 # Test files
│   ├── service/          # Service tests
│   │   └── mocks/        # Mock implementations
├── Jenkinsfile           # Jenkins pipeline definition
├── .github/workflows/    # GitHub Actions workflows
├── Dockerfile            # Container configuration
├── Makefile              # Build automation
└── README.md             # Project documentation
```

## CI/CD Pipeline

The pipeline implements the following stages:

1. **Code Commit** (GitHub)
2. **Static Code Analysis** (SonarQube)
3. **Unit Testing** (Go test)
4. **Build Artifact** (Docker image)
5. **Deployment** (Kubernetes)
6. **Monitoring** (Prometheus/Grafana)

## Security Measures

- SAST (Static Application Security Testing) with SonarQube
- Dependency scanning with OWASP Dependency-Check
- Secret scanning with GitGuardian or similar
- Infrastructure security scanning

## Setup Instructions

### Prerequisites

- Go 1.20+
- Docker
- Jenkins
- Kubernetes cluster (optional)

### Local Development

1. Clone the repository:
   ```bash
   git clone https://github.com/AkbarFikri/devsecops-uts.git
   cd devsecops-uts
   ```

2. Build the application:
   ```bash
   make build
   ```

3. Run tests:
   ```bash
   make test
   ```

4. Run the application:
   ```bash
   make run
   ```

### Jenkins Setup

1. Create a new pipeline job in Jenkins
2. Set the pipeline definition to point to the `Jenkinsfile` in the repository
3. Configure necessary credentials for Docker and Kubernetes
4. Set up webhook in GitHub for automatic trigger on push

### GitHub Actions

The workflow files in `.github/workflows/` will automatically run on push to:

- Main branch: full CI/CD pipeline
- Feature branches: run tests and security scans

## Usage Examples

### API Endpoints

```
POST /login
Parameters:
- username: string
- password: string

Response:
- Success: 200 OK with username
- Failure: 401 Unauthorized
```

### Curl Example

```bash
curl -X POST http://localhost:3000/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"secret"}'
```

## Monitoring

The deployed application includes:

- Health checks at `/`

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- University curriculum for DevSecOps course
- Jenkins and GitHub documentation
- OWASP security tool