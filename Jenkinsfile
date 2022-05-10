pipeline {
    agent any
    stages {
        stage('Pull code') {
            steps {
                    checkout([$class: 'GitSCM', branches: [[name: 'dev']], extensions: [], userRemoteConfigs: [[credentialsId: '698862fd-b5af-4c2b-920c-42ed9ab6ceef', url: 'https://github.com/SilverSky9/final-SDTE.git']]])
                }
        }
        stage('Unit testing') {
            steps {
                sh "echo Unit testing"
            }
        }
        stage('Component testing') {
            steps {
                sh "echo Component testing"
            }
        }
        stage('Packaging') {
            steps {
                sh "echo Packaging"
            }
        }
        stage('Deploy to target server') {
            steps {
                sh "echo Deploy to target server"
            }
        }
        stage('API/UI testing') {
            steps {
                sh "echo API/UI testing"
            }
        }

    }
}