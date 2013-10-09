#! /bin/sh

# create classes folder
mkdir bin

# compile
javac -cp "lib/*:." -d bin src/rpc/*

# execute
cd bin
java -cp "../lib/*:." rpc.Main

# remove classes
cd -
rm bin -rf
