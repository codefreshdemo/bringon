pipeline {
  agent any
  stages {
    stage('') {
      steps {
        codefreshRun(cfPipeline: 'gobot', cfBranch: 'master', cfVars: [['Value' : "${BUILD_NUMBER}", 'Variable' : 'BUILD_NUMBER']] )
      }
    }
  }
}
