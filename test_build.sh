#!/bin/bash

set -e -u -x
mkdir cliargs
mv cliargs.go cliargs/

mv main.go.tpl main.go
go build -o main

#undo
mv cliargs/cliargs.go ./cliargs.go
rm -r cliargs/
mv main.go main.go.tpl