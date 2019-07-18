# Rate Limiter App

In this app, i will design a concept for the rate limiter app.

This file is supposed to contain about my journey how this all started and how i proceeded. This is just to remind me how i am progressing with this task.


## 11-Jul-2019
Today is the starting of this app. I have decided to do this app in the GO Lang.

The design I am planning to adopt is dynamic sliding with counter approach.

So the first thing i will do is to how to receive the request and forward the request to some other system.

In the first phase after going through i decided to implement a proxy type server as the first step.

## 13-Jul-2019

I am now focusing on the registration side of the API.

Created the database via docker and made a sure that the following are working

I have the endpoints for adding the contract and the api's but with out proper error handling.

Now i have to setup the error handling processing and also should focus on the test case writing

```
 docker run -d -p 80:80 \
-e "PGADMIN_DEFAULT_PASSWORD=password" \
-e "PGADMIN_DEFAULT_EMAIL=chakravarthiponmudi@gmail.com" --net setup_infranet \dpage/pgadmin4
bck-i-search: dpage_
```

## 17-Jul-2019

got a  working test case of registration module.

Now i have to understand the db modules used, that returns the error codes and handle them properly in the registration module

# Tools and other Details
### golang dep

This is the tool that i used for golang dependency management

###GORC
gorc is a command that helps running the go test command recursively inside the subfolders. it seems that go test doesn't run the test cases which are present inside the sub folders. **CHECK**

###TESTIFY

Testify is a like MOCHA. which allows you to do some tearup and teardown.... good to have feature.