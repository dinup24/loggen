# loggen

### Build
```
docker build -t loggen:v1 .
docker tag loggen:v1 us.icr.io/order-mgmt-np/oms/loggen:v1
docker push us.icr.io/order-mgmt-np/oms/loggen:v1
```

### Usage
```
./loggen <wait-time-in-minutes>

eg., ./loggen 10
```
