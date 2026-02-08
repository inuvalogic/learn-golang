package models

import "time"

type TimeRange struct {
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

type BestSeller struct {
	ProductName string `json:"nama"`
	Total       int    `json:"qty_terjual"`
}

type TodayReport struct {
	TotalRevenue     int          `json:"total_revenue"`
	TotalTransaction int          `json:"total_transaksi"`
	BestSeller       *BestSeller  `json:"produk_terlaris"`
	// for debug time on supabase
	// StartDate		time.Time	`json:"start_date"`
	// EndDate			time.Time	`json:"end_date"`
}