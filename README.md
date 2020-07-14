# loggen

### Build
```
docker build -t loggen:v1 .
docker tag loggen:v1 us.icr.io/order-mgmt-np/oms/loggen:v1
docker push us.icr.io/order-mgmt-np/oms/loggen:v1
```

### Usage
```
./loggen <log-file-path> <no-of-threads>

eg., ./loggen logs.txt 10
```
