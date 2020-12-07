#!/usr/bin/env bash

if [ $# -lt 2 ]; then
    echo "Pass three input arguments: <AoC session cookie> <Number of day (1-25)>"
    exit 1
fi

CONTENT=$(curl "https://adventofcode.com/2020/day/$2" -H "Cookie: session=$1")
CONTENT_LEN=$(echo -n "$CONTENT" | wc -l)
if [ "$CONTENT_LEN" -lt 2 ]; then
    echo "Day is not available now"
    exit 1
fi

DAY_NUMBER=$(echo -n "$CONTENT" | sed -n "s/^.*--- \(Day [0-9]*\): \([a-zA-Z0-9 -]*\) ---.*/\1/p")
DAY_DESC=$(echo -n "$CONTENT" | sed -n "s/^.*--- \(Day [0-9]*\): \([a-zA-Z0-9 -]*\) ---.*/\2/p")

if [ 10 -gt "$2" ]; then
    DAY_MODULE="Day-0$2"
else
    DAY_MODULE="Day-$2"
fi

DAY_MODULE_FROM_DESC=$(echo -n "$DAY_DESC" | sed -e "s/ /-/g")
DAY_MODULE="$DAY_MODULE-$DAY_MODULE_FROM_DESC"
DAY_DESC="$DAY_NUMBER: $DAY_DESC"

printf "\nNew day: %s\n\n" "$DAY_DESC"

if [ ! -d "../$DAY_MODULE" ]; then
  rsync -a ../Day-XX-template/* ../"$DAY_MODULE" --exclude go.mod --exclude go.sum
  echo "Creating new day module named $DAY_MODULE"
else
  echo "Module $DAY_MODULE already exists"
  exit 1
fi

find ../"$DAY_MODULE" \( -type d -name .git -prune \) -o -type f -print0 | xargs -0 sed -i '' -e "s/_DAY_NAME_/$DAY_MODULE/"
find ../"$DAY_MODULE" \( -type d -name .git -prune \) -o -type f -print0 | xargs -0 sed -i '' -e "s/_DAY_DESC_/$DAY_DESC/"

cd ../"$DAY_MODULE" || exit 1
go mod init "$DAY_MODULE"
go mod download
go mod verify
go mod tidy
go test ./...

aocdl -session-cookie "$1" -output data/input

git add .

# Remember to sync new go module ([PPM] go.mod -> Sync Go Module)