pipeline {
    agent any

    tools {
        go '1.24.1'
    }

    environment {
        DOCKER_CREDENTIALS = credentials('dockerhub')
        GIT_COMMIT = sh(script: 'git rev-parse HEAD', returnStdout: true).trim()
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Unit Tests') {
            steps {
                sh 'go test -v ./test/service/...'
            }
        }
    }
}