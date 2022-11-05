## To genearte certificates for the cluster

```
cp gen-certs.sh ./certs/gen-certs.sh
docker run --entrypoint /cockroach/certs/gen-certs.sh --rm -v ${PWD}/certs:/cockroach/certs:rw cockroachdb/cockroach:v22.1.10
```

## Start Cluster

```
docker-compose up --build
```

## init DB using

```
docker exec -it roach1 ./cockroach init --certs-dir=certs
```

[Staring Local Cluster](https://www.cockroachlabs.com/docs/stable/start-a-local-cluster-in-docker-mac.html)

## Create DB users

```
docker exec -it roach1 ./cockroach sql --certs-dir=certs
```

```
alter user root password 'welcome123';
```

Run schema.sql content in the same shell as above

## Login to console

[console URL](https://localhost:8080/#/overview/list)

## Uname/PWD

root/welcome123

** Connect to your favorite SQL ediot using ca cert and user name and password above **

## Tail app logs

`docker-compose -f docker-compose-app.yaml logs --tail=0 --follow`
