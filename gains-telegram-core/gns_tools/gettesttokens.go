package gns_tools

import "log"

func GetCollateralTokenContractAddress(collateral uint8, chain uint64) string {
	switch chain {
	case 137:
		switch collateral {
		case 1:
			return "0x8f3cf7ad23cd3cadbd9735aff958023239c6a063"
		case 2:
			return "0x7ceb23fd6bc0add59e62ac25578270cff1b9f619"
		case 3:
			return "0x2791bca1f2de4661ed88a30c99a7a9449aa84174"
		}

	case 42161:
		switch collateral {
		case 1:
			return "0xda10009cbd5d07dd0cecc66161fc93d7c9000da1"
		case 2:
			return "0x82af49447d8a07e3bd95bd0d56f35241523fbab1"
		case 3:
			return "0xaf88d065e77c8cc2239327c5edb3a432268e5831"
		}

	case 421614:
		switch collateral {
		case 1:
			return "0xfbb7e7fee1525958bf5a4f04ed8d7be547ab6d27" // TestDai
		case 2:
			return "0x4200000000000000000000000000000000000006" // Weth
		case 3:
			return "0x4cc7ebeed5ea3adf3978f19833d2e1f3e8980cd6" // TestUsdc
		}
	default:
		log.Println("Invalid chain: ", chain)
	}
	return ""
}
