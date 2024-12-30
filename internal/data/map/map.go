package mapapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
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
}

// TencentMapSearch 调用腾讯地图API搜索地点
func (d *MapApiHandler) MapSearch(ctx context.Context, keyword string, bound string) ([]MapSearchResult, error) {
	// TODO: 实现腾讯地图API调用逻辑
	// 构建请求参数
	params := map[string]string{
		"key":        "你的腾讯地图API密钥",
		"keyword":    keyword,
		"boundary":   bound,
		"page_size":  "20",
		"page_index": "1",
	}

	// 发起HTTP请求
	url := "https://apis.map.qq.com/ws/place/v1/search"
	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}

	// 添加查询参数
	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %v", err)
	}
	defer resp.Body.Close()

	// 解析响应
	var result struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
		Data    []struct {
			ID       string `json:"id"`
			Title    string `json:"title"`
			Address  string `json:"address"`
			Location struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
		} `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %v", err)
	}

	if result.Status != 0 {
		return nil, fmt.Errorf("API调用失败: %s", result.Message)
	}

	// 转换为MapSearchResult
	var searchResults []MapSearchResult
	for _, item := range result.Data {
		searchResults = append(searchResults, MapSearchResult{
			ID:      item.ID,
			Title:   item.Title,
			Address: item.Address,
			Location: struct {
				Lng float64
				Lat float64
			}{
				Lng: item.Location.Lng,
				Lat: item.Location.Lat,
			},
		})
	}
	return nil, nil
}
