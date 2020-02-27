# UnityWebPlayer
みんなの作品をWebブラウザでプレイしてもらうプラットフォーム

# 技術選定

## フロントエンド

## Server-Side tech stack
Language : Go

Framework : Gin

Architecture : Clean Archi

Struct of directory.
```
└── src
    ├── app
    │   ├── domain
    │   ├── infrastructure
    │   ├── interfaces
    │   │   ├── controllers
    │   │   └── database
    │   ├── usecase
    │   ├── server.go : Main
    │   ├── HTML : Default HTML files. Will be changed by js.
    │   ├── Wating : Will be checked applications before deploy.
    │   ├── Unity : Deployed applications.
    │   └── Deleted : Pool of deleted applications from Unity.
```


