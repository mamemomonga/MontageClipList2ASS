# MontageClipList2ASS

[WaveLab Pro](https://www.steinberg.net/ja/wavelab/)のモンタージュのXMLクリップリストファイルをASS(Advanced SubStation Alpha)に変換するツールです。

## 使い方

    $ ./montage_cliplist2ass -f [XMLファイル] -c [設定ファイル] -o [出力ASSファイル]
    
ヘルプ

    $ ./montage_cliplist2ass -h

バージョン

    $ ./montage_cliplist2ass -v

使用例

    $ ./montage_cliplist2ass -f example/ClipList.xml -c example/config.yaml -o output.ass

## バイナリ

[releases](https://github.com/mamemomonga/MontageClipList2ASS/releases/)

* montage_cliplist2ass-darwin-arm64 macOS Apple Silicon
* montage_cliplist2ass-darwin-amd64 macOS Intel
* montage_cliplist2ass-linux-amd64 Linux x86_64
* montage_cliplist2ass-linux-arm Linux ARMv7
* montage_cliplist2ass-linux-arm64 Linux ARMv8
* montage_cliplist2ass-windows-amd64.exe Windows x86_64

## ビルド

ビルド

    $ make

マルチアーキテクチャビルド

    $ make multiarch

Dockerでビルド

    $ make docker

即時実行とプレビュー

    $ ./run.sh -h