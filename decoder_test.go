package mapstringinterfacedecoder

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
		"Susie Noble,friends,0,Diane Martinez\n",
		"Susie Noble,friends,1,Keller Schneider\n",
		"Susie Noble,friends,2,Imogene Kemp\n",
	}
	exampleInput := make(map[string]interface{})
	err := json.Unmarshal([]byte(randomRubbishJSON), &exampleInput)
	if err != nil {
		log.Println(err)
		T.Fail()
	}
	d := Decoder{}
	d.Decode(exampleInput)
	b := d.Get("friends")
	actualOutput := []string{}
	for i := range b {
		out := fmt.Sprintf("%s,%s,%s,%s\n", d.Field("name"), b[i].Name, b[i].Field("id"), b[i].Field("name"))
		actualOutput = append(actualOutput, out)
	}

	if !reflect.DeepEqual(actualOutput, testOutput) {
		T.Fail()
	}
}
