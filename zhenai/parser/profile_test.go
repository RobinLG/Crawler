package parser

import (
	"testing"
	"robin/learngo/crawler/model"
	"io/ioutil"
)

func TestParserProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.txt")
	//contents, err := fetcher.Fethch("http://album.zhenai.com/u/1626200317")

	if err != nil {
		panic(err)
	}

	result := ParserProfile(contents, "风中的蒲公英")

	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 element; but was %v", result.Items)
	}

	profile := result.Items[0].(model.Profile)

	expected := model.Profile{
		"风中的蒲公英",
		"女",
		42,
		158,
		48,
		"3001-5000元",
		"离异",
		"中专",
		"公务员",
		"四川阿坝",
		"处女座",
		"已购房",
		"未购车"}

	if profile != expected {
		t.Errorf("expected profile : %v; but was %v", profile, expected)
	}

}
