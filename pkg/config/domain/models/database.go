package models

import customErr "github.com/eli95/template/pkg/custom-errors"

type Database struct {
	DriverName     string `json:"driver_name"`
	DataSourceName string `json:"data_source_name"`
}

func (d *Database) Validate() error {
	if d.DriverName == "" {
		return customErr.ErrDatabaseNameRequired
	}
	if d.DataSourceName == "" {
		return customErr.ErrDatabaseURLRequired
	}
	return nil
}
