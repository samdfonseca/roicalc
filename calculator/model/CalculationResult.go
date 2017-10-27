package model

type Membership struct {
	Total                float64
	PercentagePurchasing float64
	PurchasingMembers    float64
}

type Purchases struct {
	BaselineAnnualSpend   float64
	LoyaltyAnnualSpend    float64
	IncreaseInAnnualSpend float64
}

type Earned struct {
	EngagementPoints          float64
	PurchasePoints            float64
	PointsEarnedInYear        float64
	RedeemablePointsAvailable float64
}

type Burned struct {
	PointsRedeemed float64
	PointsExpired float64
}

type Points struct {
	Earned Earned
	Burned Burned
	EOYOutstandingPointsLiablity float64
}

type ProgramCosts struct {
	RedemptionCost       float64
	ProgramLicenseCost   float64
	ProgramMarketingCost float64
}

type ROI struct {
	ProgramBenefit float64
	ProgramCost    float64
	ProgramProfit  float64
}

type CalculationResult struct {
	Membership   Membership
	Purchases    Purchases
	Points       Points
	ProgramCosts ProgramCosts
	ROI          ROI
	ProgramROI   float64
}
