package data

import (
	"github.com/DOGTT/dm-api-server/internal/conf"
	"github.com/DOGTT/dm-api-server/internal/data/rds"
)

type DataEntry struct {
	rds *rds.Client
}

func New(conf *conf.DataConfig) (d *DataEntry, err error) {
	d = &DataEntry{}
	d.rds, err = rds.New(conf.RDS)
	if err != nil {
		return
	}

	return
}
