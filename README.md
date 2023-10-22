# SomethingSocial
Prototype of simple social network I made to practice in using
different backend technologies such as gRPC; and development practices
such as Clean Architecture, CQRS, etc.

## Features
* Sign up / sign in with username and password
* Authorization using JWT
* Get info about users:
  * Id
  * Name
  * Karma
  * Registration date
  * Last login time
  * Posts count
  * Follwers / following count
  * Who follows them / who they follow
* Follow / unfollow other users
* Create / edit / delete posts
* Upvote / downvote posts (affects post's and author's karma)
* Get feed (posts from users that you follow)
* Get posts:
  * Filter by:
    * Author
    * Time period
    * Whether you upvoted/downvoted it
  * Sort by:
    * Post date
    * Karma
  * Limit number of posts

## Prerequisites
* Docker compose
* [Goose](https://github.com/pressly/goose)

## Run
1. `make up`

## Technology stack
* Go (chi, sqlx, goose, go-grpc, oapi-codegen, logrus)
* gRPC
* PostgreSQL
* Docker