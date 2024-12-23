#!/bin/bash

if [ -d "target" ]; then
    rm -rf "target"
fi

if [ -d "vt" ]; then
    rm -rf "vt"
fi

if [ -f "Cargo.lock" ]; then 
    rm -rf "Cargo.lock"
fi

sudo docker stop $(sudo docker ps -q)
sudo docker rm $(sudo docker ps -aq)
cargo build --release
mv target/release/vt vt
