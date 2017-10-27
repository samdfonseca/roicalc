package model

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
	Redemption                float64   `json:"redemption_rate"`
	PointExpiryRate           float64   `json:"point_expiry_rate"`
	ProgramCosts              []float64 `json:"program_costs"`
}
