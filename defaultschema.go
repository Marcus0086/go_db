package main

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

type (
	dynamic map[string]interface {
	}
	defjson struct {
		ID      string  `json:"_id"`
		Dynamic dynamic `json:"_data"`
	}
)

func (d *defjson) create(collection, resource string, data json.RawMessage) error {
	db, err := New("./", nil)
	if err != nil {
		fmt.Println("Error while initializing database :", err)
		return err
	}
	m := defjson{
		ID: uuid.New().String(),
	}
	m.Dynamic = dynamic{
		resource: &data,
	}
	if err := db.Write(collection, resource, m); err != nil {
		return err
	}
	return nil
}

func (d *defjson) read(collection, resource string, v interface{}) (interface{}, error) {
	db, err := New("./", nil)
	if err != nil {
		fmt.Println("Error while initializing database :", err)
		return nil, err
	}
	record, err := db.Read(collection, resource, v)
	if err != nil {
		return nil, err
	}
	return record, nil
}

func (d *defjson) readAll(collection string) ([]string, error) {
	db, err := New("./", nil)
	if err != nil {
		fmt.Println("Error while initializing database :", err)
		return nil, err
	}
	records, err := db.ReadAll(collection)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (d *defjson) delete(collection, resource string) error {
	db, err := New("./", nil)
	if err != nil {
		fmt.Println("Error while initializing database :", err)
		return err
	}
	if err := db.Delete(collection, resource); err != nil {
		return err
	}
	return nil
}
