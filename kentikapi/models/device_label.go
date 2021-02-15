package models

import "time"

type DeviceLabel struct {
	// read-write
    Name string
    Color string

    // read-only
    Devices []DeviceItem
    ID ID
    UserID *ID
    CompanyID ID
    CreatedDate time.Time
    UpdatedDate time.Time
}

type DeviceItem struct {
    ID ID
    DeviceName string
    DeviceSubtype string
    DeviceType *string
}