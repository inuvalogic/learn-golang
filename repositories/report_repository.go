package repositories

import (
	"database/sql"
	"kasir-api/models"
)

type ReportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (repo *ReportRepository) GetReport(timerange models.TimeRange) (*models.TodayReport, error) {
	var totalRevenue int
	var totalTransaction int

	err := repo.db.QueryRow(
		`SELECT COALESCE(SUM(total_amount), 0)
		 FROM transactions
		 WHERE created_at >= $1 AND created_at < $2`,
		timerange.StartDate,
		timerange.EndDate,
	).Scan(&totalRevenue)
	if err != nil {
		return nil, err
	}

	err = repo.db.QueryRow(
		`SELECT COUNT(*)
		 FROM transactions
		 WHERE created_at >= $1 AND created_at < $2`,
		timerange.StartDate,
		timerange.EndDate,
	).Scan(&totalTransaction)
	if err != nil {
		return nil, err
	}

	var bestSeller *models.BestSeller

	row := repo.db.QueryRow(
		`
		SELECT
			p.name,
			SUM(td.quantity) AS total_quantity
		FROM transaction_details td
		JOIN products p ON td.product_id = p.id
		JOIN transactions t ON td.transaction_id = t.id
		WHERE t.created_at >= $1 AND t.created_at < $2
		GROUP BY p.id, p.name
		ORDER BY total_quantity DESC
		LIMIT 1
		`,
		timerange.StartDate,
		timerange.EndDate,
	)

	var productName string
	var totalQuantity int

	err = row.Scan(&productName, &totalQuantity)
	if err != nil {
		if err == sql.ErrNoRows {
			bestSeller = nil
		} else {
			return nil, err
		}
	} else {
		bestSeller = &models.BestSeller{
			ProductName: productName,
			Total:       totalQuantity,
		}
	}

	return &models.TodayReport{
		TotalRevenue:     totalRevenue,
		TotalTransaction: totalTransaction,
		BestSeller:       bestSeller,
		// for debug time on supabase
		// StartDate:		timerange.StartDate,
		// EndDate:			timerange.EndDate,
	}, nil
}

