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

	userAccount, err := originClient.GetUserAccount(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", userAccount.Viewer.Digital.Accounts[0].AccountID)

	startDate := time.Now().AddDate(0, 0, -1)

	filter := originenergy.AccountUsageFilter{
		StartDate: startDate,
		Type:      constants.EmbeddedElectricity,
		TimeUnit:  constants.Hourly,
	}

	accountUsage, err := originClient.GetAccountUsage(ctx, filter)
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

	fmt.Printf("Last Updated: %s\n", lastNotZeroDataPoint.EndDate.Format("02-01-2006 03:04PM"))
}
