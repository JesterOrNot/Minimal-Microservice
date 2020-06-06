#!/bin/bash

for file in kubernetes/*; do
    kubectl.exe "$1" -f "$file"
done
