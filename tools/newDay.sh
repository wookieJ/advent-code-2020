#!/usr/bin/env bash

if [ $# -lt 2 ]; then
    echo "Pass two input arguments: <Day module name> and <Day description>"
    exit 1
fi

if [ ! -d "../$1" ]; then
  cp -rn ../Day-XX-template ../"$1"
  echo "Creating new day module named $1"
else
  echo "Module $1 already exists"
  exit 1
fi

find ../Day-01-test \( -type d -name .git -prune \) -o -type f -print0 | xargs -0 sed -i '' -e "s/_DAY_NAME_/$1/"
find ../Day-01-test \( -type d -name .git -prune \) -o -type f -print0 | xargs -0 sed -i '' -e "s/_DAY_DESC_/$2/"