# Nakama

Source code of the next social network for anime fans. Still on development.

## Getting the code

Besides having [Go](https://golang.org/) installed, the server needs two external services. A SQL database; I'm using [CockroachDB](https://www.cockroachlabs.com/), but [Postgres](https://www.postgresql.org/) should work too. Also, an SMTP server; I recommend [mailtrap.io](https://mailtrap.io/) for development.

Copy the `.env.example` file into `.env`. You need to at least set `SMTP_USERNAME` and `SMTP_PASSWORD` before trying it out.

## Building

Before running the server, you need a CockroachDB instance running.

```bash
cockroach start --insecure --host 127.0.0.1
```

Then, you need to create the database and tables.

```bash
cat schema.sql | cockroach sql --insecure
```

Now, you can build and run the server.

```bash
go build
./nakama
```

## Dependencies

These are the Go libraries used in the source code. Thank you very much.
 - [github.com/disintegration/imaging](https://github.com/disintegration/imaging)
 - [github.com/hako/branca](https://github.com/hako/branca)
 - [github.com/joho/godotenv](https://github.com/joho/godotenv)
 - [github.com/lib/pq](https://github.com/lib/pq)
 - [github.com/matoous/go-nanoid](https://github.com/matoous/go-nanoid)
 - [github.com/matryer/way](https://github.com/matryer/way)
