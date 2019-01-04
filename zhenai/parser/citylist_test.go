package parser

import (
	"testing"
	"robin/learngo/crawler/fetcher"
)

func TestParseCityList(t *testing.T) {
	//以下这行代码虽是从本地那文件但是UTF8编码问题还未解决
	//contents, err := ioutil.ReadFile("citylist_test_data.html")
	contents, err := fetcher.Fethch("http://www.zhenai.com/zhenghun")

	if err != nil {
		panic(err)
	}

	result := ParseCityList(contents)

	const resultSize = 470
	expectedUrls := []string{"http://www.zhenai.com/zhenghun/aba", "http://www.zhenai.com/zhenghun/akesu", "http://www.zhenai.com/zhenghun/alashanmeng",}
	expectedCities := []string{"City 阿坝", "City 阿克苏", "City 阿拉善盟",}

	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d requests; but had %d", resultSize, len(result.Requests))
	}
	for  i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s; but was %s", i, url, result.Requests[i].Url)
		}
	}

	if len(result.Items) != resultSize {
		t.Errorf("result should have %d requests; but had %d", resultSize, len(result.Items))
	}
	for  i, city := range expectedCities {
		if result.Items[i].(string) != city {
			t.Errorf("expected item #%d: %s; but was %s", i, city, result.Items[i].(string))
		}
	}
}
