# cv-admin

GitHub Pages にて公開している [Curriculum-Vitae](https://sky0621.github.io/cv/) のデータソースを管理する。

データソース管理には [SQLite](https://www.sqlite.org/index.html) を使う。

データソースへのアクセス用に Web API エンドポイントを用意。

API定義は [OpenAPI v3](https://swagger.io/specification/) に準拠。

Yamlは、[これ](./schema/openapi.yml) （プレビュー用にdockerfile書こうと思ったけどIDEのプレビュー機能で十分だったので省略）。

[swagger.ioのエディタ](https://editor-next.swagger.io/) に貼ってもいい。

![こんな感じ](pics/cvapi.png)

Oops. 未実装が目立つ・・・。

## function

※ローカル実行は [go (v1.19)](https://go.dev/) インストール済みが前提。

cv-admin ディレクトリ直下で以下を叩く。

### server

ローカルでAPIサーバー起動。

```
go run src/main.go server
```

### migrate

[ent.(ORM)](https://entgo.io/ja/) で定義したテーブル定義を [SQLite](https://www.sqlite.org/index.html) ファイルに反映する。

```
go run src/main.go migrate
```

### export

APIサーバーにアクセスして取得したデータソースをJSON形式で指定先にエクスポートする。

これを使って [cv](https://github.com/sky0621/cv) リポジトリのフロントエンド機能がキャリアシートのWebページを表示する。

※APIサーバーが起動していることが前提。

※今は、エクスポート対象となる「ユーザーID」が [ソースに直書き](src/cmd/export.go) という、いただけない作り。

```
go run src/main.go export --dir ここに出力先パスを指定！
```

## env

### OS

```
$ cat /etc/os-release 
PRETTY_NAME="Ubuntu 22.04.1 LTS"
```

### go

```
$ go version
go version go1.19 linux/amd64
```

## setup

### project

```
$ go mod init github.com/sky0621/cv-admin
go: creating new go.mod: module github.com/sky0621/cv-admin
```

### open-api-codegen

https://github.com/getkin/kin-openapi

https://github.com/deepmap/oapi-codegen

```
go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
```

### ent

```
cd src/adapter/gateway/ent
go run -mod=mod entgo.io/ent/cmd/ent init User UserActivity UserQualification UserCareerGroup UserCareer UserCareerDescription UserCareerPeriod CareerTask CareerSkillGroup CareerSkill UserNote UserNoteItem
```

### cobra

https://github.com/spf13/cobra

https://github.com/spf13/cobra-cli/blob/main/README.md

```
go install github.com/spf13/cobra-cli@latest
cobra-cli init
```

#### add command

```
cd src
cobra-cli add server
```

### ozzo-validation

https://github.com/go-ozzo/ozzo-validation