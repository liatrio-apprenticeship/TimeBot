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
                    withCredentials([string(credentialsId: 'jenkins-credential-slack', variable: 'SLACK_TOKEN'),string(credentialsId: 'jenkins-credential-slack-verification', variable: 'SLACK_VERIFICATION_TOKEN')]){
                        // withCredentials([string(credentialsId: 'jenkins-credential-slack-verification', variable: 'SLACK_VERIFICATION_TOKEN')]){
                            sh "helm upgrade timebot -i --set istioDomain=${env.productionDomain} --set config.slack_token=${SLACK_TOKEN} --set config.slack_verification_token=${SLACK_VERIFICATION_TOKEN} --set image.repository=${SKAFFOLD_DEFAULT_REPO}/timebot:lastest --tiller-namespace ${productionNamespace} --namespace ${productionNamespace} ./charts/timebot/"
                        // }
                    }
                }
            }
            post {
                failure {
                    container('skaffold'){
                        sh "helm --tiller-namespace ${productionNamespace} delete --purge timebot"
                    }
                }
            }
        }
    }
    
}