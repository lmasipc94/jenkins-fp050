pipeline {
  agent any
  stages {
    stage('Build') {
      steps {
        sh 'docker build -t miapp:$BUILD_NUMBER -t miapp:latest .'
        sh 'minikube image load miapp:latest'
      }
    }
    stage('Deploy') {
      steps {
        sh 'kubectl apply -f k8s/'
        sh 'kubectl rollout restart deployment/miapp'
      }
    }
  }
}
