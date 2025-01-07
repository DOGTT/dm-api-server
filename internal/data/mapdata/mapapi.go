package mapdata

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DOGTT/dm-api-server/internal/conf"
)

// MapSearchResult 地图搜索结果
type MapSearchResult struct {
	ID       string
	Title    string
	Address  string
	Location struct {
		Lng float64
		Lat float64
	}
}

type MapApiHandler struct {
	conf *conf.MapDataConfig
}

func New(conf *conf.MapDataConfig) (d *MapApiHandler, err error) {
	d = &MapApiHandler{
		conf: conf,
	}
	return
}

// TencentMapSearch 调用高德地图API搜索地点
func (d *MapApiHandler) MapSearch(ctx context.Context, keyword string) ([]MapSearchResult, error) {

	// 构建请求URL
	url := fmt.Sprintf("%s/place/text?key=%s&keywords=%s&city=全国&children=1&offset=20&page=1&extensions=base",
		d.conf.Endpoint, d.conf.Key, keyword)

	// 发送HTTP GET请求
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("请求高德地图API失败: %v", err)
	}
	defer resp.Body.Close()

	// 解析响应
	var result struct {
		Status string `json:"status"`
		Pois   []struct {
			ID       string `json:"id"`
			Name     string `json:"name"`
			Address  string `json:"address"`
			Location string `json:"location"`
		} `json:"pois"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %v", err)
	}

	if result.Status != "1" {
		return nil, fmt.Errorf("高德地图API返回错误")
	}

	// 转换为MapSearchResult格式
	var searchResults []MapSearchResult
	for _, poi := range result.Pois {
		var location struct {
			Lng float64
			Lat float64
		}
		fmt.Sscanf(poi.Location, "%f,%f", &location.Lng, &location.Lat)

		searchResults = append(searchResults, MapSearchResult{
			ID:       poi.ID,
			Title:    poi.Name,
			Address:  poi.Address,
			Location: location,
		})
	}

	return searchResults, nil
}
