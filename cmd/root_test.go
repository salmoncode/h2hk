/*
Copyright © 2025 salmoncode <salmon2073@gmail.com>
*/
package cmd

import (
	"testing"
)

func TestConvertHiraganaToKatakana(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "basic hiragana",
			input:    "あいうえお",
			expected: "アイウエオ",
		},
		{
			name:     "mixed hiragana",
			input:    "こんにちは",
			expected: "コンニチハ",
		},
		{
			name:     "hiragana with non-hiragana characters",
			input:    "たなか　たろう",
			expected: "タナカ　タロウ",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "non-hiragana only",
			input:    "ABC123",
			expected: "ABC123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := convertHiraganaToKatakana(tt.input)
			if result != tt.expected {
				t.Errorf("convertHiraganaToKatakana(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestConvertToHankakuKatakana(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "hiragana to hankaku katakana",
			input:    "こんにちは",
			expected: "ｺﾝﾆﾁﾊ",
		},
		{
			name:     "hiragana with full-width space",
			input:    "たなか　たろう",
			expected: "ﾀﾅｶ ﾀﾛｳ",
		},
		{
			name:     "full-width katakana to half-width",
			input:    "アイウエオ",
			expected: "ｱｲｳｴｵ",
		},
		{
			name:     "voiced katakana",
			input:    "ガギグゲゴ",
			expected: "ｶﾞｷﾞｸﾞｹﾞｺﾞ",
		},
		{
			name:     "semi-voiced katakana",
			input:    "パピプペポ",
			expected: "ﾊﾟﾋﾟﾌﾟﾍﾟﾎﾟ",
		},
		{
			name:     "small katakana",
			input:    "ァィゥェォ",
			expected: "ｧｨｩｪｫ",
		},
		{
			name:     "punctuation",
			input:    "。、・",
			expected: "｡､･",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := convertToHankakuKatakana(tt.input)
			if result != tt.expected {
				t.Errorf("convertToHankakuKatakana(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestBuildZenkakuToHankakuReplacer(t *testing.T) {
	replacer := buildZenkakuToHankakuReplacer()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "basic katakana",
			input:    "アイウエオ",
			expected: "ｱｲｳｴｵ",
		},
		{
			name:     "voiced sounds",
			input:    "ガザダバ",
			expected: "ｶﾞｻﾞﾀﾞﾊﾞ",
		},
		{
			name:     "semi-voiced sounds",
			input:    "パピプペポ",
			expected: "ﾊﾟﾋﾟﾌﾟﾍﾟﾎﾟ",
		},
		{
			name:     "small characters",
			input:    "ャュョッ",
			expected: "ｬｭｮｯ",
		},
		{
			name:     "symbols",
			input:    "ー。、・",
			expected: "ｰ｡､･",
		},
		{
			name:     "full-width space to half-width",
			input:    "　",
			expected: " ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := replacer.Replace(tt.input)
			if result != tt.expected {
				t.Errorf("replacer.Replace(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}
