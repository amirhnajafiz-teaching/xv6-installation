# Stallion Race

Load testing on ```Stallion``` message broker with ```Race```. By using
this service you can perform a load test on ```Stallion``` cluster in order
to measure system performance over your custom load.

## configs

Race configs are as follows, they need to be inside a ```config.yaml``` file.

```yaml
host: st://@localhost:7025 # stallion host
message: hello # publish message
debug: false # display results in stdout
consumers: 10 # number of consumers
providers: 10 # number of providers
offset: 2 # providers publish interval
topics: # system topics
  - snapp

```

## setup

In order to set up race on docker use the following image:

```shell
docker pull amirhossein21/stallion-race:v0.1.0
```

## metrics

System metrics will be available on ```localhost:8080/metrics```. Example:

```json
{
  "providers": 10,
  "consumers": 10,
  "failed_publish": 0,
  "publish": {
    "snapp": 20,
    "titan": 15,
    "workers": 15
  },
  "subscribe": {
    "snapp": 4,
    "titan": 3,
    "workers": 3
  },
  "received_message": {
    "snapp": 73,
    "titan": 40,
    "workers": 44
  },
  "topics": ["snapp", "titan", "workers"]
}
```
