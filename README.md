## Go & Vue.js - Demo Web Application

This is a simple web application with a Go server/backend and a Vue.js SPA (Single Page Application) frontend.


## Repo Structure

```txt
/
├── frontend         Root of the Vue.js project
│   └── src          Vue.js source code
│   └── tests        Unit tests
├── server           Go backend server
│   └── main.go      Server main / exec
│   └── main_test.go Test for testing the API 
├── seeder.py         Python script for downloading data files and creating tables and inserting data
```


## Building & Running Locally

### Pre-reqs

- Be using Linux, WSL or MacOS, with bash, make etc
- [Node.js](https://nodejs.org/en/) [Go 1.16+](https://golang.org/doc/install) - for running locally, linting, running tests etc
- [Postgress](https://www.postgresql.org/download/windows/)
- [Python] for running the seeder.py file


Clone the project to any directory where you do development work

```
git clone https://github.com/zubairriaz/packform-excersice
```


##### 1 - There is a .env file at the root of the project with the following variables

##### Config

Environmental variables

- `PORT` - PORT of the postgress sql databse (default: `5432`)
- `PASSWORD` - Password of the postgress sql
- `USER` - User of the postgress sql
- `DBNAME` - Name of the DB by which you want the Db to be created or the already existent DB name

   #### Assign appropriate values to the variables

##### 2 - Download python dependencies by running `pip install -r requirements.txt`

##### 3 - Run seeder.py file by executing the command `python seeder.py` in the terminal

##### 4 - Now Open the Server folder and download all dependencies using the command `go get -d ./...`

##### 5 - Now go the the frontend folder and download all dependencies by executing the command `npm install`

##### 6- Now start the server by executing the command `go run main.go app.go` from the server folder this will start the server at http://localhost:8001

##### 7- Now start the front end by executing the command `npm run server`

##### 8- This will start the project at http://localhost:8080/ but due to google chrome security issues we cannot access the API from our front end due to CORS issue so to test the app close all chrome windows and follow these steps
1-Right click on desktop, add new shortcut
2-Add the target as "[PATH_TO_CHROME]\chrome.exe" --disable-web-security --disable-gpu --user-data-dir=~/chromeTemp
3-Click OK.

##### 9- Now open the chrome from the shortcut you have created on desktop and access http://localhost:8080/







   


