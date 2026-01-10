pipeline {
    agent {
        // Usamos una imagen de Go para no tener que instalar Go manualmente en Jenkins
        docker { image 'golang:1.25.5' }
    }
    stages {
        stage('Build') {
            steps {
                sh 'go build -o mi-app'
            }
        }
        stage('Test') {
            steps {
                // Ejecuta los tests y genera un reporte básico
                sh 'go test -v ./...'
            }
        }
    }
}