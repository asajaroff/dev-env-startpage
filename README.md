# dev-env-startpage

## Introduction
This program is a simple api and frontend that handles creation, deletion and credentials for the Cloud Development Enviroment project.

## Deployment
This software is distributed as a single `helm chart` that creates 2 deployment:
1.  A frontend nginx container
2.  An API that exposes the vcluster actions -such as list, create, destroy-

## Development
### Backend
#### Prerequisites
Take a look at the [go.mod](./backend/go.mod) file.

### Frontend