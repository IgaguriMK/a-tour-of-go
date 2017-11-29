#! /bin/bash

if [ $# -lt 1 ]; then
	echo "Need arg."
	exit 1
fi

if [ $# -ge 2 ]; then
	cp $2 $1.go
fi

echo $1 >> .gitignore
gvim $1.go
