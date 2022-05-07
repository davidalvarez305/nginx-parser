#!/bin/bash

if [[ $1 != "-f" ]]
then
    echo "Missing path to folder as flag without the trailing slash"
    exit 1
fi

DIR_PATH=$2

for i in {2..14}
do
    gunzip $DIR_PATH/access.log.$i.gz
    gunzip $DIR_PATH/error.log.$i.gz
done