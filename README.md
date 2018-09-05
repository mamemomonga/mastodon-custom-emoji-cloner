# mastodon-custom-emoji-cloner

* マストドンのインスタンスが発見した絵文字を片っ端からコピーするツールです。
* 運が悪いとあなたの大切なインスタンスを壊す可能性がありますので、それを了承の上お使いください。
* emoji-cloner-linux-arm のみ動作を確認しています。
* S3ストレージには対応していません。

# 実行方法

* config-example.yaml を参考に config.yaml を作成してください。
* PostgreSQLにTCPで接続できるようにする必要があります。

OS、ARCHにあったバイナリをダウンロードし、以下の実行します。

	$ ./emoji-cloner-linux-arm config.yaml

