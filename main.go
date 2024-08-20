package main

import (
	"fmt"
	"log"
	"time"

	"github.com/sirjager/gopkg/tokens"
	"github.com/sirjager/gopkg/utils"
)

type Payload struct {
	UserEmail string `json:"email"`
}

func main() {
	builder, err := tokens.NewPasetoBuilder(utils.RandomString(32))
	if err != nil {
		log.Fatal(err)
	}
	data := &Payload{UserEmail: utils.RandomEmail()}
	token, _, err := builder.CreateToken(data, time.Hour)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(token)
}
