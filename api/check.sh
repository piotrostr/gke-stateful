#!/bin/bash

binary_name=stateful-api

look_for_existing () {
        gcloud artifacts repositories list \
                | grep $binary_name \
                | awk '{print $1}'
}

printf "Looking for existing '%s' artifact repository...\n" $binary_name
if [ -z "$(look_for_existing)" ]; then
        printf "Doesn't exist, creating...\n"
        gcloud artifacts repositories create $binary_name \
                --repository-format=docker \
                --location=us-central1
else
  printf "Exists!\n"
fi
