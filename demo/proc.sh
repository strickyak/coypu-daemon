#!/bin/sh

cd `dirname $0`
mkdir -p coypu
cd text

for x in *
do
	# start with the contents of the text file
	cp $x ../coypu/31-$x
	# finish with 1024 of ^Z, which is at least enough.
	python3 -c 'print( 1024 * "\032")' >>../coypu/31-$x
done

python3 ../create-directory.py * > ../coypu/30-
