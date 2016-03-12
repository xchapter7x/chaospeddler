# local dev

##
- download wercker cli:
    http://wercker.com/cli/install/

## using wercker to run tests
**(`init` will pull fresh docker images)**
```bash

$ ./testrunner
or 
$ ./testrunner init
```

## using wercker to run the service locally

**create a env file**

```bash
$ touch pws-creds.env 
$ vim pws-creds.env
```

**sample file contents**
```bash
X_CF_USER=mycfuser@gmail.com
X_CF_PASS=somepass
X_CF_LOGIN_URL=https://login.run.pivotal.io
X_CF_CC_URL=https://api.run.pivotal.io
```

**run service locally**
**(`init` will pull fresh docker images)**
```bash
$ ./start_service 
or 
$ ./start_service init
```

**hitting service**
```bash
curl $DOCKER_HOST/v1/api:8080
```
