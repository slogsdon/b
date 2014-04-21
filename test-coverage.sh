#!/bin/bash

echo "mode: set" > acc.out

returnval=`gocov test github.com/slogsdon/b | gocov report`
echo ${returnval}
if [[ ${returnval} != *FAIL* ]]
then
  if [ -f coverage.json ]
  then
      cat coverage.json | grep -v "mode: set" >> acc.out 
  fi
else
  exit 1
fi

for Dir in $(find ./* -maxdepth 10 -type d ); 
do
  if ls $Dir/*.go &> /dev/null;
  then
    returnval=`gocov test github.com/slogsdon/b/$Dir | gocov report`
    echo ${returnval}
    if [[ ${returnval} != *FAIL* ]]
    then
        if [ -f coverage.json ]
        then
            cat coverage.json | grep -v "mode: set" >> acc.out 
        fi
      else
        exit 1
      fi  
    fi
done
if [ -n "$COVERALLS_TOKEN" ]
then
  goveralls -service drone.io -repotoken $COVERALLS_TOKEN -coverprofile=acc.out
fi  

rm -rf ./coverage.json
rm -rf ./acc.out