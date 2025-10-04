pipeline {
    agent {
        label 'windows && docker'
    }
    stages {
        stage('Build') {
            steps {
                pwsh './test.ps1'
            }
        }
    }
}
