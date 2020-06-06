package routes

import (
	"encoding/json"
	"net/http"

	"github.com/csmartin1024/goopy/helper"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteGoop(w http.ResponseWriter, r *http.Request) {
	// Set header
	w.Header().Set("Content-Type", "application/json")

	// get params
	var params = mux.Vars(r)

	// string to primitve.ObjectID
	id, err := primitive.ObjectIDFromHex(params["id"])

	// collection := helper.GetDB()

	// prepare filter.
	filter := bson.M{"_id": id}

	deleteResult, err := helper.Collections.Goops.DeleteOne(helper.CTX, filter)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(deleteResult)
}
