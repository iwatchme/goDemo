package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello( name string) (string,error) {

	if name == "" {
       return "", errors.New("empty name")
	}

	msg := fmt.Sprintf(randomFormat(), name)

    return msg, nil
}




func init() {
	fmt.Println("init")
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"Hello1 %v",
		"Hello2 %v",
		"Hello3 %v",
	}

	return formats[rand.Intn(len(formats))]
}
