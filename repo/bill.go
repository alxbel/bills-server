package repo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
//	"log"
	"reflect"
	"strconv"
	"log"
)

// Bill

type Bill struct {
	Id    bson.ObjectId `json:"id" bson:"_id"`
	Year  int    `json:"year"`
	Month string `json:"month"`
	Value int    `json:"value"`
	Rows []interface{}	`json:"rows"`
}

type BillsCollection struct {
	Data []Bill `json:"data"`
}

type BillResource struct {
	Data Bill `json:"data"`
}

type BillRepo struct {
	Coll *mgo.Collection
}

//func (r *BillRepo) All(year string) ([]Bill, error) {
//	bsonMap := []bson.M{}
//	yearInt, _ := strconv.Atoi(year)
//
//	pipe := r.Coll.Pipe([]bson.M{
//		{"$match": bson.M{"_id": "user1"}},
//		{
//			"$project": bson.M{
//				"bills": bson.M{
//					"$filter": bson.M{
//						"input": "$bills",
//						"as":    "bill",
//						"cond": bson.M{
//							"$eq": []interface{}{
//								"$$bill.year",
//								yearInt,
//							},
//						},
//					},
//				},
//				"_id": 0,
//			},
//		},
//	})
//
//	if err := pipe.All(&bsonMap); err != nil {
//		return nil, err
//	}
//
//	bills := bson2Bills(bsonMap[0]["bills"])
////	billsMap := bsonMap[0]["bills"]
////	log.Println("billsMap:", billsMap, ".")
////	for k, v := range make([]interface{}, reflect.ValueOf(billsMap).Len()) {
////		log.Println(k, v)
////	}
//
////	for k, v := range bsonMap[0] {
////		log.Println("key:", k, ", val:", v)
////		v1 := reflect.ValueOf(v)
////		tmp := make([]interface{}, v1.Len())
////
////		for i := 0; i < v1.Len(); i++ {
////			tmp[i] = v1.Index(i).Interface()
////		}
////
////		bills = make([]Bill, len(tmp))
////		for i, v2 := range tmp {
////			v3 := reflect.ValueOf(v2)
////			var billM = make(map[string]interface{})
////			for _, k := range v3.MapKeys() {
////				billM[k.String()] = v3.MapIndex(k).Interface()
////			}
////			bills[i].Year = billM["year"].(int)
////			bills[i].Month = billM["month"].(string)
////			bills[i].Value = billM["value"].(int)
////		}
////	}
//
//	return bills, nil
//}
//
//func bson2Bills(bsonM interface{}) []Bill {
//	log.Println("bson2Bills:", bsonM, ".")
//	bsonMVal := reflect.ValueOf(bsonM)
//	log.Println("reflect:", reflect.TypeOf(bsonMVal), bsonMVal)
//
//	billsM := make([]interface{}, bsonMVal.Len())
//	for i := 0; i < bsonMVal.Len(); i++ {
//		billsM[i] = bsonMVal.Index(i).Interface()
//	}
//	log.Println("billsM", billsM)
//
//	bills := make([]Bill, len(billsM))
//	for i, billBsonM := range billsM {
//		billBsonMVal := reflect.ValueOf(billBsonM)
//		var billM = make(map[string]interface{})
//		for _, k := range billBsonMVal.MapKeys() {
//			billM[k.String()] = billBsonMVal.MapIndex(k).Interface()
//		}
//		bills[i].Id = billM["_id"].(bson.ObjectId)
//		bills[i].Year = billM["year"].(int)
//		bills[i].Month = billM["month"].(string)
//		bills[i].Value = billM["value"].(int)
//	}
//
//	return bills
//}

////////////////////////////////////////////////////////////////////////////////////////////////

func (r *BillRepo) All(year string) (BillsCollection, error) {
	billsC := BillsCollection{[]Bill{}}
	bsonMap := []bson.M{}
	yearInt, _ := strconv.Atoi(year)

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
								"$$bill.year",
								yearInt,
							},
						},
					},
				},
				"_id": 0,
			},
		},
	})


	if err := pipe.All(&bsonMap); err != nil {
		return billsC, err
	}
	billsC = bson2BillsC(bsonMap[0]["bills"])

	return billsC, nil
}

func bson2BillsC(bsonM interface{}) BillsCollection {
	log.Println("bsonM:", bsonM)
	bsonMVal := reflect.ValueOf(bsonM)

	billsM := make([]interface{}, bsonMVal.Len())
	for i := 0; i < bsonMVal.Len(); i++ {
		billsM[i] = bsonMVal.Index(i).Interface()
	}

	result := BillsCollection{make([]Bill, len(billsM))}
	for i, billBsonM := range billsM {
		billBsonMVal := reflect.ValueOf(billBsonM)
		var billM = make(map[string]interface{})
		for _, k := range billBsonMVal.MapKeys() {
			billM[k.String()] = billBsonMVal.MapIndex(k).Interface()
		}
		result.Data[i].Id = billM["_id"].(bson.ObjectId)
		result.Data[i].Year = billM["year"].(int)
		result.Data[i].Month = billM["month"].(string)
		result.Data[i].Value = billM["value"].(int)
//		result.Data[i].PublicUtility = billM["rows"].([]interface{})
		log.Println("rows:", reflect.TypeOf(billM["rows"]), billM["rows"])
	}

	return result
}