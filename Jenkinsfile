pipeline {
    agent any
    environment { 
        HUB_DOMAIN = 'hub.k8s.com', // 'docker 私有仓库域'
        PROJECT_NAME = 'cms', // 项目名称
        NAMESPACE_NAME = 'cms', // namespace名称
        DEPLOYMENT_NAME = 'backend', // deployment 名称
        CONTAINER_NAME = 'backend', // 容器名称
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
                echo '开始构建镜像...'
                sh "go build -o cms"
                withCredentials([usernamePassword(credentialsId: 'docker-register', passwordVariable: 'dockerPassword', usernameVariable: 'dockerUser')]) {
                    sh "./build_script/build_image.sh ${HUB_DOMAIN} ${dockerUser} ${dockerPassword} ${PROJECT_NAME} ${CONTAINER_NAME}"
                }
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying....'
                sh "./build_script/deploy_image.sh ${NAMESPACE_NAME} ${DEPLOYMENT_NAME} ${CONTAINER_NAME}"
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