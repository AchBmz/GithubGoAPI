# GithubGoAPI
Github API with Golang
Simple Rest API to search Github repositories by coding language sorted by stars

# Backend
- Developed with Go to call Github API

### RUN Backend
1. Install go
    See https://golang.org/doc/install
2. Open the project folder
3. Run command 
    ```
    go run main.go
    ```
4. Open web-browser url : http://localhost:3000/?language=<language>

# Frontend
Developed with Angular to interact with the Backend server

### RUN Frontend
1. Install NodeJS
     See https://nodejs.org/en/
2. Install Angular CLI
    ```
    npm install -g @angular/cli
    ```
3. Open the project folder and install the npm packages
    ```
    npm install
    ```
4. Run the Application 
    ```
    ng serve -o
    ```
# Timeline
*Thursday 02-01-2020*
- Set up DEV environment and IDE (Visual Studio Code)
- First contact with Go : Documentation 

*Friday 03-01-2020*
- First 'Hello World' example with Go
- First static web service with Go
- Testing Github API with curl

*Saturday 04-01-2020*
- Implement the Github API
- Test the call for Github API
- Create the frontend project

*Monday 06-01-2020*
- Create Angular components
- Call Backend server
- Configure Backend to accept call from Frontend
- Display return on console

*Tuesday 06-01-2020*
- Work on Angular componenet to display returned values
- Edit Readme file