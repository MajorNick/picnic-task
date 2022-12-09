package parser

import (
	"encoding/csv"
	
	"log"
	"os"
	"strconv"
	oset "github.com/MajorNick/orderedSet"
)

type Respodent struct{
	Question string
	SegmentType  string
	SegmentDesc string
	Answer 	 string
	Count int
	Perc float64
} 
var respodents  []Respodent



func init(){
	file, err := os.Open("WhatsgoodlyData-10.csv")
	if err != nil{
		log.Fatalf("Error appeard in Oppening file %v", err)
	}
	defer file.Close()
	
	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err!=nil{
		log.Fatalf("Error appeard in Reading file %v", err)
	}
	respodents = ReadData(data)

}


func ReadData(data [][]string)[]Respodent{
	

	for i,v:= range data{
		if i>0{
			
		tmpResp:=dataProcessing(v)
		respodents = append(respodents,tmpResp)

		}
	}
	return respodents
}

func dataProcessing(line []string) Respodent {
	res := Respodent{}
	var err error
	for i,v := range line{ 
		switch i{
		case 0:
			res.Question = v
		case 1:
			res.SegmentType = v
		case 2:
			res.SegmentDesc = v
		case 3:
			res.Answer = v
		case 4:
			res.Count,err = strconv.Atoi(v)
		case 5:
			res.Perc,err = strconv.ParseFloat(v,8)
		default:
			log.Printf("Error in Parsing Line: %v", err)
		}
	}
	if err != nil{
		log.Printf("Error in Parsing Line: %v", err)
	}
	return res 
}


func GetRawData()*[]Respodent{
	return &respodents
}


func GetMappedDataOfAnswers()map[string]int{
	mp := make(map[string]int)
	
	for _,v := range respodents{
		mp[v.Answer] += v.Count
	}
	return mp
}

func GenerateOSet()oset.OrderedSet{
	cmp := func (a,b interface{}) int {
		a1 := a.(Respodent)
		b1 := b.(Respodent)
		   if a1.Count == b1.Count {
			   if a1.SegmentDesc>b1.SegmentDesc{
				return 1
			   }else{
				return -1
			   }
		}
		if(a1.Count<b1.Count){
			return 1
		}
		return -1
	}
	set := oset.NewSet(cmp)
	for _,v:=range respodents{
		set.Insert(v)
	}
	return set
}


