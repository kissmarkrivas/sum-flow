package models

import "gorm.io/gorm"

type Sum struct{
	gorm.Model

	SumFirst 	int `json:"sum_first"`
	SumSecond int `json:"sum_second"`
	SumRes		int	 `json:"sum_res"`
	SumID			uint `json:"sum_id"`
}