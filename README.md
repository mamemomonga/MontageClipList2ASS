# MontageClipList2ASS

[WaveLab Pro](https://www.steinberg.net/ja/wavelab/)のモンタージュのXMLクリップリストファイルをASS(Advanced SubStation Alpha)に変換するツールです。

## 使い方

    $ ./montage_cliplist2ass -f [XMLファイル] -t [テンプレートファイル] -o [出力ASSファイル]
    
ヘルプ

    $ ./montage_cliplist2ass -h

バージョン

    $ ./montage_cliplist2ass -v

使用例

    $ ./montage_cliplist2ass -f example/ClipList.xml -t example/ass.tpl -o output.ass

## ビルド

ビルド

    $ make

マルチアーキテクチャビルド

    $ make multiarch

Dockerでビルド

    $ make docker

即時実行

    $ ./run.sh -h