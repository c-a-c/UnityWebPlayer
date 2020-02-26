# UnityWebPlayer
みんなの作品をWebブラウザでプレイしてもらうプラットフォーム

# 技術選定

## フロントエンド

## サーバサイド
・言語：Go

・フレームワーク：Gin


クリーンアーキテクチャ
```
└── src
    ├── app
    │   ├── domain
    │   ├── infrastructure
    │   ├── interfaces
    │   │   ├── controllers
    │   │   └── database
    │   ├── usecase
    │   ├── server.go
    │   └── UnityGames
```


