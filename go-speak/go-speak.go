package main

import (
	"fmt"

	speech "github.com/nicolaifsf/go-speak" //Importerer go-speak
)

func main() {
	speech.SetWitKey("EGBCQYQCVKPQKIX2BNQDLKJA4SVZLOLE") //Setter API nøkkel til wit
	responce := speech.SendWitVoice("Hei.wav")           //Sender lyfil til wit som konverter tale til tekst og returner JSON struktur med teksten
	fmt.Println(responce)
}
