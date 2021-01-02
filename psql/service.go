package psql

import (
	"avtoService/models"
	"context"
	"github.com/pkg/errors"
	"log"
)

func (d *DB) AddService(ctx context.Context, svc models.Svc) (err error) {
	_, err = d.conn.Exec(ctx, `insert into services (service_name, region) values ($1, $2)`, svc.Name, svc.Region)
	if err != nil {
		log.Println("Can't add service: ", err)
		return errors.Wrap(err, "Can't add service to BD")
	}
	return
}

func (d *DB) ServiceList(ctx context.Context) (resp []models.SelectSvc, err error) {
	rows, err := d.conn.Query(ctx, `select id, service_name, region from services where remove = false`)
	if err != nil {
		log.Println("Can't get service list: ", err)
		return nil, errors.Wrap(err, "Can't get service list")
	}

	defer rows.Close()

	resp = make([]models.SelectSvc, 0)

	for rows.Next() {
		item := models.SelectSvc{}
		err = rows.Scan(&item.Id, &item.Name, &item.Region)
		if err != nil {
			return nil, errors.Wrap(err, "Can't scan ServiceList")
		}

		resp = append(resp, item)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return
}

func (d *DB) EditService(ctx context.Context, svc models.Svc, id int64) (err error) {
	_, err = d.conn.Exec(ctx, `update services set service_name = $1, region=$2 where id = $3`, svc.Name, svc.Region, id)
	return
}

func (d *DB) RemoveService(ctx context.Context, id int64) (err error) {
	_, err = d.conn.Exec(ctx, `update services set remove = true where id=$1`, id)
	return
}