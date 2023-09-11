pipeline {

    agent any

    /*tools { // Tools if we don't use wrappers
        jdk '<jdk-installation-name>'
        maven '<maven-installation-name>'
        gradle '<gradle-installation-name>'
    }*/

    stages {

        stage("Build") {

            steps {
                echo 'Building the application...'
                sh 'java -version' // For Linux
                // bat 'java -version' // For Windows
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

            steps {
                echo 'Testing the application...'
            }
        }

        stage("Deploy") {

            steps {
                echo 'Deploying the application...'
            }
        }
    }
}
