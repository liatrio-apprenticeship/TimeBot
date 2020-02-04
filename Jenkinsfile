pipeline {
    agent {
        label "lead-toolchain-skaffold"
    }
    stages {
        stage('Build') {
            steps {
                container('skaffold') {
                    sh "docker build -t ${SKAFFOLD_DEFAULT_REPO}/timebot:lastest ./flottbot/"
                    sh "docker push ${SKAFFOLD_DEFAULT_REPO}/timebot:lastest"
                }
                
            }
        }
    }
    
}