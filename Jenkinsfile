pipeline {
    agent any
    stages {
        stage('Pull code') {
            steps {
                    checkout([$class: 'GitSCM', branches: [[name: 'dev']], extensions: [], userRemoteConfigs: [[credentialsId: '698862fd-b5af-4c2b-920c-42ed9ab6ceef', url: 'https://github.com/SilverSky9/final-SDTE.git']]])
                }
        }
        stage('Docker') {
            steps {
                sh "docker -v"
            }
        }
        stage('Docker compose') {
            steps {
                sh "docker-compose version"
            }
        }
    }
}