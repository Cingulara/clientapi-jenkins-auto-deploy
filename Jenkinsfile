def CONTAINER_NAME="client-api-pipeline"
def CONTAINER_TAG="latest"
def CONTAINER_DB_NAME="client-api-db-pipeline"
def CONTAINER_DB_TAG="latest"
def OPENSHIFT_PROJECT_NAME="client-api-pipeline"
def OPENSHIFT_IP="172.30.1.1:5000"

node {

    stage('Checkout') {
        checkout scm
    }

    stage("Compile Source"){
        try {
            bat "set CGO_ENABLED=0"
            bat "set GOOS=linux"
            bat "c:\\go\\bin\\go get -u github.com/go-sql-driver/mysql"
            bat "c:\\go\\bin\\go get -u github.com/gorilla/mux"
            bat "c:\\go\\bin\\go build -a -installsuffix cgo -o client-api-pipeline ."
        } catch(error){
            echo "The go compile failed with ${error}"
        }
    }

    stage('SonarQube'){
        try {
            bat "c:\\projects\\sonar-scanner\\bin\\sonar-scanner -Dsonar.projectKey=clientapi -Dsonar.sources=."
        } catch(error){
            echo "The sonar server could not be reached: ${error}"
        }
     }

    stage("Image Prune"){
        imagePrune(CONTAINER_NAME)
        imagePrune(CONTAINER_DB_NAME)
    }

    stage('DB Image Build'){
        imageBuild(CONTAINER_DB_NAME, CONTAINER_DB_TAG, "database\\")
    }

    stage('Application Image Build'){
        imageBuild(CONTAINER_NAME, CONTAINER_TAG, ".")
    }

    stage('Push DB to OpenShift Registry'){
        withCredentials([usernamePassword(credentialsId: 'openshift-docker-registry-account', usernameVariable: 'USERNAME', passwordVariable: 'PASSWORD')]) {
            pushToImage(OPENSHIFT_IP,OPENSHIFT_PROJECT_NAME, CONTAINER_DB_NAME, CONTAINER_DB_TAG, USERNAME, PASSWORD)
        }
    }

    stage('Push App to OpenShift Registry'){
        withCredentials([usernamePassword(credentialsId: 'openshift-docker-registry-account', usernameVariable: 'USERNAME', passwordVariable: 'PASSWORD')]) {
            pushToImage(OPENSHIFT_IP,OPENSHIFT_PROJECT_NAME, CONTAINER_NAME, CONTAINER_TAG, USERNAME, PASSWORD)
        }
    }
}

def imagePrune(containerName){
    try {
        bat "\"C:\\Program Files\\Docker\\Docker\\Resources\\bin\\docker\" image prune -f"
        bat "\"C:\\Program Files\\Docker\\Docker\\Resources\\bin\\docker\" stop $containerName"
    } catch(error){}
}

def imageBuild(containerName, tag, pathToDockerfile){
    bat "\"C:\\Program Files\\Docker\\Docker\\Resources\\bin\\docker\" build -t $containerName:$tag --pull --no-cache $pathToDockerfile"
    echo "Image build complete"
}

def pushToImage(openshiftIP, projectName, containerName, tag, dockerUser, dockerPassword){
    bat "\"C:\\Program Files\\Docker\\Docker\\Resources\\bin\\docker\" login -u $dockerUser -p $dockerPassword $openshiftIP"
    bat "\"C:\\Program Files\\Docker\\Docker\\Resources\\bin\\docker\" tag $containerName:$tag $openshiftIP/$projectName/$containerName:$tag"
    bat "\"C:\\Program Files\\Docker\\Docker\\Resources\\bin\\docker\" push $openshiftIP/$projectName/$containerName:$tag"
    echo "Image push complete to OpenShift"
}
