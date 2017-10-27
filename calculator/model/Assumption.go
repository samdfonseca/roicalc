package model

type Range struct {
	Distribution string  `json:"distribution"`
	High         float64 `json:"high"`
	Low          float64 `json:"low"`
}

type Assumption struct {
	Aov                       float64   `json:"aov"`
	PurchasePerYear           float64   `json:"purchase_per_year"`
	PointPerDollarSpend       float64   `json:"points_per_spend"`
	DollarValuePerPoint       float64   `json:"dollar_point_value"`
	DiscountToRewards         float64   `json:"discount_to_rewards"`
	AnnualPerMemberMarketCost float64   `json:"member_marketing_cost"`
	MembershipStartingLevel   float64   `json:"starting_members"`
	MembershipGrowth          float64   `json:"membership_growth"`
	PurchasingMembers         float64   `json:"purchasing_members"`
	EngagementPointsPerMember float64   `json:"engagement_points"`
	LiftToSpend               float64   `json:"lift_to_spend"`
	LiftToSpendRange          Range     `json:"lift_to_spend_range"`
	Redemption                float64   `json:"redemption_rate"`
	RedemptionRateRange       Range     `json:"redemption_rate_range"`
	PointExpiryRate           float64   `json:"point_expiry_rate"`
	PointExpiryRateRange      Range     `json:"point_expiry_rate_range"`
	ProgramCosts              []float64 `json:"program_costs"`
}
