# ThirdfortAPI


## Setup

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


To launch the mongod service:
```
sudo service mongod start
```

Note: For different operating systems this process will be different.

To install Go:

See here: https://golang.org/doc/install

To launch the notes backend run the command:
```
./run.sh
```



## API Description
To see the API documentation head to this link: https://documenter.getpostman.com/view/5708033/TVKEXcxH

The way the API works is to first register a user with an email and password. Next, a user wanting to interact with the api must login by sending their email and password to the `/login` endpoint. This will return an authentication token and a user id. These fields are required for all other api requests as the headers `X-Auth-Token` and `X-User-Id`.

## Technical Choices
I chose to use Go for the backend implementation because I have used it before and I didn't want to spend too much time getting acustomed to a new framework.  On top of that, the Go language is typed and can help me spot errors at compile time, otherwise I would of had to wait to hit the endpoint to find the type errors.

I chose to use MongoDB because it is a no sql framework. When I began designing the API I realised the notes themselves wouldn't need to store many fields, so I thought the overhead of designing an sql schema although not difficult, would yield little benefit. Futher, I didn't want to spend too much time implementing a user system as this has been done a lot, and I found this library (https://github.com/yasaricli/gah) which links MongoDB and Go so it was a natural fit.

## Extensions
If I were to continue, I need to add a https connection for the login and register as plain text passwords are being sent, which could be sniffed by a man in the middle. I would also want to learn how to mock the database connection so I can unit test my handlers.