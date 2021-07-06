# Stateless Postgres Query Router

SPQR is a query router for lighweight Postgres OLTP sharding. It acts as a connection pooler, but routes transaction to different Postgres clusters. Routing is configured by set of shards and rules.


Shard is a connection string to HA cluster. e.g. 'host=rc1b-1337elrqataonjh7.mdb.yandexcloud.net port=6432 dbname=db1 user=user1'
Rule is a SELECT statement describing keyspace. e.g. 'SELECT * FROM pgbench_history JOIN pgbench_accounts USING (aid) WHERE aid < 31337'
Rule R is taking one of three possible descisions for a given query Q:
- Q within R
- Q outside R
- Inconclusive

Each decision can lead to another rule or error out.
