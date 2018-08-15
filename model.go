// model.go

package main

import (
	"database/sql"
	"fmt"
	"time"
)

type Client struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Address string    `json:"address"`
	City    string    `json:"city"`
	State   string    `json:"state"`
	ZIP     string    `json:"zip"`
	Web     string    `json:"web"`
	Email   string    `json:"email"`
	Phone   string    `json:"phone"`
	Created time.Time `json:"created"`
}

type Clients []Client

func (c *Client) getClient(db *sql.DB) error {
	statement := fmt.Sprintf("SELECT name, city, state, web, email, phone FROM clients WHERE id=%d", c.ID)
	return db.QueryRow(statement).Scan(&c.Name, &c.City, &c.State, &c.Web, &c.Email, &c.Phone)
}

func (c *Client) updateClient(db *sql.DB) error {
	statement := fmt.Sprintf("UPDATE clients SET name='%s', address='%s', city='%s', state='%s', zip='%s', web='%s',email='%s', phone='%s' WHERE id=%d", c.Name, c.Address, c.City, c.State, c.ZIP, c.Web, c.Email, c.Phone, c.ID)
	_, err := db.Exec(statement)
	return err
}

func (c *Client) deleteClient(db *sql.DB) error {
	statement := fmt.Sprintf("DELETE FROM clients WHERE id=%d", c.ID)
	_, err := db.Exec(statement)
	return err
}

func (c *Client) createClient(db *sql.DB) error {
	statement := fmt.Sprintf("INSERT INTO clients(name, address, city, state, zip, web, email, phone, created) VALUES('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s')", c.Name, c.Address, c.City, c.State, c.ZIP, c.Web, c.Email, c.Phone, c.Created)
	_, err := db.Exec(statement)

	if err != nil {
		return err
	}

	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&c.ID)

	if err != nil {
		return err
	}

	return nil
}

func getClients(db *sql.DB, start, count int) (Clients, error) {
	statement := fmt.Sprintf("SELECT id, name, city, state, web, email, phone FROM clients LIMIT %d OFFSET %d", count, start)
	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	clients := Clients{}

	for rows.Next() {
		var c Client
		if err := rows.Scan(&c.ID, &c.Name, &c.City, &c.State, &c.Web, &c.Email, &c.Phone); err != nil {
			return nil, err
		}
		clients = append(clients, c)
	}

	return clients, nil
}
