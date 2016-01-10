package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"reflect"
)

type Bill struct {
	Year       int           `json:"year" bson:"year"`
	Month      string        `json:"month" bson:"month"`
	Value      int           `json:"value" bson:"value"`
}

type BillRepo struct {
	coll *mgo.Collection
}

func (r *BillRepo) All(year string) ([]bson.M, error) {
	bills := []bson.M{}
	//	yearInt, _ := strconv.Atoi(year)
	pipe := r.coll.Pipe([]bson.M{
		{ "$match": bson.M{"_id": "user1"} },
		{
			"$project": bson.M{
				"bills": bson.M{
					"$filter": bson.M{
						"input": "$bills",
						"as": "bill",
						"cond": bson.M{
							"$eq": []interface{} {
								"$$bill.year",
								2015,
							},
						},
					},
				},
				"_id": 0,
			},
		},
	})

	err := pipe.All(&bills)
	log.Println("Map:")
	log.Println(bills)
	log.Println("Map[0]:")
	log.Println(reflect.TypeOf(bills[0]), bills[0])
	//	log.Println(bills[0])
	for k, v := range bills[0] {
		log.Println("k:", k, ", v:", v)
		//		vtype := reflect.TypeOf(v)
		//		log.Println("vType: ", vtype)
		v1 := reflect.ValueOf(v)
		tmp := make([]interface{}, v1.Len())

		for i:= 0; i < v1.Len(); i++ {
			tmp[i] = v1.Index(i).Interface()
		}
		log.Println(reflect.TypeOf(tmp))
		for _, v2 := range tmp {
			log.Println(reflect.TypeOf(v2), v2)
		}
	}
	tmp := []bson.M{}
	if err != nil {
		return tmp, err
	}

	return tmp, nil
}