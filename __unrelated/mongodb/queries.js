db.bills.find({_id:"user1"}, {"bills": { $elemMatch: {year: 2016}}, _id:0 }).pretty()

// Return all bills for 'user1' in 2016
db.bills.aggregate([ 
	{ 
		$match: { _id:"user1" } 
	},
	{ 	
		$project: {
			bills: { 
				$filter: {
					input: "$bills",
					as: "bill",
					cond: { $eq: ["$$bill.year", 2016]}
				}
			},
			_id: 0
		}
	}
]).pretty()

// Return rows for bill
db.bills.aggregate([
	{
		$match: { _id:"user1" }
	},
	{
		$project: {
			bills: {
				$filter: {
					input: "$bills",
					as: "bill",
					cond: { $eq: ["$$bill._id", ObjectId("5698cc36127d85b94ec337da")]}
				}
			},
			_id: 0,
		}
	}
]).pretty()

db.bills.aggregate([
	{
		$match: { "bills._id":ObjectId("5698cc36127d85b94ec337da") }
	},
	{
		$project: {
			bills: {
				$filter: {
					input: "$bills",
					as: "bill",
					cond: { $eq: ["$$bill._id", ObjectId("5698cc36127d85b94ec337da")]}
				}
			}
		}
	}
]).pretty()


db.bills.aggregate([
	{
		$match: { "bills._id":ObjectId("5698cc36127d85b94ec337da") }
	},
	{
		$project: {
			rows: {
			    $filter: {
			        input: "$bills.rows",
            	    as: "row",
            	    cond: { $gt: ["$$row.value", 100]}
            	}
			}
		}
	}
]).pretty()


db.bills.find({'bills':{$elemMatch:{'rows:'{$elemMatch:{'value'{$in:[1000]}}}}}})