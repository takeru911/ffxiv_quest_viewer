package scraper

import (
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"reflect"
	"strings"
	"testing"
)

func readHtml(fileName string) (*goquery.Selection, error) {
	file, _ := ioutil.ReadFile(fileName)
	stringReader := strings.NewReader(string(file))
	doc, err := goquery.NewDocumentFromReader(stringReader)
	if err != nil {
		return nil, err
	}
	selection := doc.Find("div#main > div#eorzea_db > div.clearfix > div.db_cnts > div.db__l_main")
	return selection, nil
}

func TestParseQuestName(t *testing.T) {
	selection, err := readHtml("../resources/scraper/quest.html")
	if err != nil {
		t.Fatal("tailed test, cannot read test html file")
	}

	questName, err := parseQuestName(selection)
	if err != nil {
		t.Fatalf("failed test, %v", err)
	}
	expected := "冒険者への手引き"
	if questName != expected {
		t.Fatalf("failed test, incorrect result, actual = %v, expected = %v", questName, expected)
	}
}

func TestParseQuestType(t *testing.T) {
	selection, err := readHtml("../resources/scraper/quest.html")
	if err != nil {
		t.Fatal("tailed test, cannot read test html file")
	}
	questType, err := parseQuestType(selection)
	if err != nil {
		t.Fatalf("failed test, %v", err)
	}
	expected := "新生エオルゼア"
	if questType != expected {
		t.Fatalf("failed test, incorrect result, actual = %v, expected = %v", questType, expected)
	}
}

func TestParseQuestClient(t *testing.T) {
	selection, err := readHtml("../resources/scraper/quest.html")
	if err != nil {
		t.Fatal("tailed test, cannot read test html file")
	}

	questClient, err := parseQuestClient(selection)
	if err != nil {
		t.Fatalf("failed test, %v", err)
	}
	expected := "ミューヌ"
	if questClient != expected {
		t.Fatalf("failed test, incorrect result, actual = %v, expected = %v", questClient, expected)
	}
}

func TestParseQuestPlace(t *testing.T) {
	selection, err := readHtml("../resources/scraper/quest.html")
	if err != nil {
		t.Fatal("tailed test, cannot read test html file")
	}
	questClient, err := parseQuestPlace(selection)
	if err != nil {
		t.Fatalf("failed test, %v", err)
	}
	expected := "グリダニア：新市街"
	if questClient != expected {
		t.Fatalf("failed test, incorrect result, actual = %v, expected = %v", questClient, expected)
	}
}

func TestParseQuestXY(t *testing.T) {
	selection, err := readHtml("../resources/scraper/quest.html")
	if err != nil {
		t.Fatal("tailed test, cannot read test html file")
	}
	x, y, err := parseQuestXY(selection)
	if err != nil {
		t.Fatalf("failed test, %v", err)
	}
	expectedX := 11.7
	expectedY := 13.5
	if x != expectedX {
		t.Fatalf("failed test, incorrect result, actual = %v, expected = %v", x, expectedX)
	}

	if y != expectedY {
		t.Fatalf("failed test, incorrect result, actual = %v, expected = %v", y, expectedY)
	}

}

func TestParseQuestConditions(t *testing.T) {

	selection, err := readHtml("../resources/scraper/quest.html")
	if err != nil {
		t.Fatal("tailed test, cannot read test html file")
	}
	condition, err := parseQuestConditions(selection)
	if err != nil {
		t.Fatalf("failed test, %v", err)
	}

	expected := map[string]string{
		"initJob":      "槍術士",
		"class":        "いずれかのクラス・ジョブ Lv 1～",
		"grandCompany": "指定なし",
		"content":      "指定なし",
	}
	if !reflect.DeepEqual(condition, expected) {
		t.Fatalf("failed test, incorrect result, actual = %v, expected = %v", condition, expected)
	}
}

func TestParseQuestReward(t *testing.T) {
	selection, err := readHtml("../resources/scraper/quest.html")
	if err != nil {
		t.Fatal("tailed test, cannot read test html file")
	}
	rewards, err := parseQuestReward(selection)
	if err != nil {
		t.Fatalf("failed test, %v", err)
	}

	expected := map[string]int{
		"exp": 100,
		"gil": 107,
	}
	if !reflect.DeepEqual(rewards, expected) {
		t.Fatalf("failed test, incorrect result, actual = %v, expected = %v", rewards, expected)
	}
}

func TestParsePremisQuests(t *testing.T) {
	selection, err := readHtml("../resources/scraper/quest.html")
	if err != nil {
		t.Fatal("tailed test, cannot read test html file")
	}
	premisQuests, _ := parsePremisQuests(selection)
	expected := []Quest{
		{
			"/lodestone/playguide/db/quest/298088846dc/",
			"森の都グリダニアへ",
			nil,
		},
		{
			"/lodestone/playguide/db/quest/298088846dc/",
			"森の都グリダニアへ2",
			nil,
		},
	}

	if !reflect.DeepEqual(premisQuests, expected) {
		t.Fatalf("failed test, incorrect result, actual = %v, expected = %v", premisQuests, expected)
	}

}

func TestParsePremisQuestsNothing(t *testing.T) {
	selection, err := readHtml("../resources/scraper/quest_simple.html")
	if err != nil {
		t.Fatal("tailed test, cannot read test html file")
	}
	premisQuests, _ := parsePremisQuests(selection)
	expected := []Quest{}

	if !reflect.DeepEqual(premisQuests, expected) {
		t.Fatalf("failed test, incorrect result, actual = %v, expected = %v", premisQuests, expected)
	}

}

func TestParseUnlockQuests(t *testing.T) {
	selection, err := readHtml("../resources/scraper/quest.html")
	if err != nil {
		t.Fatal("tailed test, cannot read test html file")
	}
	unlockQuests, _ := parseUnlockQuests(selection)
	expected := []Quest{
		{
			"/lodestone/playguide/db/quest/088a43daa15/",
			"バノック練兵所へ",
			nil,
		},
		{
			"/lodestone/playguide/db/quest/2ee087b785d/",
			"ツリースピーク厩舎の金具",
			nil,
		},
		{
			"/lodestone/playguide/db/quest/8d6c6da2282/",
			"小さな預かり物",
			nil,
		},
	}

	if !reflect.DeepEqual(unlockQuests, expected) {
		t.Fatalf("failed test, incorrect result, actual = %v, expected = %v", unlockQuests, expected)
	}

}

func TestParseUnlockQuestsNothing(t *testing.T) {
	selection, err := readHtml("../resources/scraper/quest_simple.html")
	if err != nil {
		t.Fatal("tailed test, cannot read test html file")
	}
	unlockQuests, _ := parseUnlockQuests(selection)
	expected := []Quest{}

	if !reflect.DeepEqual(unlockQuests, expected) {
		t.Fatalf("failed test, incorrect result, actual = %v, expected = %v", unlockQuests, expected)
	}

}

func TestSelectReward(t *testing.T) {
	selection, err := readHtml("../resources/scraper/quest.html")
	if err != nil {
		t.Fatal("tailed test, cannot read test html file")
	}
	selectReward, _ := parseSelectReward(selection)
	expected := []Item{
		{
			"/lodestone/playguide/db/item/96029646ce8/",
			"カリガ",
		},
		{
			"/lodestone/playguide/db/item/76059262ca9/",
			"ハードレザーサンダル",
		},
		{
			"/lodestone/playguide/db/item/ce2e67f0442/",
			"アラグ銅貨",
		},
	}

	if !reflect.DeepEqual(selectReward, expected) {
		t.Fatalf("failed test, incorrect result, actual = %v, expected = %v", selectReward, expected)
	}
}

func TestSelectRewardNothing(t *testing.T) {
	selection, err := readHtml("../resources/scraper/quest_simple.html")
	if err != nil {
		t.Fatal("tailed test, cannot read test html file")
	}
	selectReward, _ := parseSelectReward(selection)
	expected := []Item{}

	if !reflect.DeepEqual(selectReward, expected) {
		t.Fatalf("failed test, incorrect result, actual = %v, expected = %v", selectReward, expected)
	}
}
