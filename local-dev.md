# local dev

##
- download wercker cli:
    http://wercker.com/cli/install/

## using wercker to run tests
```bash

$ ./testrunner
```

## using wercker to run the service locally

**create a env file**

```bash
$ touch local-dev.md
$ vim local-dev.md
```

**sample file contents**
```bash
X_CF_USER=mycfuser@gmail.com
X_CF_PASS=somepass
X_CF_LOGIN_URL=https://login.run.pivotal.io
X_CF_CC_URL=https://api.run.pivotal.io
```

**run service locally**
```bash
$ wercker --environment pws-creds.env dev --publish 8080
```

**hitting service**
```bash
curl $DOCKER_HOST/v1/api:8080
```
