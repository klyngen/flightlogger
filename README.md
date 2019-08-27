# A better flightlogger
flightlog.org is a great site. This project is a modern implementation of that webpage with a similar data-model, written in golang

# Architecture
Here we use a layered model in order to ensure security. The backend API will basically have two layers

### Frontend
- Angular Cli (coming soon)

### Backend
- Presentation-layer
- Datalayer
- (ORM library)

# Project setup
The backend is really a simple golang project. You need golang installed as well as **[dep](https://github.com/golang/dep)**. After dependencies do a `go get github.com/klyngen/flightlogger` followed by `dep ensure`. 

# Contributions
All help is appreciated. Send me an email if you wonder how you can contribute or just make a PR. 

## Guidelines
- Ensure that the API is easy to use. Please reed [this guide](https://blog.florimond.dev/restful-api-design-13-best-practices-to-make-your-users-happy)
- **Write tests** the application is layered for a reason....
- Be nice :) 



**Check out our board on [taiga.io](https://tree.taiga.io/project/klyngen-better-flightlog/timeline)**
