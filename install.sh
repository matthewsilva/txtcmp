#!/bin/bash
echo "Welcome to txtcmp installer."
if [ "$EUID" -ne 0 ]
then
    echo "This script needs to be run as root (e.g. use sudo)"
    echo "exiting..."
    exit
fi
echo "Please select installation method:"
echo "1. Install from precompiled binary"
echo "2. Build from source"
read option
echo "Option ="
echo $option
if [ $option -eq 1 ]
then
    echo "Downloading binary..."
    echo
    wget https://github.com/matthewsilva/txtcmp/raw/master/bins/txtcmp
    echo
    echo "Finished donwloading"
elif [ $option -eq 2 ]
then
    echo "Downloading source..."
    echo
    wget https://github.com/matthewsilva/txtcmp/raw/master/txtcmp.go
    echo
    echo "Finished donwloading"
    echo
    echo "Building..."
    go build txtcmp.go
    echo "Finished building"
else
    echo "Invalid option, exiting"
    exit
fi



echo
echo "Installing txtcmp..."
chmod 777 ./txtcmp
cp ./txtcmp /usr/bin/
echo "txtcmp installed"
echo
echo "Cleaning up files..."
rm -f ./txtcmp*
echo "Files cleaned up"
echo
echo "Installation successful!"
echo "Usage: txtcmp DOC_1_FILEPATH DOC_2_FILEPATH"
echo
