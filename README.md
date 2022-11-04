To genearte certificates ofr teh cluster

[code]
cp gen-certs.sh ./certs/gen-certs.sh
docker run --entrypoint /cockroach/certs/gen-certs.sh --rm -v ${PWD}/certs:/cockroach/certs:rw cockroachdb/cockroach:v22.1.10

init DB using
docker exec -it roach1 ./cockroach init --certs-dir=certs

https://www.cockroachlabs.com/docs/stable/start-a-local-cluster-in-docker-mac.html
