package services

import (
	"time"

	"temux/internal/repository"
)

func ProcessProfits(
	investmentRepo *repository.InvestmentRepository,
) error {

	investments, err :=
		investmentRepo.GetActiveInvestments()

	if err != nil {
		return err
	}

	for _, inv := range investments {

		days :=
			int(
				time.Since(
					inv.StartDate,
				).Hours() / 24,
			)

		if days < 0 {
			days = 0
		}

		profit :=
			inv.Amount *
				(inv.DailyRate / 100) *
				float64(days)

		err =
			investmentRepo.UpdateProfit(
				inv.ID,
				profit,
			)

		if err != nil {
			return err
		}
	}

	return nil
}
