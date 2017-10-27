package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/samdfonseca/roicalc/calculator"
	"github.com/samdfonseca/roicalc/calculator/model"
)

type ROIResponse struct {
	RoiMean              []float64 `json:"roi_mean"`
	RoiStandardDeviation []float64 `json:"roi_standard_deviation"`
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
	var paramsBody model.Assumption
	err = json.Unmarshal(body, &paramsBody)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		errMsg := fmt.Sprintf("Unprocessable entity: %s", err.Error())
		w.Write([]byte(errMsg))
		return
	}

	var calc calculator.ROICalculator
	res := calc.Calculate(paramsBody)
	spew.Dump(res)
	var roi_mean []float64
	roi_mean = append(roi_mean, res[0].ProgramROI)
	roi_mean = append(roi_mean, res[1].ProgramROI)
	roi_mean = append(roi_mean, res[2].ProgramROI)
	roi := ROIResponse{
		RoiMean: roi_mean,
	}
	resp, _ := json.Marshal(roi)
	w.Write(resp)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", baseHandler)
	r.HandleFunc("/params", configHandler).Methods("POST")

	handler := cors.Default().Handler(r)
	fmt.Println("Serving at :3000")
	http.ListenAndServe(":3000", handler)
}
