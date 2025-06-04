# MontageClipList2ASS

[WaveLab Pro](https://www.steinberg.net/ja/wavelab/)のモンタージュのXMLクリップリストファイルをASS(Advanced SubStation Alpha)に変換するツールです。

## 使い方

    $ mcl2ass -f [XMLファイル] -c [設定ファイル] -o [出力ASSファイル] -x [出力YouTube 詳細]
    
ヘルプ

    $ mcl2ass -h

バージョン

    $ mcl2ass -v

使用例

    $ mcl2ass -f example/ClipList.xml -c example/config.yaml -o output.ass

## バイナリ

[releases](https://github.com/mamemomonga/MontageClipList2ASS/releases/)

ファイル名 | OS
---|---
mcl2ass-darwin-arm64 | macOS Apple Silicon
mcl2ass-darwin-amd64 | macOS Intel
mcl2ass-linux-amd64 | Linux x86_64
mcl2ass-linux-arm | Linux ARMv7
mcl2ass-linux-arm64 | Linux ARMv8
mcl2ass-windows-amd64.exe | Windows x86_64

## ビルド

ビルド

    $ make

マルチアーキテクチャビルド

    $ make multiarch

Dockerでビルド

    $ make docker

即時実行とプレビュー

    $ ./run.sh -h