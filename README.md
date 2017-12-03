# test hangs on GAE/Go

required
 - [gcloud sdk](https://cloud.google.com/sdk/downloads)
 - app engine sdk / go
 - GOMAXPROCS > 1

## Install

Install gcloud sdk
```
$ curl https://sdk.cloud.google.com | bash
$ exec -l $SHELL
$ gcloud init
```

Install app engine sdk
```
$ gcloud components install app-engine-go
$ chmod +x /path/to/google-cloud-sdk/platform/google_appengine/go*
```

Install go-test-hangs-sample
```
$ mkdir -p $GOPATH/src/github.com/utahta
$ cd $GOPATH/src/github.com/utahta
$ git clone https://github.com/utahta/go-test-hangs-sample.git
```

Install libraries
```
$ cd go-test-hangs-sample
$ dep ensure
```

## Hangs

```
$ goapp test -v ./bar ./foo
```

