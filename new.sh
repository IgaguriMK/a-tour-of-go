#! /bin/bash

if [ $# -lt 1 ]; then
	echo "Need arg."
	exit 1
fi

echo $1 >> .gitignore
gvim $1.go
