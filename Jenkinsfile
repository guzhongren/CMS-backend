pipeline {
    agent any
    // agent {
    //     docker { 
    //         image 'golang'
    //         args '-u 0:0'
    //     }
    // }
    parameters {
        string(name: 'hub_domain', defaultValue: 'hub.k8s.com', description: 'docker 私有仓库域')
        string(name: 'project_name', defaultValue: 'cms', description: '项目名称')
        string(name: 'app_name', defaultValue: 'backend', description: '容器名称')
    }
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
                sh 'export GO111MODULE=on'
                sh 'export GOPROXY=https://goproxy.cn'
                sh 'go clean -cache'
                sh 'go test ./... -v -short'
            }
        }
        stage('构建并推送镜像') {
            steps{
                echo '开始构建镜像。。。'
                withCredentials([usernamePassword(credentialsId: 'docker-register', passwordVariable: 'dockerPassword', usernameVariable: 'dockerUser')]) {
                    sh "./build_script/build_image.sh ${params.hub_domain} ${dockerUser} ${dockerPassword} ${params.project_name} ${params.app_name}"
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
                sh "./build_script/deploy_image.sh ${params.project_name} ${params.app_name}"
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