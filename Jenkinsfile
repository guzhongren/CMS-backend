pipeline {
    agent any
    ws("${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}/") {
        withEnv(["GOPATH=${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"]) {
            stages {
                stage('检出...') {
                    steps{
                        checkout scm
                    }
                }
                stage('Test') {
                    steps {
                        sh 'export CGO_ENABLED=0'
                        sh 'env'
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
    }
}