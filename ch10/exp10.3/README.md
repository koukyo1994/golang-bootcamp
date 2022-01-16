# 練習問題10.3

`fetch http://gopl.io/ch1/helloworld?go-get=1`を使って、この本のサンプルコードをホストしているサービスを調べなさい。(go getからのHTTPリクエストはgo-getパラメータを含んでいるので、サーバは通常のブラウザのリクエストと区別することができます。)

## 結果

```shell
fetch "http://gopl.io/ch1/helloworld?go-get=1"
# =>
# <!DOCTYPE html>
# <html>
# <head>
# <meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
# <meta name="go-import" content="gopl.io git https://github.com/adonovan/gopl.io">
# </head>
# <body>
# </body>
# </html>
```

であるため、この本のサンプルコードをホストしているサービスは**github**である。
