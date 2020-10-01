pipeline {
    agent any
    stages {
        stage('Build') {
            steps {
                sh 'echo Building...'
            }
        }
        stage('Lint go code') {
            steps {
                sh 'golint client/client.go'
                sh 'golint server/server.go'
            }
        }
        stage('Build docker image') {
            steps {
                sh 'docker build -t server ./server'
                sh 'docker build -t client ./client'
            }
        }
        stage('Push docker image') {
            steps {
                withDockerRegistry([url:"", credentialsId: "docker-hub"]) {
                    sh "docker tag client calebikhuohon/cloud-devops-capstone-client"
                    sh "docker tag server calebikhuohon/cloud-devops-capstone-server"
                    sh "docker push calebikhuohon/cloud-devops-capstone-client"
                    sh "docker push calebikhuohon/cloud-devops-capstone-server"
                }
            }
        stage('Deploying') {
            steps {
                echo 'Deploying to AWS..'
                withAWS(credentials: 'aws', region: 'us-west-2') {
                  sh "aws eks --region us-west-2 update-kubeconfig --name grpc-microservice-cluster"
                  sh "kubectl config use-context arn:aws:eks:us-west-2:988212813982:cluster/grpc-microservice-cluster"
                  sh "kubectl set image deployments/client client=calebikhuohon/cloud-devops-capstone-client:latest"
                  sh "kubectl set image deployments/server server=calebikhuohon/cloud-devops-capstone-server:latest"
                  sh "kubectl apply -f kubernetes-manifests/"
                  sh "kubectl get nodes"
                  sh "kubectl get deployment"
                  sh "kubectl get pod -o wide"
                  sh "kubectl get service/client"
                  sh "kubectl get service/server"
                }
            }
        }
    }
}