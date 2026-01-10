pipeline {
    agent {
      
        docker { image 'golang:1.25.5' }
    }
    stages {
        stage('Limpieza y Checkout') {
            steps {
                deleteDir() 
                checkout scm
            }
        }
        stage('Build') {
            steps {
                sh 'go build -o mi-app'
            }
        }
        stage('Test') {
            steps {
                sh 'go test -v ./...'
            }
        }
    }
}