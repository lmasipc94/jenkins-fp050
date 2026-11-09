pipeline {
  agent any
  stages {
    stage('Build') { steps { sh 'docker build -t uoc-web:$BUILD_NUMBER .' } }
    stage('Deploy') { steps { sh 'kubectl apply -f k8s/' } }
  }
}
