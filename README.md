# newrelic-remoe

Nature Remo Eの電力情報をNew Relicに送信する小さいものです。

## 使い方
### Nature Remo Eを買う
[買ってください](https://amzn.to/3qWZ0yh)。

### Nature API Tokenを取得する
Remo Eを良い感じにセットアップしていただいて、 https://home.nature.global/ にログインしたらなんとなくtokenが発行できます。

発行したtokenは、環境変数 `NATURE_API_TOKEN` として定義してください。

### New Relic API Keyを取得する
[Insights insert key (別名 "Insert key")](https://docs.newrelic.com/docs/apis/get-started/intro-apis/types-new-relic-api-keys)を発行してください。

発行したキーは、環境変数 `NEW_RELIC_INSERT_API_KEY` として定義してください。

### 動かす

```shell script
make run
```

でとりあえず動きます。試すなら手動で数回叩いてもいいですし、cronやwatchコマンドなどをつかって60秒ごとに実行させてもよいです。

### 結果を見る

New Relic Oneを開いて、右上のQuery Your Dataからメトリックの情報をクエリできます。Data Explorerから行っても良いですし、たとえばNRQLはこんな感じに発行してみましょう。

```sql
FROM Metric SELECT average(remoe.MeasuredInstantaneous) FACET modelId TIMESERIES 
```
