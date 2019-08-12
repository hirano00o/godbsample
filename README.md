# godbsample
database unit test sample  

use sqlmock

# usage
1. run `make dbtest`

2. set environment variable

```
export DB_USER=mysql
export DB_PASSWORD=mysql
export DB_HOST=localhost
export DB_PORT=13306
export DB_NAME=test
```

3. `make mysql.start` starts mysql on docker
4. `make start` starts program
5. `mysql.stop` stop mysql on docker

