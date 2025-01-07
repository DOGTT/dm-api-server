package data

import (
	"github.com/DOGTT/dm-api-server/internal/conf"
	"github.com/DOGTT/dm-api-server/internal/data/fds"
	"github.com/DOGTT/dm-api-server/internal/data/mapdata"
	"github.com/DOGTT/dm-api-server/internal/data/rds"
)

type DataEntry struct {
	*rds.RDSClient
	*fds.FDSClient
	*mapdata.MapApiHandler
}

func New(conf *conf.DataConfig) (d *DataEntry, err error) {
	d = &DataEntry{}
	d.RDSClient, err = rds.New(conf.RDS)
	if err != nil {
		return
	}
	d.FDSClient, err = fds.New(conf.FDS)
	if err != nil {
		return
	}
	d.MapApiHandler, err = mapdata.New(conf.MapData)
	if err != nil {
		return
	}
	return
}
