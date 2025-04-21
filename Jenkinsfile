pipeline {
    agent any

    tools {
        go '1.24'
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
                sh 'go test -v ./...'
            }
        }

        stage('Security Scan') {
            steps {
                sh 'docker run --rm -v $(pwd):/zap/wrk owasp/zap2docker-stable zap-baseline.py -t http://localhost:8080 -r report.html'
                archiveArtifacts 'report.html'
            }
        }

        stage('Build Docker Image') {
            steps {
                sh '''
                    docker build -t akbarfikri/devsecops-uts:${GIT_COMMIT} .
                    docker tag akbarfikri/devsecops-uts:${GIT_COMMIT} akbarfikri/devsecops-uts:latest
                '''
            }
        }

        stage('Push to Docker Hub') {
            steps {
                withCredentials([usernamePassword(
                    credentialsId: 'dockerhub',
                    usernameVariable: 'akbarfikri',
                    passwordVariable: 'dckr_pat_FSbcAl0Tr_EPCh5-nZnYGl5nTmM'
                )]) {
                    sh '''
                        echo $DOCKER_PASS | docker login -u $DOCKER_USER --password-stdin
                        docker push akbarfikri/devsecops-uts:${GIT_COMMIT}
                        docker push akbarfikri/devsecops-uts:latest
                    '''
                }
            }
        }
    }

    post {
        always {
            cleanWs()
        }
        failure {
            emailext body: 'Build ${BUILD_NUMBER} failed!', subject: 'Build Failed', to: 'akbarfikriabdillah@gmail.com'
        }
    }
}