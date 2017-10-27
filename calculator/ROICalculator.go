package calculator

import (
	"github.com/samdfonseca/roicalc/calculator/model"
	"go/types"
)

type ROICalculator struct {
	FirstYearROIResult  model.CalculationResult
	SecondYearROIResult model.CalculationResult
	ThirdYearROIResult  model.CalculationResult
}

func calculateYearlyROI(previousYear model.CalculationResult, assumption model.Assumption, yearIndex int) model.CalculationResult {
	var result model.CalculationResult
	// Membership
	result.Membership.Total = assumption.MembershipStartingLevel
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
	result.Points.Earned.EngagementPoints = assumption.EngagementPointsPerMember;
	result.Points.Earned.PurchasePoints = result.Points.Earned.EngagementPoints * assumption.PointPerDollarSpend
	result.Points.Earned.PointsEarnedInYear = result.Points.Earned.EngagementPoints*result.Membership.Total + result.Points.Earned.PurchasePoints*result.Membership.PurchasingMembers
	if previousYear != nil {
		result.Points.Earned.RedeemablePointsAvailable = previousYear.Points.Earned.PointsEarnedInYear + result.Points.Earned.RedeemablePointsAvailable
	} else {
		result.Points.Earned.RedeemablePointsAvailable = result.Points.Earned.PointsEarnedInYear
	}
	//Burned
	result.Points.Burned.PointsRedeemed = result.Points.Earned.RedeemablePointsAvailable * assumption.Redemption
	result.Points.Burned.PointsExpired = result.Points.Earned.PointsEarnedInYear * assumption.PointExpiryRate

	result.Points.EOYOutstandingPointsLiablity = result.Points.Earned.RedeemablePointsAvailable - result.Points.Burned.PointsExpired - result.Points.Burned.PointsRedeemed

	//Program Cost
	result.ProgramCosts.RedemptionCost = result.Points.Burned.PointsRedeemed *
		assumption.DollarValuePerPoint * assumption.DiscountToRewards
	result.ProgramCosts.ProgramLicenseCost = assumption.ProgramCosts[yearIndex]
	result.ProgramCosts.ProgramMarketingCost = result.Membership.Total * assumption.AnnualPerMemberMarketCost

	//ROI
	result.ROI.ProgramBenefit = result.Membership.PurchasingMembers * result.Purchases.IncreaseInAnnualSpend
	result.ROI.ProgramCost = result.ProgramCosts.RedemptionCost + result.ProgramCosts.ProgramMarketingCost + result.ProgramCosts.ProgramLicenseCost
	result.ROI.ProgramProfit = result.ROI.ProgramBenefit - result.ROI.ProgramCost

	result.ProgramROI = result.ROI.ProgramProfit / result.ROI.ProgramCost

	return result
}

func (r ROICalculator) Calculate(assumption model.Assumption) model.ROICalculationResult {
	var result model.ROICalculationResult
	result[0] = calculateYearlyROI(nil, assumption, 1);
	for i := 1; i < 3; i++ {
		result[i] = calculateYearlyROI(result[i-1], assumption, i);
	}
	return result
}
