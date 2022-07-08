package gen3

import (
	"PKHex_Go/utils"
	"encoding/binary"
)

func getSecurityKeyOrGameCode(bytesA []byte, bytesB []byte, save SaveGen3) KeyCode {
	var blockA = binary.LittleEndian.Uint32(bytesA)
	var blockB = binary.LittleEndian.Uint32(bytesB)

	if blockA == 0x00000000 {
		save.gameVer = RubySapphire
		return KeyCode{
			game_code:    0x00000000,
			security_key: 0,
		}
	} else if blockA == 0x00000001 {
		save.gameVer = FireRedLeafGreen
		return KeyCode{
			game_code:    0x00000001,
			security_key: blockB,
		}
	} else {
		save.gameVer = Emerald
		return KeyCode{
			game_code:    0,
			security_key: blockA,
		}
	}
}

func ParseTrainer(sectionBytes []byte, save SaveGen3) TrainerData {
	return TrainerData{
		section_info: SectionData{
			section_id: binary.LittleEndian.Uint16(utils.GetOffset(sectionBytes, 0x0FF4, 2)),
			checksum:   binary.LittleEndian.Uint16(utils.GetOffset(sectionBytes, 0x0FF6, 2)),
			signature:  binary.LittleEndian.Uint32(utils.GetOffset(sectionBytes, 0x0FF8, 4)),
			save_index: binary.LittleEndian.Uint32(utils.GetOffset(sectionBytes, 0x0FFC, 4)),
		},
		name:   utils.GetOffset(sectionBytes, 0x0, 7),
		gender: sectionBytes[8],
		id: TrainerId{
			trainer_id: binary.LittleEndian.Uint32(utils.GetOffset(sectionBytes, 0x000A, 4)),
			secret_id:  binary.LittleEndian.Uint16(utils.GetOffset(sectionBytes, 0x000A, 2)),
			public_id:  binary.LittleEndian.Uint16(utils.GetOffset(sectionBytes, 0x000A+2, 2)),
		},
		time: PlayedTime{
			time:    utils.GetOffset(sectionBytes, 0x000E, 5),
			hours:   binary.LittleEndian.Uint16(utils.GetOffset(sectionBytes, 0x000E, 2)),
			minutes: sectionBytes[0x000E+2],
			seconds: sectionBytes[0x000E+3],
			frames:  sectionBytes[0x000E+4],
		},
		security: getSecurityKeyOrGameCode(
			utils.GetOffset(sectionBytes, 0x00AC, 4),
			utils.GetOffset(sectionBytes, 0x0AF8, 4), save),
	}
}
