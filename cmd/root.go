/*
Copyright © 2025 salmoncode <salmon2073@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var zenkakuToHankaku = strings.NewReplacer(
	"ア", "ｱ", "イ", "ｲ", "ウ", "ｳ", "エ", "ｴ", "オ", "ｵ",
	"カ", "ｶ", "キ", "ｷ", "ク", "ｸ", "ケ", "ｹ", "コ", "ｺ",
	"サ", "ｻ", "シ", "ｼ", "ス", "ｽ", "セ", "ｾ", "ソ", "ｿ",
	"タ", "ﾀ", "チ", "ﾁ", "ツ", "ﾂ", "テ", "ﾃ", "ト", "ﾄ",
	"ナ", "ﾅ", "ニ", "ﾆ", "ヌ", "ﾇ", "ネ", "ﾈ", "ノ", "ﾉ",
	"ハ", "ﾊ", "ヒ", "ﾋ", "フ", "ﾌ", "ヘ", "ﾍ", "ホ", "ﾎ",
	"マ", "ﾏ", "ミ", "ﾐ", "ム", "ﾑ", "メ", "ﾒ", "モ", "ﾓ",
	"ヤ", "ﾔ", "ユ", "ﾕ", "ヨ", "ﾖ",
	"ラ", "ﾗ", "リ", "ﾘ", "ル", "ﾙ", "レ", "ﾚ", "ロ", "ﾛ",
	"ワ", "ﾜ", "ヲ", "ｦ", "ン", "ﾝ",
	"ガ", "ｶﾞ", "ギ", "ｷﾞ", "グ", "ｸﾞ", "ゲ", "ｹﾞ", "ゴ", "ｺﾞ",
	"ザ", "ｻﾞ", "ジ", "ｼﾞ", "ズ", "ｽﾞ", "ゼ", "ｾﾞ", "ゾ", "ｿﾞ",
	"ダ", "ﾀﾞ", "ヂ", "ﾁﾞ", "ヅ", "ﾂﾞ", "デ", "ﾃﾞ", "ド", "ﾄﾞ",
	"バ", "ﾊﾞ", "ビ", "ﾋﾞ", "ブ", "ﾌﾞ", "ベ", "ﾍﾞ", "ボ", "ﾎﾞ",
	"パ", "ﾊﾟ", "ピ", "ﾋﾟ", "プ", "ﾌﾟ", "ペ", "ﾍﾟ", "ポ", "ﾎﾟ",
	"ァ", "ｧ", "ィ", "ｨ", "ゥ", "ｩ", "ェ", "ｪ", "ォ", "ｫ",
	"ャ", "ｬ", "ュ", "ｭ", "ョ", "ｮ", "ッ", "ｯ",
	"ー", "ｰ", "。", "｡", "、", "､", "・", "･",
	"「", "｢", "」", "｣", "゛", "ﾞ", "゜", "ﾟ", "　", " ",
)

var rootCmd = &cobra.Command{
	Use:   "h2hk",
	Short: "Convert Hiragana to half-width Katakana",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Usage: h2hk <text>")
			return
		}
		var result strings.Builder
		for _, r := range args[0] {
			if r >= 'ぁ' && r <= 'ゖ' {
				r += 0x60
			}
			result.WriteString(string(r))
		}
		fmt.Println(zenkakuToHankaku.Replace(result.String()))
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
