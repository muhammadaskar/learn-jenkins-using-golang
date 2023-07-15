pipeline {
    agent any
    
    stages {
        stage('Pull Repositories') {
            steps {
                echo 'Hello, World Koding!'
            }
        }
        
        // stage('Stop Container') {
        //     steps {
        //         echo 'Stopping the running container...'
        //         sh 'docker stop golang-container || true'
        //         sh 'docker rm golang-container || true'
        //         echo 'Container stopped.'
        //     }
        // }
        
        // stage('Docker Images') {
        //     steps {
        //         echo 'Building Docker images...'
                
        //         // Menghapus image sebelumnya
        //         sh 'docker rmi golang-images:latest || true'
                
        //         echo 'Proses Build'
        //         sh 'docker build -t golang-images:latest .'
        //         echo 'Menampilkan hasil images'
        //         sh 'docker images'
        //     }
        // }
        
        stage('Deploy') {
            steps {
                echo 'Running the compose...'
                sh 'docker compose up'
            }
        }
    }
}