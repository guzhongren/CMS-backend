pipeline {
    agent any
    // agent {
    //     docker { 
    //         image 'golang'
    //         args '-u 0:0'
    //     }
    // }
    tools {go "go1.12"}
    stages {
        stage('获取SCM') {
            steps{
                checkout scm
            }
        }
        stage('Test') {
            steps {
                sh 'export CGO_ENABLED=0'
                sh 'export GOPROXY=https://goproxy.io'
                sh 'go clean -cache'
                sh 'go test ./... -v -short'
            }
        }
        stage('构建并推送镜像') {
            steps{
                echo '开始构建镜像。。。'
                withCredentials([usernamePassword(credentialsId: 'docker-register', passwordVariable: 'dockerPassword', usernameVariable: 'dockerUser')]) {
                    sh "./build_script/build_image.sh ${dockerUser} ${dockerPassword} backend"
                }
            }
        }
        // stage('保留本地最新的三个镜像') {
        //     steps{
        //         echo '删除最新三个以外的镜像...'
        //         // sh './build_script/build_image.sh cms-backend'
        //     }
        // }
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