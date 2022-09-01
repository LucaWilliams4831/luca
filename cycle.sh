#!/bin/sh
a=0
b=0
while [ $a -lt 10 ]
do
   echo $b
   sleep 10
   b=`expr $b + 1`
done
