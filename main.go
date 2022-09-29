package main

import (
	"github.com/brianvoe/gofakeit/v6"
	"math/rand"
	"time"
	"encoding/json"
	"bufio"
	"log"
	"os"
)

type Step struct {
	Step1 		string  	`json:"step1"`
	Timestamp1  int64		`json:"timestamp1"`
	Step2 		string		`json:"step2"`
	Timestamp2  int64		`json:"timestamp2"`
	Step3 		string		`json:"step3"`
	Timestamp3  int64		`json:"timestamp3"`
	Step4 		string		`json:"step4"`
	Timestamp4  int64		`json:"timestamp4"`
}

type LogQueue struct {
	Event			string  `json:"event"`
	Id 				string	`json:"id"`
	Type 			string  `json:"type"`
	Channel_number	string  `json:"channel_number"`
	Steps	 		Step    `json:"steps"`
	Number_of_tries int     `json:"number_of_tries"`
	Status          bool    `json:"status"`
}

func main() {
	GenerateRandomData(1000)
}

func RandomChoice(l int) int{
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	ts := r.Intn(l)
	return ts
}

func GenerateRandomData(n int64){
	var teste []LogQueue
	mt := []string{"text", "link", "image", "video"}
	ta := []int64{5,10,12,16,20,24,28,30,50,60}

	var i int64

	for i = 0; i < n; i++ {
		ts := RandomChoice(len(mt))
		t1 := RandomChoice(len(ta))
		t2 := RandomChoice(len(ta))
		t3 := RandomChoice(len(ta))

		s := Step{
			Step1: "verify_queue",
			Timestamp1: time.Now().Unix()+i,
			Step2: "send_to_route",
			Timestamp2: time.Now().Unix()+int64(t1)+i,
			Step3: "wait_for_confirmation",
			Timestamp3: time.Now().Unix()+int64(t1)+int64(t2)+i,
			Step4: "remove_from_queue",
			Timestamp4: time.Now().Unix()+int64(t1)+int64(t2)+int64(t3)+i}

		l := LogQueue{
			Event: "send_message",
			Id: gofakeit.UUID(),
			Type: mt[ts],
			Channel_number: gofakeit.Phone(),
			Steps: s,
			Number_of_tries: gofakeit.Number(1, 7),
			Status: gofakeit.Bool(),
		}
		teste = append(teste, l)
	}

	jp, _ := json.Marshal(teste)
	file, err := os.OpenFile("teste.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
 
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
 
	datawriter := bufio.NewWriter(file)
 
	_, _ = datawriter.WriteString(string(jp) + "\n")
 
	datawriter.Flush()
	file.Close()
}
