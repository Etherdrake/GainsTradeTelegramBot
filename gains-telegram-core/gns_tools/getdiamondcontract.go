package gns_tools

import "log"

func GetDiamondContractAddress(chain uint64) string {
	switch chain {
	case 137:
		return "0x209A9A01980377916851af2cA075C2b170452018"
	case 42161:
		return "0xFF162c694eAA571f685030649814282eA457f169"
	case 421614:
		return "0xd659a15812064C79E189fd950A189b15c75d3186"
	default:
		log.Println("Invalid chain: ", chain)
	}
	return ""
}
