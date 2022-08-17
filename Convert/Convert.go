package Convert

import (
	"bufio"
	"fmt"
	"github.com/ikawaha/kagome-dict/ipa"
	"github.com/ikawaha/kagome/v2/filter"
	"github.com/ikawaha/kagome/v2/tokenizer"
	"strings"
)

var ichi = map[string]string{"ワタシ": "", "ボク": "", "アタシ": "",
	"ワタクシ": "", "オレ": "", "ジブン": "", "アッシ": "",
	"アタイ": "", "ウチ": "", "ワシ": "", "オラ": "", "オイラ": "",
	"コチラ": "", "アチキ": "", "ワガハイ": ""}
var ni = map[string]string{"アナタ": "", "キミ": "", "アンタ": "",
	"ソチラ": "", "オマエ": "", "キサマ": "", "ソナタ": "",
	"ナンジ": "", "ウヌ": "", "オタク": "", "オヌシ": ""}

func Convert(originalText string) string {

	t, err := tokenizer.New(ipa.Dict(), tokenizer.OmitBosEos())
	if err != nil {
		panic(err)
	}

	//sampleText := `恥の多い生涯を送って来ました。自分には、人間の生活というものが、見当つかないのです。自分は東北の田舎に生れましたので、汽車をはじめて見たのは、よほど大きくなってからでした。自分は停車場のブリッジを、上って、降りて、そうしてそれが線路をまたぎ越えるために造られたものだという事には全然気づかず、ただそれは停車場の構内を外国の遊戯場みたいに、複雑に楽しく、ハイカラにするためにのみ、設備せられてあるものだとばかり思っていました。しかも、かなり永い間そう思っていたのです。ブリッジの上ったり降りたりは、自分にはむしろ、ずいぶん垢抜けのした遊戯で、それは鉄道のサーヴィスの中でも、最も気のきいたサーヴィスの一つだと思っていたのですが、のちにそれはただ旅客が線路をまたぎ越えるための頗る実利的な階段に過ぎないのを発見して、にわかに興が覚めました。また、自分は子供の頃、絵本で地下鉄道というものを見て、これもやは`

	scanner := bufio.NewScanner(strings.NewReader(originalText))
	scanner.Split(filter.ScanSentences)
	var result string
	for scanner.Scan() {
		// tokenize
		tokens := t.Tokenize(scanner.Text())

		for index, token := range tokens {
			fmt.Printf("%s\t%s\t%s\n", token.Surface, token.Features()[0], token.Features()[1])
			if token.Features()[1] == "代名詞" {
				S, _ := token.Pronunciation()
				if _, ok := ichi[S]; ok {
					token.Surface = "まろ"
				}
				if _, ok := ni[S]; ok {
					token.Surface = "そなた"
				}
			}
			if token.Features()[1] == "終助詞" {
				token.Surface = "のぅ"
			}
			if token.Features()[0] == "形容詞" {
				if index == len(tokens)-1 {
					token.Surface += "のぅ"
				} else {
					if tokens[index+1].Surface == "。" {
						token.Surface += "のぅ"
					}
				}
			}
			if token.Features()[0] == "助動詞" {
				if token.Surface == "です" || token.Surface == "ます" {
					token.Surface = "のぅ"
				} else {
					if index == len(tokens)-1 {
						token.Surface += "のぅ"
					} else {
						if tokens[index+1].Surface == "。" {
							token.Surface += "のぅ"
						}
					}
				}
			}
			if token.POS()[0] == "動詞" {
				if index < len(tokens)-1 {
					if tokens[index+1].Surface == "です" || tokens[index+1].Surface == "ます" {
						token.Surface, _ = token.BaseForm()
					}
					if tokens[index+1].Surface == "。" {
						token.Surface += "のぅ"
					}
				} else {
					token.Surface += "のぅ"
				}
			}
			result += token.Surface
		}
		/*
				if token.Features()[1] == "代名詞" {
					S, _ := token.Pronunciation()
					if _, ok := shugo[S]; ok {
						token.Surface = "ワガハイ"
					}
				}
				if token.Features()[1] == "終助詞" {
					token.Surface = "ナリ"
				}
				if token.Features()[0] == "形容詞" {
					if index == len(tokens)-1 {
						token.Surface += "ナリ"
					} else {
						if tokens[index+1].Surface == "。" {
							token.Surface += "ナリ"
						}
					}
				}
				if token.Features()[0] == "助動詞" {
					if token.Surface == "です" || token.Surface == "ます" {
						token.Surface = "ナリ"
					} else {
						if index == len(tokens)-1 {
							token.Surface += "ナリ"
						} else {
							if tokens[index+1].Surface == "。" {
								token.Surface += "ナリ"
							}
						}
					}
				}
			if token.Features()[0] == "感動詞" {
						if index == len(tokens)-1 {
							token.Surface += "ナリ"
						}
				if token.POS()[0] == "動詞" {
					if index < len(tokens)-1 {
						if tokens[index+1].Surface == "です" || tokens[index+1].Surface == "ます" {
							token.Surface, _ = token.BaseForm()
						}
						if tokens[index+1].Surface == "。" {
							token.Surface += "ナリ"
						}
					} else {
						token.Surface += "ナリ"
					}
				}
				result += token.Surface
			}*/
	}
	return result
}
