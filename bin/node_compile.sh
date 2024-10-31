#!/bin/bash

function exit_if() {
    extcode=$1
    msg=$2
    if [ $extcode -ne 0 ]
    then
        if [ "msg$msg" != "msg" ]; then
            echo $msg >&2
        fi
        exit $extcode
    fi
}

echo Compiling node interfaces...
mkdir -p node

#protoc -I ./ --java_out=./java --grpc-java_out=./java dapplink/*.proto
protoc --js_out=import_style=commonjs,binary:./node dapplink/*.proto

exit_if $?
echo Done
