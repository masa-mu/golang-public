
## midleware
- mysql
- docker
- 
- 
- 

## go plugin
- github.com/jinzhu/gorm
- 
- 
- 
- 

## memo
- AWSにするならMysqlのアカウント情報はSecretManagerに保持する
- dep使うなら.gitignoreにvender/を含めないとPUSHされてしまうので重い

## 実行コマンド
```aidl
$ docker-compose up --build
```

## サンプルリクエスト
```aidl
$ curl http://localhost:8000/ddd/v1/message?id=12345
```
