package repo

import "gopkg.in/mgo.v2"

type PublicUtility struct {
	Name string `json:"name"`
}

type PublicUtilitiesCollection struct {
	Data []PublicUtility `json:"data"`
}

type PublicUtilityRepo struct {
	Coll *mgo.Collection
}

func (r *PublicUtilityRepo) All() (PublicUtilitiesCollection, error) {
	result := PublicUtilitiesCollection{[]PublicUtility{}}
	err := r.Coll.Find(nil).All(&result.Data)
	if err != nil {
		return result, err
	}

	return result, nil
}
