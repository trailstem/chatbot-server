# speakbot-backend

# heroku のデプロイ先作成

- デプロイ先作成

```
heroku create -a hogehuga
```

-

```
heroku git:remote -a hogehuga
```

```
git push heroku main
```

ローカルのメインブランチ以外からのデプロイ
git push heroku testbranch:main

---

# heroku.yml を使用して Docker イメージ

### スタック設定

heroku stack:set container

### push

git push heroku master

### 特定アプリにアドオンを紐づける

heroku addons:create <addon_name> -a <app_name>
