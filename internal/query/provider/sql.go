package provider

import (
	"database/sql"
	"time"
)

func (p *Provider) GetTimeLastVisit(name string) (string, bool, error) {
	var msg string
	row := p.conn.QueryRow("SELECT time FROM labs where name=($1)", name) //чтоб наверняка последнее взяли
	err := row.Scan(&msg)
	if err == sql.ErrNoRows {
		return "Hello, " + name, false, nil
	}
	if err != nil {
		return "", false, err
	}
	return "Hello, " + name + " Your last visit: " + msg, true, nil
}

func (p *Provider) UpdateTimeLastVisit(name string) error {
	_, err := p.conn.Exec("update labs set time = ($1) where name = ($2)", time.Now().Format("2006-01-02 15:04:05"), name)
	return err
}
func (p *Provider) SetTimeLastVisit(name string) error {
	_, err := p.conn.Exec("insert into labs (count, name, time) values (0, ($1), ($2))",
		name, time.Now().Format("2006-01-02 15:04:05"))
	return err
}
