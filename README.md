# Client API using Golang and MySQL
This is a single Golang API microservice with MySQL backend as a POC. It is to show how to do a single API around a particular domain with a database that is its own in a separate programming language and database server. This requires Docker, Golang 1.10.x (or latest), and a web browser to run. Optionally you can use a SQL tool such as MySQL Workbench to view the database if you so desire.

Optionally, point Jenkins to the Jenkinsfile and setup your environment to match. You will need a token for a user to go into Jenkins for the authentication to OpenShift.

## Option A: Run the API with Docker Compose
* run 'docker-compose up -d' to run this as a networked set of containers
* modify the main.go if you want to separate these and run independently for localhost:3306
* Note: you need to see Option B: Golang application container below to see how to compile the Golang API

## Option B: Database Container

* run the 'docker build -t client-api-db-pipeline .' from within the database directory of client-api-pipeline
* run 'docker run -d --name client-api-db-pipeline --rm -p 3306:3306  client-api-db-pipeline' to launch the db individually

# Option B: Golang application container
* run 'set CGO_ENABLED=0' (or export)
* run 'set GOOS=linux' (or export)
* run 'go build -a -installsuffix cgo -o client-api-pipeline .' to build a go application self contained for a small image
* run 'docker run -it -p 8080:8080 --rm --name client-api-pipeline client-api-pipeline' to run the API, substituting the external port for what you require (it runs 8080 internal)

## Database Structure for MySQL
```
CREATE TABLE clients (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    address VARCHAR(300) NOT NULL,
    city VARCHAR(100) NOT NULL,
    state VARCHAR(50) NOT NULL,
    zip VARCHAR(10) NOT NULL,
    web VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    phone VARCHAR(12) NOT NULL,
    created VARCHAR(100) NOT NULL
);
```

## API Calls

GET http://localhost:xxxx/api/clients/ gets back a JSON listing of the Clients class.

GET http://localhost:xxxx/api/clients/1 gets back a JSON listing of the Client class for the first record. You must add a record to view anything.

POST http://localhost:xxxx/api/clients will add this client with the information below
```
{
	"Name": "KBRwyle", 
	"Address": "22309 Exploration Drive", 
	"City": "Lexington Park", 
	"State": "Maryland", 
	"ZIP": "20619", 
	"Web": "https://www.kbrwyle.com", 
	"Email": "info@wyle.com", 
	"Phone": "301-555-1212"
}
```
PUT http://localhost:xxxx/api/clients will update this client with the information below
```
{
    "Id" : 1,
	"Name": "KBRwyle", 
	"Address": "22309 Exploration Drive", 
	"City": "Lexington Park", 
	"State": "Maryland", 
	"ZIP": "20619", 
	"Web": "https://www.kbrwyle.com", 
	"Email": "info@wyle.com", 
	"Phone": "301-555-1212"
}
```

## ToDo's

* auto-generate swagger API docs
* Log events to an event stream