# gogeocheck
Simple checker geo/asn by ip

``` 
$ go run geocheck.go 
8.8.8.8,United States,,15169,Google LLC
109.22.12.33,France,Chauny,15557,SFR SA
8.8.4.4,United States,,15169,Google LLC
213.132.122.4,Sweden,,12552,IP-Only Networks AB
```

Short check one ip
```
$ go run geocheck.go -a 8.8.8.8
8.8.8.8,United States,,15169,Google LLC
```

