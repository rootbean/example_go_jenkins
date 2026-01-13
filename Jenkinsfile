pipeline {
    agent {
      
        docker { image 'golang:1.25.5-alpine' }
    }
    environment {
        // Nombre de la imagen
        IMAGE_NAME = "mi-app-go"
        // Tag basado en el número de build para no sobreescribir siempre
        TAG = "${env.BUILD_NUMBER}"
    }
    stages {
        stage('Limpieza y Checkout') {
            steps {
                deleteDir() 
                checkout scm
            }
        }

        stage('Preparar Entorno') {
            steps {
                script {
                    // Instala Docker y Git dentro del contenedor de Jenkins si no existen
                    sh '''
                        if ! command -v docker &> /dev/null; then
                            echo "Instalando Docker..."
                            apk add --no-cache docker-cli
                        fi
                        if ! command -v git &> /dev/null; then
                            echo "Instalando Git..."
                            apk add --no-cache git
                        fi
                        if [ -f /etc/alpine-release ]; then
                            apk add --no-cache docker-cli git curl
                        else
                            apt-get update && apt-get install -y docker.io git curl
                        fi
                    '''
                }
            }
        }
      
        stage('Test') {
            steps {
                // 1. Construimos la imagen de test usando el nuevo archivo
                sh 'docker build -t ruber-go-tester -f Dockerfile.test .'
                
                // 2. Ejecutamos el comando de test
                // Esto garantiza que el go.mod esté presente y la versión sea la correcta
                sh 'docker run --rm ruber-go-tester go test -v .'
                
            }
        }

        stage('Build Image') {
            steps {
                sh "docker build -t ${IMAGE_NAME}:${TAG} ."
                sh "docker tag ${IMAGE_NAME}:${TAG} ${IMAGE_NAME}:latest"
            }
        }

        stage('Deploy') {
            steps {
                script {
                    echo "Iniciando despliegue de la versión ${TAG}..."
                    
                    // 1. Detener y borrar el contenedor si ya existe uno corriendo
                    // El "|| true" es para que no falle si el contenedor no existe la primera vez
                    sh "docker stop my-running-app 2>/dev/null || true"
                    sh "docker rm my-running-app 2>/dev/null || true"
                    
                    // 2. Ejecutar el nuevo contenedor
                    // -d lo corre en segundo plano, -p mapea el puerto (ej. 8080 de la app al 9090 del Mac)
                    sh "docker run -d --name my-running-app -p 9091:8089 ${IMAGE_NAME}:${TAG}"
                    
                    echo "✅ Despliegue completado. App disponible en http://localhost:9091"
                }
            }
        }

        stage('Smoke Test') {
            steps {
                script {
                    echo "Verificando que la app responda..."
                    // Esperamos unos segundos a que el servidor de Go levante
                    sleep 5
                    
                    // Intentamos conectar al puerto 9091 (host de Jenkins/Mac)
                    // Usamos la IP especial 'host.docker.internal' para salir del contenedor de Jenkins hacia el Mac
                    def response = sh(script: "curl -s http://host.docker.internal:9091", returnStatus: true)
                    
                    if (response == 0) {
                        echo "✅ ¡Prueba exitosa! La aplicación responde correctamente."
                    } else {
                        error "❌ La aplicación no responde en el puerto 9091."
                    }
                }
            }
        }
    }
}