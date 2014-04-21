#!/bin/bash

echo "mode: set" > acc.out

returnval=`go test -coverprofile=profile.out github.com/slogsdon/b`
echo ${returnval}
if [[ ${returnval} != *FAIL* ]]
then
  if [ -f profile.out ]
  then
      cat profile.out | grep -v "mode: set" >> acc.out 
  fi
else
  exit 1
fi

for Dir in $(find ./* -maxdepth 10 -type d ); 
do
  if ls $Dir/*.go &> /dev/null;
  then
    returnval=`go test -coverprofile=profile.out github.com/slogsdon/b/$Dir`
    echo ${returnval}
    if [[ ${returnval} != *FAIL* ]]
    then
        if [ -f profile.out ]
        then
            cat profile.out | grep -v "mode: set" >> acc.out 
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

rm -rf ./profile.out
rm -rf ./acc.out