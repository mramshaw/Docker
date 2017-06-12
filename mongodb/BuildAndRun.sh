#!/bin/bash

if [ $EUID -ne 0 ]
then
    echo "This script should be run as the root user or else using 'sudo'"
    exit 1
fi

mongodir=/var/mongodb

populatemongo=false

echo "Building docker images ..."
docker build -t mramshaw4docs/pumongodb ./Database
docker build -t mramshaw4docs/puorder   ./Order
docker build -t mramshaw4docs/puweb     ./Clients
echo " "

if [ -d $mongodir ]
then
	echo 'MongoDB data directory exists';
	populatemongo=false
else
	echo 'Creating MongoDB data directory ...';
	mkdir -p $mongodir
	populatemongo=true
fi
echo " "

echo "Running docker images ..."
docker run -it -d --name db    -p 27017:27017 -p 28017:28017 -v $mongodir:/data/db mramshaw4docs/pumongodb
docker run -it -d --name order -p  8080:8080                 --link db:mongo       mramshaw4docs/puorder
docker run -it -d --name web   -p    80:8080                                       mramshaw4docs/puweb
echo " "

if [ $populatemongo = true ]
then
	echo "Populating MongoDB ..."
	docker exec db mongo ordering /tmp/MongoRecords.js
fi
echo " "

echo "Access the Parts Unlimited MRP system at 'http://localhost/mrp'"
