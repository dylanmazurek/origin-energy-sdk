package constants

import "time"

const (
	API_BASE_URL = "https://api.rx.originenergy.com.au"
	GRAPHQL_PATH = "/v1/gateway/graphql"
)

const (
	AUTH_CLIENT_ID    = "yOHRT97N3yH85jzTDlqN2A7Cf2D0cmQe"
	AUTH_DOMAIN       = "id.originenergy.com.au"
	AUTH_TENANT       = "originenergy"
	AUTH_BASE_URL     = "https://id.originenergy.com.au/"
	AUTH_CALLBACK_URL = "https://www.originenergy.com.au/auth/callback"
	AUTH_AUDIENCE     = "https://digitalapi"
)

type ServiceType string

const (
	Electricity         ServiceType = "ELECTRICITY"
	EmbeddedElectricity ServiceType = "EMBEDDED_ELECTRICITY"
	Gas                 ServiceType = "GAS"
	HotWater            ServiceType = "HOT_WATER"
	EmbeddedHotWater    ServiceType = "EMBEDDED_HOT_WATER"
)

type UsageTimeUnit string

const (
	Hourly  UsageTimeUnit = "HOURLY"
	Daily   UsageTimeUnit = "DAILY"
	Monthly UsageTimeUnit = "MONTHLY"
	Yearly  UsageTimeUnit = "YEARLY"
)

type DateTime string

func (d *DateTime) FromTime(t time.Time) {
	timeString := t.Format(time.RFC3339)
	*d = DateTime(timeString)
}
