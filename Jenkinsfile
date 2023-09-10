pipeline {

    agent any

    /*tools { // tools if we don't use wrappers
        jdk '<jdk-installation-name>'
        maven '<maven-installation-name>'
        gradle '<gradle-installation-name>'
    }*/

    stages {

        stage("build") {

            steps {
                echo 'building the application...'
                sh 'java -version'
                /*withMaven(maven: '<maven-installation-name>') { // Wrappers - Install Pipeline Maven Integration Plugin
                    sh 'mvn -v'
                }
                nodejs('<nodejs-installation-name>') { // Install NodeJS plugin
                    sh '<command>'
                }
                withGradle() { // Gradle Plugin
                    sh '<command>'
                }
                */
                echo 'application built'
            }
        }

        stage("test") {

            steps {
                echo 'testing the application...'
            }
        }

        stage("deploy") {

            steps {
                echo 'deploying the application...'
            }
        }
    }
}
