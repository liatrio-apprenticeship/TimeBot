pipeline {
    agent {
        label "lead-toolchain-skaffold"
    }
    environment {
        VERSION = version()
    }
    stages {
        stage('Build') {
            steps {
                container('skaffold') {
                    sh "docker build -t ${SKAFFOLD_DEFAULT_REPO}/timebot:${VERSION} ./flottbot/"
                    sh "docker push ${SKAFFOLD_DEFAULT_REPO}/timebot:${VERSION}"
                }
                
            }
        }
        stage('Deployment') {
            when {
                branch 'master'
            }
            steps {
                container('skaffold') {
                    sh "helm init --client-only"
                    sh "helm dependency update ./charts/timebot/"
                    withCredentials([string(credentialsId: 'jenkins-credential-slack', variable: 'SLACK_TOKEN'), string(credentialsId: 'jenkins-credential-slack-verification', variable: 'SLACK_VERIFICATION_TOKEN')]){
                        sh "helm upgrade timebot -i --set istioDomain=${env.productionDomain} --set config.slack_token=${SLACK_TOKEN} --set config.slack_verification_token=${SLACK_VERIFICATION_TOKEN} --set image.repository=${SKAFFOLD_DEFAULT_REPO}/timebot:${VERSION} --tiller-namespace ${productionNamespace} --namespace ${productionNamespace} ./charts/timebot/"
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
def version() {
    return sh(script: "git describe --tags --dirty", returnStdout: true).trim();
}