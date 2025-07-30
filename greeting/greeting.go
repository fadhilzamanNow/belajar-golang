package greeting

import (
	"fmt"
	"errors"
	"math/rand"
)

func Hello(name string) (string,error) {
	if(name == ""){
		return "", errors.New("nama gak boleh kosong ya kocak");
	}
	message := fmt.Sprintf(randomMessage(), name);
	return message,nil;
}

func Hellos(names []string) (map[string]string, error){
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil,err
		}
		messages[name] = message
	}
	return messages, nil
}

func randomMessage() string {
	formats := []string{
		"Hello %v",
		"Hai %v sayang",
		"Kelas dekk %v",
	}
	return formats[rand.Intn(len(formats))]
}
