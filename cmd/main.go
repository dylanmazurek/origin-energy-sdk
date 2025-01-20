package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/dylanmazurek/origin-energy-sdk/pkg/originenergy"
	"github.com/dylanmazurek/origin-energy-sdk/pkg/originenergy/constants"

	"github.com/markkurossi/tabulate"
)

func main() {
	ctx := context.Background()
	originClient, err := originenergy.New(ctx)
	if err != nil {
		panic(err)
	}

	printAccountDetails(originClient)
	printUsage(originClient)
}

func printAccountDetails(client *originenergy.Client) {
	ctx := context.Background()

	userAccount, err := client.GetUserAccount(ctx)
	if err != nil {
		panic(err)
	}

	tab := tabulate.New(tabulate.Unicode)
	tab.Header("Account Type")

	row := tab.Row()
	row.Column(userAccount.Viewer.Digital.User.CustomerType)

	fmt.Println(tab.String())

	tab2 := tabulate.New(tabulate.Unicode)
	tab2.Header("Backend Service ID")
	tab2.Header("Status")
	tab2.Header("Type")

	for _, service := range userAccount.Viewer.Digital.Services {
		row := tab2.Row()
		row.Column(service.BackendServiceID)
		row.Column(service.Status)
		row.Column(service.Type)
	}

	fmt.Println(tab2.String())
}

func printUsage(client *originenergy.Client) {
	ctx := context.Background()

	agreementId, agreementIdSet := os.LookupEnv("AGREEMENT_ID")
	if !agreementIdSet || agreementId == "" {
		panic(originenergy.ErrAgreementIDNotSet)
	}

	startDate := time.Now().AddDate(0, 0, -1)

	filter := originenergy.AccountUsageFilter{
		AgreementID: agreementId,
		StartDate:   startDate,
		EndDate:     time.Now(),
		Type:        constants.EmbeddedElectricity,
		TimeUnit:    constants.Hourly,
	}

	accountUsage, err := client.GetAccountUsage(ctx, filter)
	if err != nil {
		panic(err)
	}

	tab := tabulate.New(tabulate.Unicode)
	tab.Header("Date")
	tab.Header("Time")
	tab.Header("Consumed Wh")
	tab.Header("UsageCost")
	tab.Header("SupplyPrice")

	for _, dataPoint := range accountUsage.Viewer.Kraken.Service.EmbeddedElectricityService.Usage.DataPoints {
		row := tab.Row()
		row.Column(dataPoint.StartDate.Format("02-01-2006"))

		row.Column(fmt.Sprintf("%s - %s", dataPoint.StartDate.Format("03:04PM"), dataPoint.EndDate.Format("03:04PM")))

		consumedKiloWattHours := float64(dataPoint.TotalConsumedEnergy) / 1000
		row.Column(fmt.Sprintf("%.2f kWh", consumedKiloWattHours))

		usagePrice := os.Getenv("USAGE_PRICE")
		usagePriceFloat, _ := strconv.ParseFloat(usagePrice, 32)
		row.Column(fmt.Sprintf("$%.2f", consumedKiloWattHours*usagePriceFloat))

		daysBetween := dataPoint.EndDate.Sub(dataPoint.StartDate).Hours() / 24
		daySupplyPrice := os.Getenv("DAY_SUPPLY_PRICE")
		daySupplyPriceFloat, _ := strconv.ParseFloat(daySupplyPrice, 32)
		supplyPrice := daySupplyPriceFloat * float64(daysBetween)
		row.Column(fmt.Sprintf("$%.2f", supplyPrice))
	}

	fmt.Println(tab.String())

	embedElectricityService := accountUsage.Viewer.Kraken.Service.EmbeddedElectricityService

	hasMissingCosts := embedElectricityService.Usage.HasMissingCosts

	dataPoints := embedElectricityService.Usage.DataPoints
	lastNotZeroDataPoint := dataPoints[len(dataPoints)-1]
	if hasMissingCosts {
		for i := len(dataPoints) - 1; i >= 0; i-- {
			if dataPoints[i].TotalConsumedEnergy != 0 {
				lastNotZeroDataPoint = dataPoints[i]
				break
			}
		}
	}

	fmt.Printf("last updated: %s\n", lastNotZeroDataPoint.EndDate.Format("02-01-2006 03:04PM"))
}
