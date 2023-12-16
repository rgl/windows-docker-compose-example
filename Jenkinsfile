pipeline {
    agent {
        label 'windows && docker'
    }
    stages {
        stage('Build') {
            steps {
                powershell './test.ps1'
            }
        }
    }
}
