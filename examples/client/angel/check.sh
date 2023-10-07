#!/bin/bash

# Example check script for angel.
# Will check if a file called "flag" is found
# in the same directory of this script
example="$(dirname $0)/flag"

if [[ -f $example ]]
then
	exit 0
else
	exit 1
fi