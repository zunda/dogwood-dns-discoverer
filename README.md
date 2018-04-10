# dogwood-dns-discoverer
An app that exposes private DNS

## Run locally
```
$ go run main.go
```

## Build dependency
```
$ go get github.com/julienschmidt/httprouter
$ godep save ./...
```

## Adding SSH key to connect to your EC2 instances
```
$ heroku config:set "SSH_RSA_PRIV_KEY=`cat <rsa-prvate-key-file>`"
$ heroku config:set "KNOWN_HOSTS=`grep ec2 ~/.ssh/known_hosts`"
```
