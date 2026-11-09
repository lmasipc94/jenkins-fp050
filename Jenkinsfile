pipeline {
  agent any
  stages {
    stage('Build') { steps { sh 'docker build -t web-uoc:$BUILD_NUMBER .' } }
    stage('Deploy') { steps { sh 'kubectl apply -f k8s/' } }
  }
}
