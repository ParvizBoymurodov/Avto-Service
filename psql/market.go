package psql

import (
	"avtoService/models"
	"context"
	"github.com/pkg/errors"
	"log"
)

func (d *DB) AddMarket(ctx context.Context, mrk models.Market) (err error) {
	_, err = d.conn.Exec(ctx, `insert into  servicemarket (market_name, service_id, title) VALUES ($1, $2, $3)`, mrk.Name, mrk.ServiceId, mrk.Title)
	if err != nil {
		log.Println("Can't add market: ", err)
		return errors.Wrap(err, "Can't add market to BD")
	}
	return
}

func (d *DB) MarketList(ctx context.Context) (resp []models.Market, err error) {
	rows, err := d.conn.Query(ctx, `select sm.id,sm.market_name,  sm.title, s.service_name, s.region from servicemarket sm inner join services s on s.id = sm.service_id where sm.remove = false and s.remove = false`)
	if err != nil {
		log.Println("Can't get market list: ", err)
		return nil, errors.Wrap(err, "Can't get market list")
	}

	defer rows.Close()

	resp = make([]models.Market, 0)

	for rows.Next() {
		item := models.Market{}
		err := rows.Scan(&item.Id, &item.Name, &item.Title, &item.Service.Name, &item.Service.Region)
		if err != nil {
			return nil, errors.Wrap(err, "Can't scan MarketList")
		}

		resp = append(resp, item)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return
}

func (d *DB) RemoveMarket(ctx context.Context, id int64) (err error) {
	_, err = d.conn.Exec(ctx, `update servicemarket set remove = true where id= $1`, id)
	return
}