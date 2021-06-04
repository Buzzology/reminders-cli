# Golang Tutorial by Steve Hook
https://www.youtube.com/watch?v=-9CbX2MncZg

A simple cli client that can perform CRUD on reminders. Also has a single endpoint for health checks.

This tutorial was not using go modules. To disable them ensure that the following environment variable is set: `$env:GO111MODULE = "auto"`.  

## Notifier
Simple nodejs application that displays desktop notifications.

![image](https://user-images.githubusercontent.com/27006526/120097269-efe31000-c172-11eb-8584-1ca27512163f.png)

### Run the node notifier
node notifier.js  

## Using make
Note that on windows this needs to be installed separately. You may also need to add an "exe" extension to the dll in order to circumvent the "Choose a program to run this file" prompt.  

## Examples

### Os-Args
Handle os args with manual parsing.

### Flagsets
Handle os args using flagsets. Note that running `./flagsets greet --help` provides built in dialog.

### Flag-Value
Using flagsets to grab custom types from command line args. Usage example:  
```
> ./flag-value --id=1 --id=5 --id=2523 -name="Chris"
1,5,2523
Hello, my name is Chris and I was born 2021
```


## Client commands
After running `make client` you can use the following:
Create: `./bin/client.exe create t="Some title" m="My Message" d=1`  
Edit: `./bin/client.exe create id=1 t="Some title" m="My Message" d=1`  
Delete: `./bin/client.exe delete id=1`  
Health: `./bin/client.exe health`  

Or to run without creating a bin use:
Edit: `go run cmd/client/main.go edit -id="1" t="test Title" m="test message" d=3`



