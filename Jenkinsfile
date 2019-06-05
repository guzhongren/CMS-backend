pipeline {
    agent any

    stages {
        stage('获取SCM') {
            steps{
                checkout scm
            }
        }
        stage('Test') {
            steps {
                sh 'export CGO_ENABLED=0'
                // sh 'go clean -cache'
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