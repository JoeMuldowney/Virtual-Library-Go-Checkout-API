pipeline {
    agent { label 'gopipe' }


environment {
       DOCKERHUB_CREDENTIALS = d324a1f1-35d3-4aca-956a-334b18fb2852
}
stages {

        steps('Remove old build'){
            steps{
            // Stop and remove current docker container to free up space
                sh 'docker ps -q --filter name=Golibrarybackend | xargs -r docker stop'
                sh 'docker ps -q --filter name=Golibrarybackend | xargs -r docker rm'
                sh 'docker system prune -a'
            }
        }
        stage('Build') {
                    steps {
                        // Checkout source code
                        checkout scm

                        // Build Docker image
                        sh 'docker build -t Golibrarybackend .'
                    }
                }
        stage('Push to Docker Hub') {
             steps {

                withCredentials([usernamePassword(credentialsId: DOCKERHUB_CREDENTIALS, usernameVariable: 'DOCKERHUB_USERNAME', passwordVariable: 'DOCKERHUB_PASSWORD')]){

                    sh 'docker login -u $DOCKERHUB_USERNAME -p $DOCKERHUB_PASSWORD'

                    // Tag Docker image
                    sh 'docker tag Golibrarybackend joemuldowney/virtual_library_cart'

                    // Push Docker image to Docker Hub
                    sh 'docker push joemuldowney/virtual_library_cart'
                }
             }
        }
        stage('deploy on server'){

                 steps {
                 sh 'docker run -d -p 8020:8020 --name Golibrarybackend joemuldowney/virtual_library_auth'
                 }

        }
}