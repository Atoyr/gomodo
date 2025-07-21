# 🍅 GoModo - Pomodoro CLI Timer for Developers

**GoModo** は、開発者のためのシンプルな Go 製ポモドーロタイマー CLI ツールです。  
Alacritty や zellij のような軽量ターミナル環境に最適で、集中と休憩のリズムを自然に取り入れられます。

---

## 🚀 特徴

- ✅ シンプルなコマンド：`gomodo start`
- 🕒 25分作業 → 5分休憩 のポモドーロサイクル
- 🖥️ macOS 通知センターで開始・終了を通知
- 🛑 `Ctrl+C` でいつでも中断可能
- 🧠 習慣化・集中力向上に最適

---

## 🔧 インストール

```bash
git clone https://github.com/yourname/gomodo.git
cd gomodo
go build -o gomodo
```

または Go モジュールとして使用：

```bash
go install github.com/yourname/gomodo@latest
```


## 📝 使い方
```bash
gomodo start
```

コピーする
編集する
作業開始の通知が表示され、25分間のタイマーが始まります。

作業終了後、自動的に 5分間の休憩タイマーが開始されます。

このサイクルは Ctrl+C で停止するまで継続します。



