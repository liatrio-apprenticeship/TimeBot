pipeline {
    agent {
        label "lead-toolchain-skaffold"
    }
    stages {
        stage('Build') {
            steps {
                sh "docker build -t ${SKAFFOLD_DEFAULT_REPO}/TimeBot:lastest ./flottbot/"
                sh "docker push ${SKAFFOLD_DEFAULT_REPO}/TimeBot:lastest"
                
            }
        }
    }
    
}