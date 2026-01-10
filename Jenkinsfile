pipeline {
    agent any
    
    stages {
        stage('Limpiar y Clonar') {
            steps {
                // Borra todo el contenido de la carpeta de trabajo actual
                deleteDir()
                
                // Clona manualmente el repositorio
                sh 'git clone https://github.com/rootbean/example_go_jenkins.git .'
            }
        }
        stage('Test') {
            steps {
                // Verificamos que estamos en una carpeta git
                sh 'git status'
                sh 'docker run --rm -v $(pwd):/app -w /app golang:1.22-alpine go test -v ./...'
            }
        }
    }
}