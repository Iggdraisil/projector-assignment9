# SQL Benchmarks
 - To launch run `docker run -d -e MYSQL_ROOT_PASSWORD=1234 -p 3306:3306 --name my-mysql mysql`
 - To generate users (csv) run datagen
 - To run inserter rest api run inserter

## Index testing:

| query                                                                      | no index | BTREE     | HASH      |
|----------------------------------------------------------------------------|----------|-----------|-----------|
| `SELECT * FROM user WHERE birthdate BETWEEN '2000-12-01' AND '2001-01-01'` | **42s**  | **4.5s**  | **5.1s**  |
| `SELECT * FROM user WHERE birthdate BETWEEN '2000-01-01' AND '2001-01-01'` | **53s**  | **42s**   | **1m 9s** |
| `SELECT * FROM user WHERE birthdate = '2000-01-01'`                        | **25s**  | **325ms** | **495ms** |

## Insertion testing concurrency 400
`siege -b -t1m -c400 'http://127.0.0.1:3000/person POST'`

| innodb_flush_log_at_trx_commit | rps     | 
|--------------------------------|---------|
| 0                              | **284** |
| 1                              | **151** |
| 2                              | **611** | 

