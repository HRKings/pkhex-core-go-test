package gen3

type Gen3Game int8

const (
	RubySapphire     Gen3Game = 0
	FireRedLeafGreen Gen3Game = 1
	Emerald          Gen3Game = 2
)

type KeyCode struct {
	game_code    uint32
	security_key uint32
}

type PlayedTime struct {
	time    []byte
	hours   uint16
	minutes byte
	seconds byte
	frames  byte
}

type TrainerId struct {
	trainer_id uint32
	secret_id  uint16
	public_id  uint16
}

type TrainerData struct {
	section_info SectionData
	name         []byte
	gender       byte
	id           TrainerId
	time         PlayedTime
	security     KeyCode
}

type SectionData struct {
	section_id uint16
	checksum   uint16
	signature  uint32
	save_index uint32
}

type SaveGen3 struct {
	gameVer        Gen3Game
	TrainerSection TrainerData
}
