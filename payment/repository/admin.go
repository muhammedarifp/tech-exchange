package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/muhammedarifp/tech-exchange/payments/commonHelp/response"
	"github.com/muhammedarifp/tech-exchange/payments/repository/interfaces"
	"gorm.io/gorm"
)

type adminPaymentDb struct {
	db *gorm.DB
}

func NewAdminPaymentDb(db *gorm.DB) interfaces.AdminPaymentRepo {
	return &adminPaymentDb{
		db: db,
	}
}

func (d *adminPaymentDb) AddPlan(ctx context.Context, plan map[string]interface{}) (response.Plans, error) {
	qury := `INSERT INTO plans(plan_id,name,description,interval,period,amount,is_active) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING *`
	item := plan["item"].(map[string]interface{})
	planid := plan["id"].(string)
	name := item["name"].(string)
	disc := item["description"].(string)
	intervel := int(plan["interval"].(float64))
	period := plan["period"].(string)
	amount := item["amount"].(float64)

	var data response.Plans
	err := d.db.Raw(qury, planid, name, disc, intervel, period, amount, true).Scan(&data).Error
	if err != nil {
		return response.Plans{}, err
	}

	fmt.Println(data)

	return data, nil
}

func (d *adminPaymentDb) RemovePlan(ctx context.Context, planid string) (response.Plans, error) {
	var fplan response.Plans

	tx := d.db.Begin()

	fetchQury := `SELECT * FROM plans WHERE plan_id = ?`
	if err := tx.Raw(fetchQury, planid).Scan(&fplan).Error; err != nil {
		tx.Rollback()
		return response.Plans{}, err
	}

	if !fplan.IsActive {
		tx.Rollback()
		return response.Plans{}, errors.New("this plan already deactivated")
	}

	qury := `UPDATE plans SET is_active = false WHERE plan_id = $1 RETURNING *`
	var splan response.Plans
	if err := tx.Raw(qury, planid).Scan(&splan).Error; err != nil {
		tx.Rollback()
		return response.Plans{}, err
	}

	if err := tx.Commit().Error; err != nil {
		return response.Plans{}, err
	}

	return splan, nil
}
