pipeline {

    agent any

    /*tools { // tools if we don't use wrappers
        jdk '<jdk-installtion-name>'
        maven '<maven-installation-name>'
        gradle '<gradle-installation-name>'
    }*/

    stages {

        stage("build") {

            steps {
                echo 'building the application...'
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
