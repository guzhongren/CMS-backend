pipeline {
    // agent any
    agent {
        docker { image 'golang' }
    }
    // tools {go "go1.12"}
    stages {
        stage('获取SCM') {
            steps{
                checkout scm
            }
        }
        // stage('Build') {                
        //     steps {      
        //         // Create our project directory.
        //         sh 'cd ${GOPATH}/src'
        //         sh 'mkdir -p ${GOPATH}/src/CMS-backend'

        //         // Copy all files in our Jenkins workspace to our project directory.                
        //         sh 'cp -r ${WORKSPACE}/* ${GOPATH}/src/CMS-backend'

        //         // Copy all files in our "vendor" folder to our "src" folder.
        //         // sh 'cp -r ${WORKSPACE}/vendor/* ${GOPATH}/src'

        //         // Build the app.
        //         sh 'go build'
        //     }            
        // }
        stage('Test') {
            steps {
                sh 'whoami'
                sh 'export CGO_ENABLED=0'
                sh 'go clean -cache'
                sh 'go test ./... -v -short'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying....'
            }
        }
    }
    post{
        always {
            echo '执行完毕'
        }
        failure{
            echo '执行失败'
        }
        success{
            echo 'pipeline 执行成功'
        }
        unstable{
            echo '测试失败，需要检查测试或者查看编码规范！'
        }
        aborted{
            echo 'Pipeline 被终止，如需继续，请重新运行 pipeline'
        }
    }
}