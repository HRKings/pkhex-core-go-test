package main

import (
	"PKHex_Go/save/gen3"
	"PKHex_Go/utils"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	fileBytes, err := ioutil.ReadFile("assets/test.sav")
	if err != nil {
		log.Fatal(err)
	}

	var saveBBytes = utils.GetOffset(fileBytes, 0x00E000, 57344)

	var save = gen3.SaveGen3{}

	for i := 0; i < 14; i++ {
		var currentSection = utils.GetOffset(saveBBytes, 4096*i, 4096)
		var sectionId = binary.LittleEndian.Uint16(utils.GetOffset(currentSection, 0x0FF4, 2))

		switch sectionId {
		case 0:
			save.TrainerSection = gen3.ParseTrainer(currentSection, save)
			break

		default:
			fmt.Println("The section ID is", sectionId)
			break
		}
	}

	fmt.Printf("%+v\n", save)
}
