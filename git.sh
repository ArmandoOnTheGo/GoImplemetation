#!/bin/bash

echo "Enter Commit Message"
read mes

git add .
git commit -m "$mes"
git push
