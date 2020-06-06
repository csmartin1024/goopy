package routes

import (
	"encoding/json"
	"net/http"

	"github.com/csmartin1024/goopy/helper"
	"github.com/csmartin1024/goopy/models"
)

func CreateGoop(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var goop models.Goop

	// we decode our body request params
	_ = json.NewDecoder(r.Body).Decode(&goop)

	// connect db
	// collection := helper.GetDB()

	// insert our goop model.
	result, err := helper.Collections.Goops.InsertOne(helper.CTX, goop)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}
