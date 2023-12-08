package interfaces

type UserPaymentUsecase interface {
	FetchAllPlans()
	CreateSubscription()
	CancelSubscription()
	ChangePlan()
}
