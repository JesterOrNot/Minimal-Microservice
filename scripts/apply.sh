#!/bin/bash

for file in kubernetes/*; do
    kubectl apply -f $file
done
