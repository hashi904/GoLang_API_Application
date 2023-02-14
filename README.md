# golang API Application

## getting started
```bash
# build and start containers
$ docker-compose build
$ docker-compose up

# into golang app container
$ docker-compose exec app sh

# execute starting command
$ cd /go/src 
$ gin -i --all run main.go
```

## What I want to do
- [ ] saveしたらコンパイルできるようにする。
- [ ] container起動時にサーバーも起動する。
- [ ] プログラムを止めてデバッグできるようにする。
- [ ] 大規模アプリケーションに対応できるように設計を考える。
- [ ] 自動テストの導入。
- [ ] APIをpostgresqlに繋いでデータを管理する。

