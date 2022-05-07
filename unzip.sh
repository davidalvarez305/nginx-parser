#!/bin/bash

for i in {2..14}
do
    gunzip ./nginx/access.log.$i.gz
    gunzip ./nginx/error.log.$i.gz
done