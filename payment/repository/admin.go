package repository

import (
	"context"
	"fmt"

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

func (d *adminPaymentDb) AddPlan(ctx context.Context, plan map[string]interface{}) {
	qury := `INSERT INTO plans(plan_id,name,description,interval,period,amount,is_active) VALUES ($1,$2,$3,$4,$5,$6,$7)`
	item := plan["item"].(map[string]interface{})
	var a interface{}
	planid := plan["id"].(string)
	name := item["name"].(string)
	disc := item["description"].(string)
	intervel := int(plan["interval"].(float64))
	period := plan["period"].(string)
	amount := item["amount"].(float64)

	fmt.Printf("Val : %d, type: %T", intervel, intervel)
	err := d.db.Raw(qury, planid, name, disc, 1, period, amount, true).Scan(&a).Error
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Okk")
	fmt.Println(a)
}
func (d *adminPaymentDb) RemovePlan(ctx context.Context) {}
