# Golang × OpenAi API × OCR
PDFや画像データをスキャンし特定のデータを取得するためのAPI

---

## 1. 動かすために必要なこと

### 1.1 .envファイルの作成
```shell
cp .env.sample .env
```

・Lineでアカウントを作成して、チャンネルシークレットとアクセストークをセットする<br>
・OpenAIのAPIキーをセットする
```text
LINE_CHANNEL_SECRET=
LINE_CHANNEL_ACCESS_TOKEN=

OPENAI_API_KEY=
```

---

## 2. 起動・停止方法
### 2.1 起動方法
imageの作成
```shell
docker compose build
```

imageからコンテナの起動
```shell
docker compose up -d
```

---

### 2.2 停止方法
```shell
docker compose down
```

---



