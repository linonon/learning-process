package main

import (
	"encoding/json"
	"fmt"
	trippb "ginstart/protogo/proto"

	"google.golang.org/protobuf/proto"
)

func main() {
	trip := trippb.Trip{
		Start:       "one",
		End:         "twenty",
		DurationSec: 9900,
		FeeCent:     2800,
		// StartPos 為指針所指向的Location結構，所以要用 “&”
		StartPos: &trippb.Location{
			Latitude:  30,
			Longitude: 120,
		},
		// 同理
		EndPos: &trippb.Location{
			Longitude: 115,
			Latitude:  35,
		},
		PathPos: []*trippb.Location{
			{
				Latitude:  31,
				Longitude: 119,
			},
			{
				Latitude:  32,
				Longitude: 118,
			},
		},
		Status: trippb.TripStatus_IN_PROGRESS,
	}
	fmt.Println(&trip)

	b, err := proto.Marshal(&trip)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("%X\n", b)
	}

	var trip2 trippb.Trip
	err = proto.Unmarshal(b, &trip2)
	if err != nil {
		panic(err)
	}
	fmt.Println(&trip2)

	b, err = json.Marshal(&trip2)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("%s\n", b)
	}
}
