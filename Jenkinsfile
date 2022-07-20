def getEnvFromBranch(branch) {
    if (branch ==~ /^(main|edge_conductor_v\d+\.\d+\.\d+)$/) {
        return 'checkmarx,snyk,virus,protex'
    }
    else {
        return 'checkmarx,snyk,virus'
    }
}

pipeline {
    triggers {
    // nightly build between 23:00 a.m. - 23:59 a.m.(Etc/UTC), Monday - Friday:
    cron(env.BRANCH_NAME =~ /main|edge_conductor_v*/ ? 'H 23 * * 1-5' : '')
    }
    agent {
        label 'docker'
    }
    options { disableConcurrentBuilds() }
    environment {
        GOVERSION = "${sh(script:'cat build/Makefile | grep "GOVERSION =" | cut -d "=" -f2', returnStdout: true).trim()}"
    }
    stages {
        stage ('Check license headers') {
            steps {
                withCredentials([string(credentialsId: 'sys_pmce_git_github_token', variable: 'GITHUB_API_TOKEN')]) {
                    sh 'git clone https://${GITHUB_API_TOKEN}@github.com/intel-sandbox/applications.analyzers.infrastructure.license-header-checker.git header_scan'
                    sh '''
                    cd header_scan 
                    echo -e "Copyright (c) {dates} Intel Corporation.\n\nSPDX-License-Identifier: Apache-2.0" > license_header_template.txt
                    python3 license_header_checker.py check ../ -r
                    '''
                }
            }
        }
        stage('Go Pipeline') {
            environment {
                MAKE_TEST_LOG = '/tmp/make_test_step.log'
                FUNCTION_COV_FILE = 'cover.function-coverage.log'
                STATEMENT_COV_FILE = '/tmp/cover.statement-coverage.csv'
                HTML_COVERAGE_FILE = 'cover.html'
            }
            agent {
                docker {
                    image 'golang:' + env.GOVERSION
                    reuseNode true
                }
            }
            stages {
                stage('Prep') {
                    steps {
                        sh 'apt-get update -y && apt-get install git make unzip wget gcc docker bash gcc musl-dev curl -y'
                    }
                }
                stage('Build') {
                    steps {
                        sh 'make GO_BUILD_CMD=""'
                    }
                }
                stage('Test') {
                    steps {
                        // pipefail - the return value of a pipeline is the value of the last (rightmost) command to
                        // exit with a non-zero status or zero if all commands in the pipeline exit successfully
                        sh 'bash -c \'set -o pipefail ; make test GO_BUILD_CMD="" |& tee ${MAKE_TEST_LOG}\''
                    }
                }
                stage('Artifact') {
                    steps {
                        // copy relevant files of code coverage to /tmp/ dir because 'make artifact' will cleanup
                        // these files (git clean -f -d -X), and we need to upload them to the artifactory
                        sh '''
                            cat ${MAKE_TEST_LOG} | grep -E "FAIL|^ok|no test file" | grep -vE "_FAIL|^FAIL$" | \
                            awk -v OFS=',' 'BEGIN {print "Status","PKG","Time","Cov"} \
                            {if ($1=="?") {print $1,$2,$10,$3" "$4" "$5} \
                            else if ($1=="FAIL") {print $1,$2,$3,$10} \
                            else if ($1=="---") {print substr($2,1,length($2)-1),$3,$4,$10} \
                            else {print $1,$2,$3,$5}}' | sed 's/[()]//g' > ${STATEMENT_COV_FILE}
                        '''
                        sh 'mv ${FUNCTION_COV_FILE} ${HTML_COVERAGE_FILE} /tmp/'
                        sh 'rm -r header_scan'
                        sh 'make artifact GO_BUILD_CMD=""'
                        archiveArtifacts artifacts: '**/*.tar.gz',
                        allowEmptyArchive: true,
                        fingerprint: true,
                        onlyIfSuccessful: true
                        sh '''
                            [ -d "artifacts/ut-coverage" ] || mkdir -p artifacts/ut-coverage
                            mv  EdgeConductor-*.tar.gz artifacts/
                            mv /tmp/cover.* artifacts/ut-coverage/
                        '''
                        publishArtifacts([artifactsRepo: 'edge-conductor-or-local'])
                    }
                }
            }
        }
        stage('Scan'){
            environment {
                PROJECT_NAME               = 'edge-peak'
                SCANNERS                   = getEnvFromBranch(env.BRANCH_NAME)

                // protex details
                PROTEX_PROJECT_NAME        = 'edge_conductor_opensource_master'

                // publishArtifacts details
                ARTIFACT_RETENTION_PERIOD  = ''
                ARTIFACTORY_URL            = 'https://ubit-artifactory-or.intel.com/artifactory'
                ARTIFACTS_REPO             = 'edge-conductor-or-local'
                PUBLISH_TO_ARTIFACTORY     = true

                CHECKMARX_INCREMENTAL_SCAN = false
                CHECKMARX_PROJECT_NAME     = 'edge-peak-opensource'
                CHECKMARX_FORCE_SCAN       = true
                CHECKMARX_USER_AUTH_DOMAIN = 'GER'

                SNYK_MANIFEST_FILE         = 'go.mod'
                SNYK_GET_COMPLETE_REPORT   = true

                // .virus_scan in order to avoid the copy of the repo files into another scan steps
                VIRUS_SCAN_DIR             = '.virus_scan'
            }
            when {
                anyOf {
                    branch 'master';
                    branch 'main';
                    changeRequest();
                }
            }
            steps {
                script{
                    // move built sources in /tmp because next step will clean everything
                    // and this file is needed for McAfee virus/malware scan
                    sh 'mv artifacts/EdgeConductor-*.tar.gz /tmp/'
                    scmCheckout {
                        clean = true
                    }
                    // move the built sources in current workspace, in a 'hidden' dir,
                    // in order to not be copied into the workspace of the other scans with 'cp $WORKSPACE/* <some_dir>'
                    sh 'mkdir .virus_scan && mv /tmp/EdgeConductor-*.tar.gz .virus_scan/'
                }
                rbheStaticCodeScan()
            }
        }
        stage('Start validation') {
            when {
                anyOf {
                    branch pattern: "^(main|edge_conductor_v\\d+\\.\\d+\\.\\d+)\$", comparator: "REGEXP"
                }
            }
            steps {
                build job: '../Cluster-deployment/Kind-deployment',
                    parameters: [string(name: 'EC_SOURCES_REVISION', value: env.GIT_COMMIT)]
            }
        }
    }
    post {
        always {
            script {
                if (currentBuild.currentResult == 'FAILURE') {
                    emailext body: 'Check console output at $BUILD_URL to view the results. \n\n ${CHANGES} \n\n -------------------------------------------------- \n${BUILD_LOG, maxLines=100, escapeHtml=false}', 
                    to: "edge.platform.dev.bb@intel.com", 
                    subject: 'Build failed in Jenkins: $PROJECT_NAME - #$BUILD_NUMBER'
                }
            }
            cleanWs()
        }
    }
}
