# generic mapStringInterface decoder

```TL;DR - use this like the pandas library but for map[string]interface{} to perform basic data extraction```
here is some EXAMPLE code:

```
package main

import (
	"encoding/json"
	"log"
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
```

```
func main() {
	exampleInput := make(map[string]interface{})
	err := json.Unmarshal([]byte(randomRubbishJSON), &exampleInput)
	if err != nil {
		log.Println(err)
		return
	}
	d := Decoder{}
	d.Decode(exampleInput)
	d.Print()
}
```

## what is output of d.Print() ?
```
Name: root
Root: 
field: [phone greeting latitude _id balance email longitude favoriteFruit age about registered name gender company guid isActive picture index eyeColor address]
array: [tags]
sub: [friends]
```

## ok so let's explore:

```
	for _, field := range d.Fields() {
		log.Println(field)
	}
```

## and:

```
	b := d.Get("friends")
	for i := range b {
		b[i].Print()
	}
  
  which results:
  
Name: friends
Root: root
field: [id name]
array: []
sub: []
```

## so now let's do something useful:
```
func main() {
	exampleInput := make(map[string]interface{})
	err := json.Unmarshal([]byte(randomRubbishJSON), &exampleInput)
	if err != nil {
		log.Println(err)
		return
	}
	d := Decoder{}
	d.Decode(exampleInput)
	b := d.Get("friends")
	for i := range b {
		fmt.Printf("%s,%s,%s,%s\n", d.Field("name"), b[i].Name, b[i].Field("id"), b[i].Field("name"))
	}
}

output:
Susie Noble,friends,0,Diane Martinez
Susie Noble,friends,1,Keller Schneider
Susie Noble,friends,2,Imogene Kemp
```

## with this decoder you can basically create any CSV dataset you want with the right instructions.
## and it's much _much_ easier than playing around with map[string]interface{}

Have fun!

- Tom
