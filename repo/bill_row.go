package repo
import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"log"
	"reflect"
)

type BillRow struct {
	Id            bson.ObjectId `json:"id" bson:"_id"`
	PublicUtility PublicUtility `json:"publicUtility" bson:"publicUtility"`
	Value         int           `json:"value" bson:"value"`
}

type BillRowsCollection struct {
	Data []BillRow `json:"data"`
}

type BillRowResource struct {
	Data BillRow `json:"data"`
}

type BillRowRepo struct {
	Coll *mgo.Collection
}

func (r *BillRowRepo) All(id string) (BillRowsCollection, error) {
	rowsC := BillRowsCollection{[]BillRow{}}
	bsonMap := []bson.M{}

	pipe := r.Coll.Pipe([]bson.M{
		{"$match": bson.M{"_id": "user1"}},
		{
			"$project": bson.M{
				"bills": bson.M{
					"$filter": bson.M{
						"input": "$bills",
						"as":    "bill",
						"cond": bson.M{
							"$eq": []interface{}{
								"$$bill._id",
								bson.ObjectIdHex(id),
							},
						},
					},
				},
				"_id": 0,
			},
		},
	})
	if err := pipe.All(&bsonMap); err != nil {
		return rowsC, err
	}
	log.Println("bsonMap:", bsonMap[0])
	bsonToRows(bsonMap[0]["bills"])

	return rowsC, nil
}

func bsonToRows0(bsonM interface{}) BillRowsCollection {
	bsonMVal := reflect.ValueOf(bsonM)

	t := reflect.ValueOf(bsonM).Index(0).Interface()
	log.Println("t", t)
	billVals := make([]interface{}, bsonMVal.Len())
	for i := 0; i < bsonMVal.Len(); i++ {
		billVals[i] = bsonMVal.Index(i).Interface()
	}
	log.Println("billVals:", billVals)
	log.Println("billVals[0]", billVals[0])

//	billM := make(map[string]interface{})
//	for _, k := range reflect.ValueOf(billVals[0]).MapKeys() {
//		log.Println("k", k)
//		billM[k.String()] = bsonMVal.MapIndex(k).Interface()
//	}
//	log.Println(billM["rows"])
//	billVals := reflect.ValueOf(billM)
//	var rowsM = make(map[string]interface{})
//	for _, k := range billVals.MapKeys() {
//		rowsM[k.String()] = billVals.MapIndex(k).Interface()
//	}
//	log.Println("rows:", rowsM["rows"])

	var billM = make(map[string]interface{})
	for _, k := range reflect.ValueOf(billVals[0]).MapKeys() {
		billM[k.String()] = reflect.ValueOf(billVals[0]).MapIndex(k).Interface()
	}
	log.Println("rows", reflect.TypeOf(billM["rows"]), billM["rows"])



	result := BillRowsCollection{make([]BillRow, len(billVals))}
	for _, billItem := range billVals {
		log.Println("billItem:", billItem)
		rowsBsonMVal := reflect.ValueOf(billItem)
		var billM = make(map[string]interface{})
		for _, k := range rowsBsonMVal.MapKeys() {
			billM[k.String()] = rowsBsonMVal.MapIndex(k).Interface()
		}
		log.Println("rows", reflect.TypeOf(billM["rows"]), billM["rows"])

//		for i, k := range reflect.ValueOf(billM["rows"]).MapKeys() {
//			log.Println(i, k)
//		}

	}

	return result
}

func bsonToRows(bsonM interface{}) BillRowsCollection {
//	bsonMVal := reflect.ValueOf(bsonM)

	billVals := reflect.ValueOf(bsonM).Index(0).Interface()
	log.Println("billVals", billVals)
//	billVals := make([]interface{}, bsonMVal.Len())
//	for i := 0; i < bsonMVal.Len(); i++ {
//		billVals[i] = bsonMVal.Index(i).Interface()
//	}
//	log.Println("billVals:", billVals)
//	log.Println("billVals[0]", billVals[0])

	//	billM := make(map[string]interface{})
	//	for _, k := range reflect.ValueOf(billVals[0]).MapKeys() {
	//		log.Println("k", k)
	//		billM[k.String()] = bsonMVal.MapIndex(k).Interface()
	//	}
	//	log.Println(billM["rows"])
	//	billVals := reflect.ValueOf(billM)
	//	var rowsM = make(map[string]interface{})
	//	for _, k := range billVals.MapKeys() {
	//		rowsM[k.String()] = billVals.MapIndex(k).Interface()
	//	}
	//	log.Println("rows:", rowsM["rows"])

	var billM = make(map[string]interface{})
	for _, k := range reflect.ValueOf(billVals).MapKeys() {
		billM[k.String()] = reflect.ValueOf(billVals).MapIndex(k).Interface()
	}
	rows := reflect.ValueOf(billM["rows"])
	log.Println("rows", reflect.TypeOf(rows), rows)

	for i := 0; i < rows.Len(); i++ {
		log.Println(i, reflect.TypeOf(rows), reflect.TypeOf(rows.Index(i).Interface()))

//		bsonMVal := rows.Index(i)
//
//		billsM := make([]interface{}, bsonMVal.Len())
//		for i := 0; i < bsonMVal.Len(); i++ {
//			billsM[i] = bsonMVal.Index(i).Interface()
//		}
//		for j := 0; j < rows.Index(i).Len(); j++ {
//			log.Println(rows.Index(i).Index(j))
//		}
//		v := rows.Index(i).Interface()
//		m := v.(map[string]interface{})
//		log.Println("m:", m)
//		for _, k := range rows.MapKeys() {
//			log.Println(k.String())
//		}
//		for v := range rows.Index(i).MapKeys() {
//			log.Println(v)
//		}
	}

//	for _, k := range rows {
//		log.Println(k)
//	}



	result := BillRowsCollection{make([]BillRow, 0)}
//	for _, billItem := range billVals {
//		log.Println("billItem:", billItem)
//		rowsBsonMVal := reflect.ValueOf(billItem)
//		var billM = make(map[string]interface{})
//		for _, k := range rowsBsonMVal.MapKeys() {
//			billM[k.String()] = rowsBsonMVal.MapIndex(k).Interface()
//		}
//		log.Println("rows", reflect.TypeOf(billM["rows"]), billM["rows"])
//
//		//		for i, k := range reflect.ValueOf(billM["rows"]).MapKeys() {
//		//			log.Println(i, k)
//		//		}
//
//	}

	return result
}