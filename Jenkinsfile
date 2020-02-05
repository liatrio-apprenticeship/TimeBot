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
        stage('Deployment') {
            steps {
                container('skaffold') {
                    sh "helm dependency update"
                    sh "helm install --name timebot --set image.repository=${SKAFFOLD_DEFAULT_REPO}/timebot:lastest --namespace ${productionNamespace} ."
                }
            }
        }
    }
    
}