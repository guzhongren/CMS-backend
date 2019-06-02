pipeline {
    agent any

    stages {
        stage('Test') {
            steps {
                sh 'export CGO_ENABLED=0'
                sh 'go clean -cache'
                sh 'go test'
            }
        }
        stage('Build') {
            steps {
                echo 'Building..'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying....'
            }
        }
    }
}