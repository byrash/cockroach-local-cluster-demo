#! /bin/bash

cockroach cert create-ca \
    --certs-dir=./certs \
    --ca-key=./certs/ca.key \
    --overwrite \
    --allow-ca-key-reuse

mkdir -p ./certs/node1
mkdir -p ./certs/node2
mkdir -p ./certs/node3
cp ./certs/ca.crt ./certs/node1/ca.crt
cp ./certs/ca.crt ./certs/node2/ca.crt
cp ./certs/ca.crt ./certs/node3/ca.crt

cockroach cert create-node \
    roach1 \
    localhost \
    127.0.0.1 \
    --certs-dir=./certs/node1 \
    --ca-key=./certs/ca.key \
    --overwrite

cockroach cert create-node \
    roach2 \
    localhost \
    127.0.0.1 \
    --certs-dir=./certs/node2 \
    --ca-key=./certs/ca.key \
    --overwrite

cockroach cert create-node \
    roach3 \
    localhost \
    127.0.0.1 \
    --certs-dir=./certs/node3 \
    --ca-key=./certs/ca.key \
    --overwrite

cockroach cert create-client \
    root \
    --certs-dir=./certs \
    --ca-key=./certs/ca.key \
    --overwrite

cp ./certs/client.root.crt ./certs/node1/client.root.crt
cp ./certs/client.root.key ./certs/node1/client.root.key
