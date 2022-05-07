#!/bin/bash

if [[ $1 != "-f" || $3 != "-k" || $5 != "-ip" ]]
then
    echo "Missing path to folder or key or ip as flag without the trailing slash"
    exit 1
fi

DIR_PATH=$2
KEY=$4
SERVER_IP=$6

scp -r -i $KEY ubuntu@$SERVER_IP:/var/log/nginx $DIR_PATH

for i in {1..14}
do
    gunzip -f $DIR_PATH/nginx/access.log.$i.gz
    gunzip -f $DIR_PATH/nginx/error.log.$i.gz
done