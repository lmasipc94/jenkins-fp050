pipeline {
  agent any
  stages {
    stage('Builds') { steps { sh 'docker build -t miapp:$BUILD_NUMBER .' } }
    stage('Deploys') { steps { sh 'kubectl apply -f k8s/' } }
  }
}
