package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/csmartin1024/goopy/helper"
	"github.com/csmartin1024/goopy/models"
	"go.mongodb.org/mongo-driver/bson"
)

func ListGoops(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// we created Goop array
	var goops = []models.Goop{}

	//Connection mongoDB with helper class
	// collection := helper.Collections.Goops

	// bson.M{},  we passed empty filter. So we want to get all data.
	cur, err := helper.Collections.Goops.Find(helper.CTX, bson.M{})

	if err != nil {
		helper.GetError(err, w)
		return
	}

	// Close the cursor once finished
	/*A defer statement defers the execution of a function until the surrounding function returns.
	simply, run cur.Close() process but after cur.Next() finished.*/
	defer cur.Close(helper.CTX)

	for cur.Next(helper.CTX) {

		// create a value into which the single document can be decoded
		var goop models.Goop
		// & character returns the memory address of the following variable.
		err := cur.Decode(&goop) // decode similar to deserialize process.
		if err != nil {
			log.Fatal(err)
		}

		// add item our array
		goops = append(goops, goop)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(goops) // encode similar to serialize process.
}
