# 5-goprojects
The goal of this repo is to get started using go in a real scenario and using go's standard library as well as a few popular external packages

## ToDo list
    To set path if go install not working:

    go env to see path

    export GOPATH='/Users/tylerjob/go/bin' 
    export PATH=$PATH:$(go env GOPATH)/bin

    a cli application that allows a user to add, delete, update, list, and list all to do tasks
    - cobra for building cli applications
	- encoding/csv
	- os
	- strconv
	- syscall
	- time

    []   text/tabwriter for writing out tab aligned output
    []   github.com/mergestat/timediff for displaying relative friendly time differences (1 hour ago, 10 minutes ago, etc)

## Backend API
    a web server that acts as a calulator with add, subtract, divide, mulitply, and sum endpoints
    - net/http

    - Extras
        []  Add in rate limiter to prevent misuse of the API
        []  Add in token authentication to prevent anyone unauthorized from using the API
        []  Add in a database to keep track of all of the calculations that have taken place
        []  Add in support for floating point numbers as well.
        []  Create an associated http client that can work with the calculator API.
        []  Create a frontend that makes use of your API.
        []  Add in a middleware that adds a request ID to the http.Request object.

## Web Scraper
    a web scraper that checks links and sublinks for dead links (ie 400 vs 200). Also using go routines
    - net/url
	- strings
	- sync
    - golang.org/x/net/html

## URL Shortener