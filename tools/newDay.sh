#!/usr/bin/env bash

if [ $# -lt 2 ]; then
    echo "Pass two input arguments: <Day module name> and <Day description>"
    exit 1
fi

if [ ! -d "../$1" ]; then
  rsync -a ../Day-XX-template/* ../"$1" --exclude go.mod --exclude go.sum
  echo "Creating new day module named $1"
else
  echo "Module $1 already exists"
  exit 1
fi

find ../"$1" \( -type d -name .git -prune \) -o -type f -print0 | xargs -0 sed -i '' -e "s/_DAY_NAME_/$1/"
find ../"$1" \( -type d -name .git -prune \) -o -type f -print0 | xargs -0 sed -i '' -e "s/_DAY_DESC_/$2/"

cd ../"$1" || exit 1
go mod init "$1"
go mod download
go mod verify
go mod tidy
go test ./...

# Remember to sync new go module ([PPM] go.mod -> Sync Go Module)

# todo - set title and description from web, add README for module