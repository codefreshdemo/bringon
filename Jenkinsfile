pipeline {
  agent any
  stages {
    stage('CallCF') {
      steps {
        codefreshRun(cfPipeline: 'bringon', cfBranch: 'master', cfVars: [['Value' : "${BUILD_NUMBER}", 'Variable' : 'BUILD_NUMBER']] )
      }
    }
  }
}
