# MontageClipList2ASS

[WaveLab Pro](https://www.steinberg.net/ja/wavelab/)のモンタージュのXMLクリップリストファイルをASS(Advanced SubStation Alpha)に変換するツールです。

output.txtの内容は、YouTubeの「説明」に貼ることで、時間リンクの見出しを作成することができます。

## 使い方

    $ mcl2ass -f [XMLファイル] -c [設定ファイル] -o [出力ASSファイル] -x [YouTube 説明テキスト出力]
    
ヘルプ

    $ mcl2ass -h

バージョン

    $ mcl2ass -v

使用例

    $ mcl2ass -f example/ClipList.xml -c example/config.yaml -o output.ass -x output.txt

使用例(ブルーバックでオーディオ入りの動画作成)

    $ ffmpeg -f lavfi -i color=c=blue:s=1920x1080:r=24 \
            -i audio.wav \
            -vf "ass=output.ass" \
            -c:v libx264 -pix_fmt yuv420p -c:a aac -shortest \
            output.mp4


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