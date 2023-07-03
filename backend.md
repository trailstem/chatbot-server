# chatbot-server

### Dockerfile.dep 検証用コマンド

```
docker build --no-cache -t dep_test -f Dockerfile.dep .
docker run --name dep dep_test
docker exec -it dep sh
```

### heroku のデプロイ先作成

```
heroku login
heroku create -a hogehuga
// local repositoryをHerokuに接続
heroku git:remote -a hogehuga
// デプロイ
git push heroku main
//local mainBranch以外からのデプロイ
git push heroku testbranch:main
```
