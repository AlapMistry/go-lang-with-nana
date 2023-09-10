pipeline {

    agent any

    tools { // tools if we don't use wrappers
        jdk 'JDK-17'
        maven 'Maven-3.9.4'
        gradle 'Gradle-8.3'
    }

    stages {

        stage("build") {

            steps {
                echo 'building the application...'
                sh 'java -version'
                sh 'maven -version'
                sh './gradlew -v'
                echo 'application built'
                /*nodejs('<nodejs-installation-name>') { // Wrappers
                    sh '<command>'
                }
                withGradle() {
                    sh '<command>'
                }
                */
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
