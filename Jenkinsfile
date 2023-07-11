pipeline {
    agent any
    
    stages {
        stage('Pull Repositories') {
            steps {
                echo 'Pull Repositories'
            }
        }

        stage('Download') {
            steps {
                echo 'Downloading...'
                sh 'go mod download'
            }
        }
        

        stage('Build') {
            steps {
                echo 'Building...'
                sh 'go build -o main .'
            }
        }

        stage('Build') {
            steps {
                echo 'Building...'
                sh './main'
            }
        }
        // stage('Pull Repositories') {
        //     steps {
        //         echo 'Hello, World!'
        //     }
        // }
        
        // stage('Stop Container') {
        //     steps {
        //         echo 'Stopping the running container...'
        //         sh 'docker stop mycontainer || true'
        //         sh 'docker rm mycontainer || true'
        //         echo 'Container stopped.'
        //     }
        // }
        
        // stage('Docker Images') {
        //     steps {
        //         echo 'Building Docker images...'
                
        //         sh 'docker rmi hello-world:latest || true'
                
        //         echo 'Proses Build'
        //         sh 'docker build -t hello-world:latest .'
        //         echo 'Menampilkan hasil images'
        //         sh 'docker images'
        //     }
        // }
        
        // stage('Deploy') {
        //     steps {
        //         echo 'Running the container...'
        //         sh 'docker run -d --name mycontainer -p 8080:8080 hello-world:latest'
        //         echo 'Container is now running.'
        //         sh 'docker ps'
        //     }
        // }
    }
}