# Weakest-Shogi

**Weakest-Shogi** は、この世で最も弱い将棋ソフトです。
自分の手番になると、投了します。

# 動作確認済み環境

```
go version go1.16.6 darwin/amd64
```

# 将棋所への設定方法

## ビルド

このソフトと対局するには、まずビルドしてあなたのPCにあったバイナリを作りましょう。

```
./script/build.sh
```

このスクリプトを実行すると `bin/weakest-shogi` というバイナリが出力されます。
