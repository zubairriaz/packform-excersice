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


# Create .env file at the root of the project with the following variables

# Config

Environmental variables

- `PORT` - PORT of the postgress sql databse (default: `5432`)
- `PASSWORD` - Password of the postgress sql
- `USER` - User of the postgress sql
- `DBNAME` - Name of the DB by which you want the Db to be created or the already existent DB name

### Optional User Sign-In Feature

The application can be configured with an optional user sign-in feature which uses Azure Active Directory as an identity platform. This uses wrapper & helper libraries from https://github.com/benc-uk/msal-graph-vue

If you wish to enable this, carry out the following steps:

- Register an application with Azure AD, [see these steps](https://github.com/benc-uk/msal-graph-vue#set-up--deployment)
- Set the environmental variable `AUTH_CLIENT_ID` on the Go server, with the value of the client id. This can be done in the `.env` file if working locally.
- _Optional_ when testing/debugging the Vue.js SPA without the Go server, you can place the client-id in `.env.development` under the value `VUE_APP_AUTH_CLIENT_ID`

# GitHub Actions CI/CD

A working set of CI and CD release GitHub Actions workflows are provided `.github/workflows/`, automated builds are run in GitHub hosted runners

### [GitHub Actions](https://github.com/benc-uk/vuego-demoapp/actions)

[![](https://img.shields.io/github/workflow/status/benc-uk/vuego-demoapp/CI%20Build%20App)](https://github.com/benc-uk/vuego-demoapp/actions?query=workflow%3A%22CI+Build+App%22)

[![](https://img.shields.io/github/workflow/status/benc-uk/vuego-demoapp/CD%20Release%20-%20AKS?label=release-kubernetes)](https://github.com/benc-uk/vuego-demoapp/actions?query=workflow%3A%22CD+Release+-+AKS%22)


[![](https://img.shields.io/github/last-commit/benc-uk/vuego-demoapp)](https://github.com/benc-uk/vuego-demoapp/commits/master)

## Updates

| When       | What                                                 |
| ---------- | ---------------------------------------------------- |
| Nov 2021   | Rewrite for Vue.js 3, new look & feel, huge refactor |
| Mar 2021   | Auth using MSAL.js v2 added                          |
| Mar 2021   | Refresh, makefile, more tests                        |
| Nov 2020   | New pipelines & code/ API robustness                 |
| Dec 2019   | Github Actions and AKS                               |
| Sept 2019  | New release pipelines and config moved to env vars   |
| Sept 2018  | Updated with weather API and weather view            |
| July 2018  | Updated Vue CLI config & moved to Golang 1.11        |
| April 2018 | Project created                                      |