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
                    sh "helm init --client-only"
                    sh "helm dependency update ./charts/timebot/"
                    sh "helm upgrade --install --name timebot --set image.repository=${SKAFFOLD_DEFAULT_REPO}/timebot:lastest --tiller-namespace ${productionNamespace} --namespace ${productionNamespace} ./charts/timebot/"
                }
            }
        }
    }
    
}