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

// buildZenkakuToHankakuReplacer creates a replacer for full-width to half-width katakana conversion.
// This function reduces code duplication by generating the replacer from structured data.
func buildZenkakuToHankakuReplacer() *strings.Replacer {
	// Basic katakana mappings (full-width -> half-width)
	basicMappings := map[string]string{
		"ア": "ｱ", "イ": "ｲ", "ウ": "ｳ", "エ": "ｴ", "オ": "ｵ",
		"カ": "ｶ", "キ": "ｷ", "ク": "ｸ", "ケ": "ｹ", "コ": "ｺ",
		"サ": "ｻ", "シ": "ｼ", "ス": "ｽ", "セ": "ｾ", "ソ": "ｿ",
		"タ": "ﾀ", "チ": "ﾁ", "ツ": "ﾂ", "テ": "ﾃ", "ト": "ﾄ",
		"ナ": "ﾅ", "ニ": "ﾆ", "ヌ": "ﾇ", "ネ": "ﾈ", "ノ": "ﾉ",
		"ハ": "ﾊ", "ヒ": "ﾋ", "フ": "ﾌ", "ヘ": "ﾍ", "ホ": "ﾎ",
		"マ": "ﾏ", "ミ": "ﾐ", "ム": "ﾑ", "メ": "ﾒ", "モ": "ﾓ",
		"ヤ": "ﾔ", "ユ": "ﾕ", "ヨ": "ﾖ",
		"ラ": "ﾗ", "リ": "ﾘ", "ル": "ﾙ", "レ": "ﾚ", "ロ": "ﾛ",
		"ワ": "ﾜ", "ヲ": "ｦ", "ン": "ﾝ",
	}

	// Voiced katakana mappings (dakuten)
	voicedMappings := map[string]string{
		"ガ": "ｶﾞ", "ギ": "ｷﾞ", "グ": "ｸﾞ", "ゲ": "ｹﾞ", "ゴ": "ｺﾞ",
		"ザ": "ｻﾞ", "ジ": "ｼﾞ", "ズ": "ｽﾞ", "ゼ": "ｾﾞ", "ゾ": "ｿﾞ",
		"ダ": "ﾀﾞ", "ヂ": "ﾁﾞ", "ヅ": "ﾂﾞ", "デ": "ﾃﾞ", "ド": "ﾄﾞ",
		"バ": "ﾊﾞ", "ビ": "ﾋﾞ", "ブ": "ﾌﾞ", "ベ": "ﾍﾞ", "ボ": "ﾎﾞ",
	}

	// Semi-voiced katakana mappings (handakuten)
	semiVoicedMappings := map[string]string{
		"パ": "ﾊﾟ", "ピ": "ﾋﾟ", "プ": "ﾌﾟ", "ペ": "ﾍﾟ", "ポ": "ﾎﾟ",
	}

	// Small katakana mappings
	smallMappings := map[string]string{
		"ァ": "ｧ", "ィ": "ｨ", "ゥ": "ｩ", "ェ": "ｪ", "ォ": "ｫ",
		"ャ": "ｬ", "ュ": "ｭ", "ョ": "ｮ", "ッ": "ｯ",
	}

	// Punctuation and symbol mappings
	symbolMappings := map[string]string{
		"ー": "ｰ", "。": "｡", "、": "､", "・": "･",
		"「": "｢", "」": "｣", "゛": "ﾞ", "゜": "ﾟ", "　": " ",
	}

	// Helper function to add mappings to the replacements slice
	addMappings := func(replacements []string, mappings map[string]string) []string {
		for k, v := range mappings {
			replacements = append(replacements, k, v)
		}
		return replacements
	}

	// Calculate total capacity needed (2 strings per mapping: key + value)
	totalMappings := len(basicMappings) + len(voicedMappings) + len(semiVoicedMappings) + len(smallMappings) + len(symbolMappings)
	replacements := make([]string, 0, totalMappings*2)

	// Combine all mappings into a single slice for the replacer
	replacements = addMappings(replacements, basicMappings)
	replacements = addMappings(replacements, voicedMappings)
	replacements = addMappings(replacements, semiVoicedMappings)
	replacements = addMappings(replacements, smallMappings)
	replacements = addMappings(replacements, symbolMappings)

	return strings.NewReplacer(replacements...)
}

var zenkakuToHankaku = buildZenkakuToHankakuReplacer()

// convertHiraganaToKatakana converts hiragana characters to katakana.
// This function encapsulates the conversion logic to improve testability and reusability.
func convertHiraganaToKatakana(input string) string {
	var result strings.Builder
	for _, r := range input {
		// Convert hiragana to katakana by adding 0x60 to the character code
		if r >= 'ぁ' && r <= 'ゖ' {
			r += 0x60
		}
		result.WriteRune(r)
	}
	return result.String()
}

// convertToHankakuKatakana converts input text to half-width katakana.
// This is the main conversion function that combines hiragana->katakana and zenkaku->hankaku conversions.
func convertToHankakuKatakana(input string) string {
	katakana := convertHiraganaToKatakana(input)
	return zenkakuToHankaku.Replace(katakana)
}

var rootCmd = &cobra.Command{
	Use:   "h2hk",
	Short: "Convert Hiragana to half-width Katakana",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Usage: h2hk <text>")
			return
		}
		fmt.Println(convertToHankakuKatakana(args[0]))
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
