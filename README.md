# 5-goprojects

https://github.com/dreamsofcode-io/goprojects

https://go.dev/doc/effective_go

To set path if go install not working:

go env to see path

export GOPATH='/Users/tylerjob/go/bin' 
export PATH=$PATH:$(go env GOPATH)/bin

1) ToDo list

    [x]  Create CSV file with ID, Description, CreatedAt, isComplete
    []  Add ToDo list item with description and created at
    []  Complete "id param" marks isComplete as true
    []  List shows all not complete tasks, flag -a shows all
    []  Delete by id

    []   encoding/csv for writing out as a csv file
    []   strconv for turning types into strings and visa versa
    []   text/tabwriter for writing out tab aligned output
    []   os for opening and reading files
    []   github.com/spf13/cobra for the command line interface
    []   github.com/mergestat/timediff for displaying relative friendly time differences (1 hour ago, 10 minutes ago, etc)