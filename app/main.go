package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"

	"github.com/zawa-t/pr-commentator/flag"
	"github.com/zawa-t/pr-commentator/golangci"
	"github.com/zawa-t/pr-commentator/platform"
	"github.com/zawa-t/pr-commentator/platform/bitbucket"
	"github.com/zawa-t/pr-commentator/platform/bitbucket/client"
	"github.com/zawa-t/pr-commentator/platform/http"
	"github.com/zawa-t/pr-commentator/test/custommock"
	"github.com/zawa-t/pr-commentator/txt"
)

/*
以下、動作確認用コマンド
```
$ go build -o pr-comment
$ ./pr-comment -n=golangci-lint -ext=json < sample.json
```
*/

/*
TODO:
*/

func main() {
	stdin := os.Stdin
	stat, err := os.Stdin.Stat() // MEMO: 標準入力の「ファイル情報（ファイルのモードやサイズ、変更日時など）」取得
	if err != nil {
		slog.Error("Stdin could not be verified.", "error", err.Error())
		os.Exit(1)
	}

	// MEMO:
	// stat.Mode()を実行することでファイルのモード情報（ファイルの種類やアクセス権）を取得。それによって設定される os.ModeCharDevice の値を用いて、
	// 入力がキャラクタデバイス（通常、ターミナル）であるか否かを確認。現時点では、標準入力がパイプやリダイレクトのみ受け付けたいため、ターミナルからの入力の場合（0 でない場合）は処理終了。
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		slog.Error("Only data from standard input can be accepted.")
		os.Exit(1)
	}

	flagValue := flag.NewValue()
	input := platform.Input{
		Name: flagValue.Name,
	}

	switch flagValue.Name {
	case "golangci-lint":
		input.Datas = golangci.MakeInputDatas(*flagValue, stdin)
	default:
		input.Datas = txt.Read(*flagValue, stdin)
	}

	slog.Info("The following data was accepted.")
	printJSON(input)

	var customClient bitbucket.CustomClient
	if flagValue.Reporter != nil && *flagValue.Reporter == "local" {
		customClient = custommock.DefaultCustomClient
	} else {
		customClient = client.NewCustomClient(http.NewClient())
	}

	if err := bitbucket.NewPullRequest(customClient).AddComments(context.Background(), input); err != nil {
		slog.Error("Failed to add comments.", "error", err.Error())
		os.Exit(1)
	}

	slog.Info("The pull request comments were successfully added.")
}

func printJSON(v any) {
	data, err := json.Marshal(v)
	if err != nil {
		slog.Error("Faild to exec json.Marshal().", "error", err.Error())
		os.Exit(1)
	}
	var buf bytes.Buffer
	err = json.Indent(&buf, data, "", "  ")
	if err != nil {
		slog.Error("Faild to exec json.Indent().", "error", err.Error())
		os.Exit(1)
	}
	fmt.Println(buf.String())
}
