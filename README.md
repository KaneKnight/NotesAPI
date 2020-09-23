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