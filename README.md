# mastodon-custom-emoji-cloner

* マストドンのインスタンスが発見した絵文字を片っ端からコピーするツールです。
* Golangの勉強ついでに作成したプログラムです。
* データベースへINSERT、ファイルのコピーを行います。試される場合は必ずファイルとデータベースのバックアップをとってください。運が悪いとあなたの大切なインスタンスを壊す可能性があります。
* ピッカーは自動更新されません。自動取得後はウェブページをリロードする必要があります。
* S3ストレージには対応していません。

# 実行方法

* config-example.yaml を参考に config.yaml を作成してください。
* PostgreSQLにTCPで接続できるようにする必要があります。

OS、ARCHにあったバイナリを[ここから](https://github.com/mamemomonga/mastodon-custom-emoji-cloner/releases/)ダウンロードし、以下の実行します。

	$ ./emoji-cloner-linux-arm config.yaml

## バックグラウンドでの実行例

	$ emoji-cloner-linux-arm config.yaml > emoji-cloner.log 2>&1 &

