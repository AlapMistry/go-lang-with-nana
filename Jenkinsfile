pipeline {

    agent any

    parameters {
        string(name: 'VERSION_S', defaultValue: '', description: 'Version value')
        choice(name: 'VERSION_C', choices: ['1.0.0', '1.1.0', '1.2.0'], description: 'Version choice')
        booleanParam(name: 'executeTests', defaultValue: true, description: 'Boolean value')
    }

    environment {
        NEW_VERSION = '1.0.0'
        SERVER_CREDENTIALS = credentials('server-credentials')
    }

    /*tools { // Tools if we don't use wrappers
        jdk '<jdk-installation-name>'
        maven '<maven-installation-name>'
        gradle '<gradle-installation-name>'
    }*/

    stages {

        stage("Build") {

            steps {
                echo "Building the application with version ${NEW_VERSION}"
                // sh 'java -version' // For Linux
                bat 'java -version' // For Windows
                /*withMaven(maven: '<maven-installation-name>') { // Wrappers - Install Pipeline Maven Integration Plugin
                    sh 'mvn -v' // For Linux
                    bat 'mvn -v' // For Windows
                }
                nodejs('<nodejs-installation-name>') { // Install NodeJS plugin
                    sh '<command>' // For Linux
                    bat '<command>' // For Windows
                }
                withGradle() { // Gradle Plugin
                    sh '<command>' // For Linux
                    bat '<command>' // For Windows
                }
                */
                echo 'Application built'
            }
        }

        stage("Test") {

            when {
                expression {
                    // BRANCH_NAME == 'dev' || BRANCH_NAME == 'feature-1'
                    // params.executeTests == true
                    params.executeTests
                }
            }

            steps {
                echo 'Testing the application...'
            }
        }

        stage("Deploy") {

            steps {
                withCredentials([ // If we don't use environment variable for it
                    usernamePassword(credentialsId: 'server-credentials', // For credential type username with password
                    usernameVariable: 'USERNAME',
                    passwordVariable: 'PASSWORD')
                ]) {
                    echo "${USERNAME}"
                    echo "${PASSWORD}"
                }

                echo "Deploying the application with ${SERVER_CREDENTIALS}"
                echo "Deploying the application with version ${params.VERSION_S}"
                echo "Deploying the application with version ${params.VERSION_C}"
                // sh "${SERVER_CREDENTIALS}" // For Linux
                bat "echo ${SERVER_CREDENTIALS}" // For Windows
            }
        }
    }

    post {

        always {
            echo 'Always'
        }

        success {
            echo 'Sucess'
        }

        failure {
            echo 'Failure'
        }
    }
}
