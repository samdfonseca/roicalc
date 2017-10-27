package model

type Assumption struct {
	Aov                       float64
	PurchasePerYear           float64
	PointPerDollarSpend       float64
	DollarValuePerPoint       float64
	DiscountToRewards         float64
	AnnualPerMemberMarketCost float64
	MembershipStartingLevel   float64
	MembershipGrowth          float64
	PurchasingMembers         float64
	EngagementPointsPerMember float64
	LiftToSpend               float64
	Redemption                float64
	PointExpiryRate           float64
	ProgramCostsFirstYear     float64
	ProgramCostsSecondYear    float64
	ProgramCostThirdYear      float64
}
