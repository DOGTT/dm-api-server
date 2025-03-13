package utils

import (
	"fmt"

	api "github.com/DOGTT/dm-api-server/api/base"
)

func PointCoordToGeometry(p *api.PointCoord) string {
	if p == nil {
		return ""
	}
	return fmt.Sprintf("SRID=4326;POINT(%f %f)", p.Lng, p.Lat)
}

func PointCoordFromGeometry(s string) *api.PointCoord {
	var lat, lng float32
	_, _ = fmt.Sscanf(string(s), "POINT(%f %f)", &lng, &lat) // 经度在前，纬度在后
	return &api.PointCoord{
		Lat: lat,
		Lng: lng,
	}
}
