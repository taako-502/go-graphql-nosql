# go-graphql-nosql

`Go` + `GraphQL` + `DynamoDB` サンプルプログラム

## Description

このリポジトリは、`Go` 言語を使用して `GraphQL` の API を構築し、AWS の `DynamoDB` に対して操作を行うサンプルプログラムを公開しています。`GraphQL` を通じて、`DynamoDB` 上のデータを簡単にクエリやミューテーションができることを示すためのものです。

## Usage

`.env_example` を参考に.`env` ファイルを作成してください。

作成後、以下コマンドを実行します。

```zsh
$ docker-compose build
$ docker-compose up -d
```

初回実行時は以下のコマンドでデータベースのマイグレーションを実行してください。

```zsh
$ make migrate
```

## Deploy

```zsh
$ make build && serverless deploy
```

## Build

```zsh
$ go build -o bin/handler handler/server/server.go
```

## Frontend

https://github.com/taako-502/staggered-scheduler
