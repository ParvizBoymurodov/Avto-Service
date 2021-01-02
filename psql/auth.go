package psql

import (
	"avtoService/models"
	"context"
	"github.com/pkg/errors"
	"log"
)



func (d *DB) Registered(ctx context.Context,reg models.Reg) (err error) {
	_, err = d.conn.Exec(ctx, `insert into registered (name, surname, login, password) VALUES ($1, $2, $3, md5($4))`, reg.Name, reg.Surname, reg.Login, reg.Password)
	if err != nil {
		log.Println("Can't registered: ", err)
		return errors.Wrap(err, "Can't add person to BD")
	}
	return nil
}

func (d *DB) Login(ctx context.Context, reg models.Login) (login string, err error) {
	err = d.conn.QueryRow(ctx, `select login from registered where login = $1`, reg.Login).Scan(&login)
	return
}

func (d *DB) ChangePass(ctx context.Context, id int64, chg models.ChangePass) (err error) {
	_, err = d.conn.Exec(ctx, `update registered set password = md5($1) where id= $2`, chg.Password, id)
	return
}

func (d *DB) PersonalList (ctx context.Context) (list []models.Reg, err error)  {
	rows, err := d.conn.Query(ctx, `select name, surname, login from registered where remove = false`)
	if err != nil {
		log.Println("Can't get personal list: ", err)
		return nil, errors.Wrap(err, "Can't get personal list")
	}

	defer rows.Close()

	list = make([]models.Reg, 0)
	for rows.Next() {
		item := models.Reg{}
		err := rows.Scan(&item.Name, &item.Surname, &item.Login)
		if err != nil {
			return nil, errors.Wrap(err, "Can't scan PersonalList")
		}
		list = append(list, item)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return
}

func (d *DB) DeletePersonal(ctx context.Context, id int64) (err error) {
	_, err = d.conn.Exec(ctx, `update registered set remove = true where id = $1`, id)
	return
}