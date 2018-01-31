# gogeocheck
Simple checker geo by ip on Golang

``` 
$ go run geocheck.go 
8.8.8.8,United States
109.22.12.33,France
8.8.4.4,United States
213.132.122.4,Sweden
```

With City use -c flag
```
$ go run geocheck.go -c
8.8.8.8,United States,
109.22.12.33,France,Chauny
8.8.4.4,United States,
213.132.122.4,Sweden,
```

Help
```
$ go run geocheck.go -h
-db	Path to database
-f	Path to file
-c	Include information about City
-a	ipaddress IP to resolve
-h	help
You can don't use flags, if import db to directory with program and use name GeoLite2-City.mmdb, use  list of ip addresses with name ip_part

```
