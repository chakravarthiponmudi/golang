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

```
 docker run -d -p 80:80 \
-e "PGADMIN_DEFAULT_PASSWORD=password" \
-e "PGADMIN_DEFAULT_EMAIL=chakravarthiponmudi@gmail.com" --net setup_infranet \dpage/pgadmin4
bck-i-search: dpage_

```

# Tools and other Details
### golang dep

This is the tool that i used for golang dependency management

