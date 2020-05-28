#!/bin/bash
kubectl.exe get pods | grep frontend | awk '{print $1}' | tr -d '\n' | xargs -0 kubectl.exe logs
rrrrrr