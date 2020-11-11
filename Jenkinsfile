pipeline {
	agent any
	tools{
	    go 'go'
	    dockerTool 'docker'
	}

	stages {
	    
	  	stage("drivers") {
        	steps {
            	sh "go get -u github.com/go-sql-driver/mysql"
        	}
    	}
    	
    	stage("Git"){
        	steps {
            	git url: "https://github.com/rogerverges/registration-service"
        	}
	    }
	   
   
        stage('Build go') {
        	steps {
            	sh 'go build post.go '
        	}
    	}
    	stage('docker build image'){
    	    steps{
    	        sh 'docker build -t post .'
				sh 'docker tag post 192.162.133.158:5000/post:latest'
				sh 'docker push 192.162.133.158:5000/post:latest'
    	    }
    	}
    }
}