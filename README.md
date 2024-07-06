# CurlARC-Service
## CurlARCのバックエンド

### 使い方
- サーバーの起動
```sh
$ docker compose up
```

- dbのテーブルデータの確認
```sh
$ docker exec -it $(container_id) bash
$ psql -U app -d app
$ \dt
$ SELECT * FROM ${table_name};
```