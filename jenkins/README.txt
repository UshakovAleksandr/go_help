1. docker pull jenkins/jenkins:lts

2. docker run -d \
  -p 8888:8080 \
  --restart always \
  --env DOCKER_TLS_CERTDIR=/certs \
  --volume jenkins-docker-certs:/certs/client \
  --volume jenkins-data:/var/jenkins_home \
  --name jenkins \
  jenkins/jenkins:2.332.3-lts-jdk11

  Locale
  Active Choices
  Nexus Platform Plugin
  Go Plugin
  Docker Pipeline
  Docker

3. localhost:8080. Там при первом запуске будет указан путь до пароля. Смотрим его в терминале командой:
   docker exec "jenkins" cat /var/jenkins_home/secrets/initialAdminPassword

4. docker exec -it -u 0 jenkins bash && chmod 666 /var/run/docker.sock

4. docker pull sonatype/nexus3:3.38.1

5. docker run -d \
  -p 8889:8081 \
  -p 8085:8085 \
  --name nexus \
  --volume nexus-data:/nexus-data \
  sonatype/nexus3:3.38.1

6. docker exec nexus cat /nexus-data/admin.password

   pipeline {
    agent any
    
    tools {
        go 'go_1.19.5'
    }

    stages {
        stage('Go Version Test') {
            steps {
                sh 'go version'
            }
        }
    }
}

{
    "insecure-registries" : ["172.22.0.3:8085"]
}

https://bhairavisanskriti.hashnode.dev/building-a-ci-pipeline-for-a-web-application-using-jenkins

docker login -u admin -p admin 127.0.0.1:8085

docker tag hello:latest 127.0.0.1:8085/my_hello:latest

docker push 127.0.0.1:8085/my_hello:latest