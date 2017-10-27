package calculator

import (
	"github.com/samdfonseca/roicalc/calculator/model"
	"gonum.org/v1/gonum/stat"
	"math"
)

type ROICalculator struct {
	FirstYearROIResult  model.CalculationResult
	SecondYearROIResult model.CalculationResult
	ThirdYearROIResult  model.CalculationResult
}

func calculateYearlyROI(previousYear *model.CalculationResult, assumption model.Assumption, yearIndex int) model.CalculationResult {
	var result model.CalculationResult
	var rate float64
	// Membership
	if previousYear == nil {
		result.Membership.Total = assumption.MembershipStartingLevel
	} else {
		rate = 1+  (assumption.MembershipGrowth / 100 )
		result.Membership.Total = previousYear.Membership.Total * rate
	}

	result.Membership.PercentagePurchasing = assumption.PurchasingMembers /100
	result.Membership.PurchasingMembers = result.Membership.Total * result.Membership.PercentagePurchasing
	//Purchases
	result.Purchases.BaselineAnnualSpend = assumption.Aov * assumption.PurchasePerYear
	result.Purchases.LoyaltyAnnualSpend = result.Purchases.BaselineAnnualSpend *  (1 + assumption.LiftToSpend / 100)
	result.Purchases.IncreaseInAnnualSpend = result.Purchases.LoyaltyAnnualSpend - result.Purchases.BaselineAnnualSpend
	//Points
	//Earned
	result.Points.Earned.EngagementPoints = assumption.EngagementPointsPerMember
	result.Points.Earned.PurchasePoints = result.Purchases.LoyaltyAnnualSpend * assumption.PointPerDollarSpend
	result.Points.Earned.PointsEarnedInYear = result.Points.Earned.EngagementPoints*result.Membership.Total + result.Points.Earned.PurchasePoints*result.Membership.PurchasingMembers
	if previousYear != nil {
		result.Points.Earned.RedeemablePointsAvailable = result.Points.Earned.PointsEarnedInYear + previousYear.Points.EOYOutstandingPointsLiablity
	} else {
		result.Points.Earned.RedeemablePointsAvailable = result.Points.Earned.PointsEarnedInYear
	}
	//Burned
	result.Points.Burned.PointsRedeemed = result.Points.Earned.RedeemablePointsAvailable * assumption.Redemption /100
	result.Points.Burned.PointsExpired = result.Points.Earned.RedeemablePointsAvailable * assumption.PointExpiryRate / 100

	result.Points.EOYOutstandingPointsLiablity = result.Points.Earned.RedeemablePointsAvailable - result.Points.Burned.PointsExpired - result.Points.Burned.PointsRedeemed

	//Program Cost
	result.ProgramCosts.RedemptionCost = result.Points.Burned.PointsRedeemed *
		assumption.DollarValuePerPoint * assumption.DiscountToRewards / 100
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
	result := model.ROICalculationResult{}
	result = append(result, calculateYearlyROI(nil, assumption, 1))
	for i := 1; i < 3; i++ {
		result = append(result, calculateYearlyROI(&result[i-1], assumption, i))
	}
	return result
}


func (r ROICalculator) CalculateROIMatrix(a model.Assumption) [][]float64{
	var results [][]float64
	var result, result2 []float64
	var firstROI, secondROI, thirdROI []float64
	iteration := 1000
	for i := 0;  i < iteration; i ++ {
		 roiResult :=  r.Calculate(model.Randomlize(a))
		 firstROI = append(firstROI, roiResult[0].ProgramROI)
		 secondROI = append(secondROI, roiResult[1].ProgramROI)
		 thirdROI = append(thirdROI, roiResult[2].ProgramROI)
	 }
	 result = append(result, stat.Mean(firstROI, nil))
	result = append(result, stat.Mean(secondROI, nil))
	result = append(result, stat.Mean(thirdROI, nil))
	result2 = append(result2, math.Sqrt(stat.Variance(firstROI, nil)))
	result2 = append(result2, math.Sqrt(stat.Variance(secondROI, nil)))
	result2 = append(result2, math.Sqrt(stat.Variance(thirdROI, nil)))

	results = append(results, result)
	results = append(results, result2)
	return results
	/*
	 for y := 0; y < 2 ; y++ {
		 tSum:= 0.0
		 stdev :=0.0
		 for i := 0; i < iteration; i++ {
		 	tSum += roiResults[i][y].ProgramROI;
		 	stdev +=
		 }
		 result = append(result, tSum/float64(iteration))
		result2 = append(result, )

	 }*/
}