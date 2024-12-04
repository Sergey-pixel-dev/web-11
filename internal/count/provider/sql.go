package provider

import (
	"strconv"
)

func (p *Provider) GetCount() (string, error) {
	var msg int
	row := p.conn.QueryRow("SELECT count FROM labs order by id LIMIT 1")
	err := row.Scan(&msg)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(msg), nil
}
func (p *Provider) UpdateCount(count string) error {
	c, _ := strconv.Atoi(count)
	_, err := p.conn.Exec("update labs set count = ($1)", c)
	return err
}
