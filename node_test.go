package map2struct

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"reflect"
)

var randomRubbishJSON = `{
	"_id": "617c120ab7a3fa5dd968346d",
	"index": 0,
	"guid": "f3a4d826-30d6-4811-ba08-cda75b221f2b",
	"isActive": true,
	"balance": "$1,241.70",
	"picture": "http://placehold.it/32x32",
	"age": 37,
	"eyeColor": "brown",
	"name": "Susie Noble",
	"gender": "female",
	"company": "ZILLA",
	"email": "susienoble@zilla.com",
	"phone": "+1 (925) 507-2761",
	"address": "266 Legion Street, Cressey, Illinois, 8982",
	"about": "Aliquip ut veniam mollit duis non minim ad amet est ex id fugiat consequat. Amet labore nulla esse adipisicing velit fugiat aute fugiat in Lorem aliquip et exercitation. Proident laborum aliqua laborum esse Lorem minim. Velit do aute cillum laborum sit excepteur minim fugiat cillum laborum ex enim. Consequat minim nisi adipisicing eu deserunt esse id id anim esse consequat sunt cillum deserunt. Sit commodo aute adipisicing irure nulla in aliquip pariatur eiusmod consectetur velit. Ea anim nulla cillum eu fugiat.\r\n",
	"registered": "2017-08-29T12:01:21 -01:00",
	"latitude": -39.228822,
	"longitude": 136.128045,
	"tags": [
	  "consequat",
	  "magna",
	  "exercitation",
	  "eu",
	  "minim",
	  "duis",
	  "qui"
	],
	"friends": [
	  {
		"id": 0,
		"name": "Diane Martinez"
	  },
	  {
		"id": 1,
		"name": "Keller Schneider"
	  },
	  {
		"id": 2,
		"name": "Imogene Kemp"
	  }
	],
	"greeting": "Hello, Susie Noble! You have 7 unread messages.",
	"favoriteFruit": "strawberry"
  }`

func TestDecode(T *testing.T) {
	testOutput := []string{
		"Susie Noble,friends,0,Diane Martinez,consequat",
		"Susie Noble,friends,1,Keller Schneider,consequat",
		"Susie Noble,friends,2,Imogene Kemp,consequat",
		"Susie Noble,friends,0,Diane Martinez,magna",
		"Susie Noble,friends,1,Keller Schneider,magna",
		"Susie Noble,friends,2,Imogene Kemp,magna",
		"Susie Noble,friends,0,Diane Martinez,exercitation",
		"Susie Noble,friends,1,Keller Schneider,exercitation",
		"Susie Noble,friends,2,Imogene Kemp,exercitation",
		"Susie Noble,friends,0,Diane Martinez,eu",
		"Susie Noble,friends,1,Keller Schneider,eu",
		"Susie Noble,friends,2,Imogene Kemp,eu",
		"Susie Noble,friends,0,Diane Martinez,minim",
		"Susie Noble,friends,1,Keller Schneider,minim",
		"Susie Noble,friends,2,Imogene Kemp,minim",
		"Susie Noble,friends,0,Diane Martinez,duis",
		"Susie Noble,friends,1,Keller Schneider,duis",
		"Susie Noble,friends,2,Imogene Kemp,duis",
		"Susie Noble,friends,0,Diane Martinez,qui",
		"Susie Noble,friends,1,Keller Schneider,qui",
		"Susie Noble,friends,2,Imogene Kemp,qui",
	}
	exampleInput := make(map[string]interface{}) // using JSON as an example dataset - we really want map[string]interface{}
	err := json.Unmarshal([]byte(randomRubbishJSON), &exampleInput)
	if err != nil {
		log.Println(err)
		T.Fail()
	}
	root := Node{}
	root.Ingest(exampleInput)
	friends := root.Get("friends") // get a list of the friends structs
	tags := root.Array("tags")     // grab the tags array
	actualOutput := []string{}
	for tagindex := range tags {
		for friendindex := range friends {
			out := fmt.Sprintf("%s,%s,%s,%s,%s", root.Field("name"),
				friends[friendindex].Name, friends[friendindex].Field("id"),
				friends[friendindex].Field("name"), tags[tagindex])
			actualOutput = append(actualOutput, out)
		}
	}

	for _, v := range actualOutput {
		log.Println(v)
	}

	if !reflect.DeepEqual(actualOutput, testOutput) {

		T.Fail()
	}
}
