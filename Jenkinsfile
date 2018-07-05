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
}
