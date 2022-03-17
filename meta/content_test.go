package meta

import (
	"encoding/json"
	"testing"
)

func TestTables(t *testing.T) {
	dataString := `
	{
		"entities": [
			{
				"name": "NewEntity3",
				"uuid": "2fac8cfc-fc71-446a-a9cf-36056a63ba78",
				"columns": [
					{
						"name": "id",
						"type": "ID",
						"uuid": "4886a2b0-82ff-4ed6-bd3b-51b97a661d34",
						"primary": true
					},
					{
						"name": "newColumn1",
						"type": "String",
						"uuid": "69135ab9-49c7-424b-855a-e6a3767be920"
					},
					{
						"name": "newColumn2",
						"type": "String",
						"uuid": "0b32cd62-1d26-401b-8d24-3a7b6d260d2b"
					}
				],
				"entityType": "Normal"
			},
			{
				"name": "NewEntity4",
				"uuid": "3e9ae743-de18-4b0c-a77e-3726be4049a8",
				"columns": [
					{
						"name": "id",
						"type": "ID",
						"uuid": "6758ae89-1e2c-462d-907c-a54baf6cf6fd",
						"primary": true
					},
					{
						"name": "newColumn1",
						"type": "String",
						"uuid": "e1afb0c4-5eee-40f3-8c34-3ce15746877b"
					},
					{
						"name": "newColumn2",
						"type": "String",
						"uuid": "8207d444-3a81-446f-9b59-a6a69f4fe771"
					},
					{
						"name": "newColumn3",
						"type": "String",
						"uuid": "bf4228e8-5e6f-49a1-8429-f5f15813dc91"
					}
				],
				"entityType": "Normal"
			}
		],
		"relations": [
			{
				"uuid": "c635a4c2-f98b-416e-9333-6d9cf22fc1d3",
				"ownerId": "2fac8cfc-fc71-446a-a9cf-36056a63ba78",
				"sourceId": "2fac8cfc-fc71-446a-a9cf-36056a63ba78",
				"targetId": "3e9ae743-de18-4b0c-a77e-3726be4049a8",
				"relationType": "manyToMany",
				"roleOnSource": "newentity41",
				"roleOnTarget": "newentity32"
			}
		]
	}	`
	m := MetaContent{}
	json.Unmarshal([]byte(dataString), &m)

	if len(m.Entities) != 2 {
		t.Errorf("Entities number is %d ,not 2", len(m.Entities))
	}

	tables := m.Tables()

	if len(tables) != 3 {
		t.Errorf("Tables number is %d ,not 3", len(tables))
	}
	//t.FailNow()
}
