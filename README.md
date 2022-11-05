# Go-JWT
## Steps to bring up the setup

#### Bring up containers
 ```docker-compose up```

#
### Postman collection
Import postman_collection.json to postman.
Create user, generate JWT and access secured resources

#
### Remote Debug 
Pre-requisite: VS-Code go plugin dlv (Go lang debugger)

Run ```docker-compose up```

VSCode: click on Run and Debug (```shift+command+D```)

Select "```Connect to Delve debug server```" and run. Now VSCode connected to remote dlv server running on ```40000``` port.

Add break point and send the rest api request to debug.