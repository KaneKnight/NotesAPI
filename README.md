# ThirdfortAPI

The stack is a Golang backend with a MongoDB database. Therefore to start the service, Mongodb and Go are needed.

To install mongodb on ubuntu 18.04 (via https://docs.mongodb.com/manual/tutorial/install-mongodb-on-ubuntu/):

```
wget -qO - https://www.mongodb.org/static/pgp/server-4.4.asc | sudo apt-key add -
sudo apt-get install gnupg
wget -qO - https://www.mongodb.org/static/pgp/server-4.4.asc | sudo apt-key add -
echo "deb [ arch=amd64,arm64 ] https://repo.mongodb.org/apt/ubuntu bionic/mongodb-org/4.4 multiverse" | sudo tee /etc/apt/sources.list.d/mongodb-org-4.4.list
sudo apt-get update
sudo apt-get install -y mongodb-org
```

To install Go:

See here: https://golang.org/doc/install


To see the API documentation head to this link: https://documenter.getpostman.com/view/5708033/TVKEXcxH

The way the API works is to first register a user with an email and password. Next, a user wanting to interact with the api must login by sending their email and password to the `/login` endpoint. This will return an authentication token and a user id. These fields are required for all other api requests as the headers `X-Auth-Token` and `X-User-Id`.