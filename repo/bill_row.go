package repo

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
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

type bill struct {
	Rows []BillRow `json:"rows"`
}

type billsCol struct {
	Data []bill `json:"data"`
}

func (r *BillRowRepo) All(id string) (BillRowsCollection, error) {
	rowsC := BillRowsCollection{[]BillRow{}}
	billsC := []billsCol{}

	pipe := r.Coll.Pipe([]bson.M{
		{"$match": bson.M{"_id": "user1"}},
		{
			"$project": bson.M{
				"data": bson.M{
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
	if err := pipe.All(&billsC); err != nil {
		return rowsC, err
	}
	rowsC.Data = billsC[0].Data[0].Rows

	return rowsC, nil
}

func (r *BillRowRepo) Delete(bid string, rid string) error {
	query := bson.M{
		"bills._id": bson.ObjectIdHex(bid),
	}

	delete := bson.M{
		"$pull": bson.M{
			"bills.$.rows": bson.M{
				"_id": bson.ObjectIdHex(rid),
			},
		},
	}

	if err := r.Coll.Update(query, delete); err != nil {
		return err
	}

	return nil
}