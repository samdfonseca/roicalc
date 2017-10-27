package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type ParamsBody struct {
	AOV                 int     `json:"aov"`
	PurchasePerYear     float64 `json:"purchase_per_year"`
	PointsPerSpend      int     `json:"points_per_spend"`
	DollarPointsValue   float64 `json:"dollar_point_value"`
	DiscountToRewards   int     `json:"discount_to_rewards"`
	MemberMarketingCost float64 `json:"member_marketing_cost"`
	StartingMembers     int     `json:"starting_members"`
	MembershipGrowth    int     `json:"membership_growth"`
	PurchasingMembers   int     `json:"purchasing_members"`
	EngagementPoints    int     `json:"engagement_points"`
	LiftToSpend         int     `json:"lift_to_spend"`
	RedemptionRate      int     `json:"redemption_rate"`
	PointExpiryRate     int     `json:"point_expiry_rate"`
	ProgramCosts        []int   `json:"program_costs"`
}

func baseHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"hello\": \"world\"}"))
}

func configHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error reading request body"))
		return
	}
	var paramsBody ParamsBody
	err = json.Unmarshal(body, &paramsBody)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		errMsg := fmt.Sprintf("Unprocessable entity: %s", err.Error())
		w.Write([]byte(errMsg))
		return
	}

	spew.Dump(paramsBody)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", baseHandler)
	r.HandleFunc("/params", configHandler).Methods("POST")

	handler := cors.Default().Handler(r)
	fmt.Println("Serving at :3000")
	http.ListenAndServe(":3000", handler)
}
