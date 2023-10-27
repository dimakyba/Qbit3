#!/bin/bash

for N in {1..10} 
do 
    mkdir Ex$N/ 
    mv $N.go Ex$N/ 
    cd Ex$N/ 
    mv $N.go main.go
    cd .. 
done

echo "Done!"
