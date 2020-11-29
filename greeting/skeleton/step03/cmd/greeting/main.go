// STEP03: パッケージを分けてみよう

// mainパッケージの定義
package main

// TODO: greetingパッケージのインポート
import "skeleton/step03/cmd/greeting"

// main関数から実行される
func main() {
	// TODO: greeting.Do関数を呼び出す
	greeting.Do()
}
