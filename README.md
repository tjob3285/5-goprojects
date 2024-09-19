# 5-goprojects

https://github.com/dreamsofcode-io/goprojects

https://go.dev/doc/effective_go

To set path if go install not working:

go env to see path

export GOPATH='/Users/tylerjob/go/bin' 
export PATH=$PATH:$(go env GOPATH)/bin

## ToDo list

    [x]  Create CSV file with ID, Description, CreatedAt, isComplete
    [x]  Add ToDo list item with description and created at
    [x]  Complete "id param" marks isComplete as true
    [x]  List shows all not complete tasks, flag -a shows all
    [x]  Delete by id

    [x]   encoding/csv for writing out as a csv file
    [x]   strconv for turning types into strings and visa versa
    []   text/tabwriter for writing out tab aligned output
    [x]   os for opening and reading files
    [x]   github.com/spf13/cobra for the command line interface
    []   github.com/mergestat/timediff for displaying relative friendly time differences (1 hour ago, 10 minutes ago, etc)

## Backend API

    [x]  Add
    [x]  Subtract
    [x]  Multiply
    [x]  Divide
    [x]  Sum

    - Extras
        []  Add in rate limiter to prevent misuse of the API
        []  Add in token authentication to prevent anyone unauthorized from using the API
        []  Add in a database to keep track of all of the calculations that have taken place
        []  Add in support for floating point numbers as well.
        []  Create an associated http client that can work with the calculator API.
        []  Create a frontend that makes use of your API.
        []  Add in a middleware that adds a request ID to the http.Request object.
