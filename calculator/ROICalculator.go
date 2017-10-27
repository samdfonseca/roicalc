package calculator

import (
	"github.com/samdfonseca/roicalc/calculator/model"
)

type ROICalculationResult struct {
	FirstYearROIResult  model.CalculationResult
	SecondYearROIResult model.CalculationResult
	ThirdYearROIResult  model.CalculationResult
}



func calculateROI(previousYear model.CalculationResult,
	assumption model.Assumption) model.CalculationResult{
	var ROIResult model.CalculationResult
	ROIResult.
}

func calculateFirstYear(assumption model.Assumption) model.CalculationResult{
	var result model.CalculationResult
	// Membership
	result.Membership.Total =  assumption.MembershipStartingLevel
	result.Membership.PercentagePurchasing = assumption.PurchasingMembers
	result.Membership.PurchasingMembers = result.Membership.Total *
		result.Membership.PercentagePurchasing
	//Purchases
	result.Purchases.BaselineAnnualSpend = assumption.Aov * assumption.PurchasePerYear;
	result.Purchases.LoyaltyAnnualSpend = result.Purchases.BaselineAnnualSpend *
		assumption.LiftToSpend
	result.Purchases.IncreaseInAnnualSpend = result.Purchases.LoyaltyAnnualSpend -
		result.Purchases.BaselineAnnualSpend
	//Points
	//Earned
	result.Points.Earned.EngagementPoints = assumption
	result.Points.Earned.PurchasePoints
	result.Points.Earned.PointsEarnedInYear
	result.Points.Earned.RedeemablePointsAvailable
}