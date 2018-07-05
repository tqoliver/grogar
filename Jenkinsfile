#!groovy
  
def project = 'ccdemoqa'
def appName = 'serenity-api'
def feSvcName = "${appName}"
def namespace = 'ccdemoqa'
def imageTag = "timothyoliver/serenity-api"
def prevImageTag = ''
def prevBuildNum = ''
def firstDeploy = false

node {
stage 'Build in QA'
openshiftBuild(namespace:'ccdemoqa', buildConfig: 'serenity-api', showBuildLogs: 'true')
stage 'Deploy to QA'
openshiftDeploy(namespace: 'ccdemoqa', deploymentConfig: 'serenity-api')
openshiftScale(namespace: 'ccdemoqa', deploymentConfig: 'serenity-api', replicaCount: '2')
stage 'Deploy to Production'
input 'Promote QA Image to Production?'
openshiftTag(namespace: 'ccdemoqa', soruceStream: 'serenity-api', sourceTag: 'latest', destinationNamespace: 'ccdemoprod', destinationStream: 'serenity-api', destionationTag: 'promote-prod')
openshiftDeploy(namespace: 'ccdemoprod', deploymentConfig: 'serenity-api')
openshiftScale(namespace: 'ccdemoprod', deploymentConfig: 'serenity-api', replicaCount: '2')
}
