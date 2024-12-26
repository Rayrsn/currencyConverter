package utils

import (
	"strconv"
)

func IsNumber(s string) bool {
    _, err := strconv.ParseFloat(s, 64)
    return err == nil
}

func IsCurrency(code string) bool {
    switch code {
	case "00":
		return true

	case "1INCH":
		return true

	case "AAVE":
		return true

	case "ABT":
		return true

	case "ACH":
		return true

	case "ACS":
		return true

	case "ADA":
		return true

	case "AED":
		return true

	case "AERGO":
		return true

	case "AFN":
		return true

	case "AGLD":
		return true

	case "AIOZ":
		return true

	case "AKT":
		return true

	case "ALCX":
		return true

	case "ALEPH":
		return true

	case "ALGO":
		return true

	case "ALICE":
		return true

	case "ALL":
		return true

	case "AMD":
		return true

	case "AMP":
		return true

	case "ANG":
		return true

	case "ANKR":
		return true

	case "ANT":
		return true

	case "AOA":
		return true

	case "APE":
		return true

	case "API3":
		return true

	case "APT":
		return true

	case "AR":
		return true

	case "ARB":
		return true

	case "ARPA":
		return true

	case "ARS":
		return true

	case "ASM":
		return true

	case "AST":
		return true

	case "ATA":
		return true

	case "ATOM":
		return true

	case "ATS":
		return true

	case "AUCTION":
		return true

	case "AUD":
		return true

	case "AUDIO":
		return true

	case "AURORA":
		return true

	case "AVAX":
		return true

	case "AVT":
		return true

	case "AWG":
		return true

	case "AXL":
		return true

	case "AXS":
		return true

	case "AZM":
		return true

	case "AZN":
		return true

	case "BADGER":
		return true

	case "BAKE":
		return true

	case "BAL":
		return true

	case "BAM":
		return true

	case "BAND":
		return true

	case "BAT":
		return true

	case "BBD":
		return true

	case "BCH":
		return true

	case "BDT":
		return true

	case "BEF":
		return true

	case "BGN":
		return true

	case "BHD":
		return true

	case "BICO":
		return true

	case "BIF":
		return true

	case "BIGTIME":
		return true

	case "BIT":
		return true

	case "BLUR":
		return true

	case "BLZ":
		return true

	case "BMD":
		return true

	case "BNB":
		return true

	case "BND":
		return true

	case "BNT":
		return true

	case "BOB":
		return true

	case "BOBA":
		return true

	case "BOND":
		return true

	case "BRL":
		return true

	case "BSD":
		return true

	case "BSV":
		return true

	case "BSW":
		return true

	case "BTC":
		return true

	case "BTCB":
		return true

	case "BTG":
		return true

	case "BTN":
		return true

	case "BTRST":
		return true

	case "BTT":
		return true

	case "BUSD":
		return true

	case "BWP":
		return true

	case "BYN":
		return true

	case "BYR":
		return true

	case "BZD":
		return true

	case "C98":
		return true

	case "CAD":
		return true

	case "CAKE":
		return true

	case "CBETH":
		return true

	case "CDF":
		return true

	case "CELO":
		return true

	case "CELR":
		return true

	case "CFX":
		return true

	case "CGLD":
		return true

	case "CHF":
		return true

	case "CHZ":
		return true

	case "CLP":
		return true

	case "CLV":
		return true

	case "CNH":
		return true

	case "CNY":
		return true

	case "COMP":
		return true

	case "COP":
		return true

	case "COTI":
		return true

	case "COVAL":
		return true

	case "CRC":
		return true

	case "CRO":
		return true

	case "CRPT":
		return true

	case "CRV":
		return true

	case "CSPR":
		return true

	case "CTSI":
		return true

	case "CTX":
		return true

	case "CUC":
		return true

	case "CUP":
		return true

	case "CVC":
		return true

	case "CVE":
		return true

	case "CVX":
		return true

	case "CYP":
		return true

	case "CZK":
		return true

	case "DAI":
		return true

	case "DAR":
		return true

	case "DASH":
		return true

	case "DCR":
		return true

	case "DDX":
		return true

	case "DEM":
		return true

	case "DESO":
		return true

	case "DEXT":
		return true

	case "DFI":
		return true

	case "DIA":
		return true

	case "DIMO":
		return true

	case "DJF":
		return true

	case "DKK":
		return true

	case "DNT":
		return true

	case "DOGE":
		return true

	case "DOP":
		return true

	case "DOT":
		return true

	case "DREP":
		return true

	case "DYP":
		return true

	case "DZD":
		return true

	case "EEK":
		return true

	case "EGLD":
		return true

	case "EGP":
		return true

	case "ELA":
		return true

	case "ENJ":
		return true

	case "ENS":
		return true

	case "EOS":
		return true

	case "ERN":
		return true

	case "ESP":
		return true

	case "ETB":
		return true

	case "ETC":
		return true

	case "ETH":
		return true

	case "ETH2":
		return true

	case "EUR":
		return true

	case "EUROC":
		return true

	case "FARM":
		return true

	case "FEI":
		return true

	case "FET":
		return true

	case "FIDA":
		return true

	case "FIL":
		return true

	case "FIM":
		return true

	case "FIS":
		return true

	case "FJD":
		return true

	case "FKP":
		return true

	case "FLOW":
		return true

	case "FLR":
		return true

	case "FORT":
		return true

	case "FORTH":
		return true

	case "FOX":
		return true

	case "FRAX":
		return true

	case "FRF":
		return true

	case "FTM":
		return true

	case "FTT":
		return true

	case "FX":
		return true

	case "FXS":
		return true

	case "GAL":
		return true

	case "GALA":
		return true

	case "GBP":
		return true

	case "GEL":
		return true

	case "GFI":
		return true

	case "GGP":
		return true

	case "GHC":
		return true

	case "GHS":
		return true

	case "GHST":
		return true

	case "GIP":
		return true

	case "GLM":
		return true

	case "GMD":
		return true

	case "GMT":
		return true

	case "GMX":
		return true

	case "GNF":
		return true

	case "GNO":
		return true

	case "GNT":
		return true

	case "GODS":
		return true

	case "GRD":
		return true

	case "GRT":
		return true

	case "GST":
		return true

	case "GT":
		return true

	case "GTC":
		return true

	case "GTQ":
		return true

	case "GUSD":
		return true

	case "GYD":
		return true

	case "GYEN":
		return true

	case "HBAR":
		return true

	case "HFT":
		return true

	case "HIGH":
		return true

	case "HKD":
		return true

	case "HNL":
		return true

	case "HNT":
		return true

	case "HOPR":
		return true

	case "HOT":
		return true

	case "HRK":
		return true

	case "HT":
		return true

	case "HTG":
		return true

	case "HUF":
		return true

	case "ICP":
		return true

	case "IDEX":
		return true

	case "IDR":
		return true

	case "IEP":
		return true

	case "ILS":
		return true

	case "ILV":
		return true

	case "IMP":
		return true

	case "IMX":
		return true

	case "INDEX":
		return true

	case "INJ":
		return true

	case "INR":
		return true

	case "INV":
		return true

	case "IOTX":
		return true

	case "IQD":
		return true

	case "IRR":
		return true

	case "ISK":
		return true

	case "ITL":
		return true

	case "JASMY":
		return true

	case "JEP":
		return true

	case "JMD":
		return true

	case "JOD":
		return true

	case "JPY":
		return true

	case "JUP":
		return true

	case "KAS":
		return true

	case "KAVA":
		return true

	case "KCS":
		return true

	case "KDA":
		return true

	case "KEEP":
		return true

	case "KES":
		return true

	case "KGS":
		return true

	case "KHR":
		return true

	case "KLAY":
		return true

	case "KMF":
		return true

	case "KNC":
		return true

	case "KPW":
		return true

	case "KRL":
		return true

	case "KRW":
		return true

	case "KSM":
		return true

	case "KWD":
		return true

	case "KYD":
		return true

	case "KZT":
		return true

	case "LAK":
		return true

	case "LBP":
		return true

	case "LCX":
		return true

	case "LDO":
		return true

	case "LEO":
		return true

	case "LINK":
		return true

	case "LIT":
		return true

	case "LKR":
		return true

	case "LOKA":
		return true

	case "LOOM":
		return true

	case "LPT":
		return true

	case "LQTY":
		return true

	case "LRC":
		return true

	case "LRD":
		return true

	case "LSETH":
		return true

	case "LSL":
		return true

	case "LTC":
		return true

	case "LTL":
		return true

	case "LUF":
		return true

	case "LUNA":
		return true

	case "LUNC":
		return true

	case "LVL":
		return true

	case "LYD":
		return true

	case "MAD":
		return true

	case "MAGIC":
		return true

	case "MANA":
		return true

	case "MASK":
		return true

	case "MATH":
		return true

	case "MATIC":
		return true

	case "MCO2":
		return true

	case "MDL":
		return true

	case "MDT":
		return true

	case "MEDIA":
		return true

	case "METIS":
		return true

	case "MGA":
		return true

	case "MGF":
		return true

	case "MINA":
		return true

	case "MIR":
		return true

	case "MKD":
		return true

	case "MKR":
		return true

	case "MLN":
		return true

	case "MMK":
		return true

	case "MNDE":
		return true

	case "MNT":
		return true

	case "MONA":
		return true

	case "MOP":
		return true

	case "MPL":
		return true

	case "MRO":
		return true

	case "MRU":
		return true

	case "MSOL":
		return true

	case "MTL":
		return true

	case "MULTI":
		return true

	case "MUR":
		return true

	case "MUSD":
		return true

	case "MUSE":
		return true

	case "MVR":
		return true

	case "MWK":
		return true

	case "MXC":
		return true

	case "MXN":
		return true

	case "MXV":
		return true

	case "MYR":
		return true

	case "MZM":
		return true

	case "MZN":
		return true

	case "NAD":
		return true

	case "NCT":
		return true

	case "NEAR":
		return true

	case "NEO":
		return true

	case "NEST":
		return true

	case "NEXO":
		return true

	case "NFT":
		return true

	case "NGN":
		return true

	case "NIO":
		return true

	case "NKN":
		return true

	case "NLG":
		return true

	case "NMR":
		return true

	case "NOK":
		return true

	case "NPR":
		return true

	case "NU":
		return true

	case "NZD":
		return true

	case "OCEAN":
		return true

	case "OGN":
		return true

	case "OKB":
		return true

	case "OMG":
		return true

	case "OMR":
		return true

	case "ONE":
		return true

	case "OOKI":
		return true

	case "OP":
		return true

	case "ORCA":
		return true

	case "ORN":
		return true

	case "OSMO":
		return true

	case "OXT":
		return true

	case "PAB":
		return true

	case "PAX":
		return true

	case "PAXG":
		return true

	case "PEN":
		return true

	case "PEPE":
		return true

	case "PERP":
		return true

	case "PGK":
		return true

	case "PHP":
		return true

	case "PKR":
		return true

	case "PLA":
		return true

	case "PLN":
		return true

	case "PLU":
		return true

	case "PNG":
		return true

	case "POLS":
		return true

	case "POLY":
		return true

	case "POND":
		return true

	case "POWR":
		return true

	case "PRIME":
		return true

	case "PRO":
		return true

	case "PRQ":
		return true

	case "PTE":
		return true

	case "PUNDIX":
		return true

	case "PYG":
		return true

	case "PYR":
		return true

	case "PYUSD":
		return true

	case "QAR":
		return true

	case "QI":
		return true

	case "QNT":
		return true

	case "QSP":
		return true

	case "QTUM":
		return true

	case "QUICK":
		return true

	case "RAD":
		return true

	case "RAI":
		return true

	case "RARE":
		return true

	case "RARI":
		return true

	case "RBN":
		return true

	case "REN":
		return true

	case "REP":
		return true

	case "REPV2":
		return true

	case "REQ":
		return true

	case "RGT":
		return true

	case "RLC":
		return true

	case "RLY":
		return true

	case "RNDR":
		return true

	case "ROL":
		return true

	case "RON":
		return true

	case "ROSE":
		return true

	case "RPL":
		return true

	case "RSD":
		return true

	case "RUB":
		return true

	case "RUNE":
		return true

	case "RVN":
		return true

	case "RWF":
		return true

	case "SAND":
		return true

	case "SAR":
		return true

	case "SBD":
		return true

	case "SCR":
		return true

	case "SDD":
		return true

	case "SDG":
		return true

	case "SEI":
		return true

	case "SEK":
		return true

	case "SGD":
		return true

	case "SHIB":
		return true

	case "SHP":
		return true

	case "SHPING":
		return true

	case "SIT":
		return true

	case "SKK":
		return true

	case "SKL":
		return true

	case "SLE":
		return true

	case "SLL":
		return true

	case "SNT":
		return true

	case "SNX":
		return true

	case "SOL":
		return true

	case "SOS":
		return true

	case "SPA":
		return true

	case "SPELL":
		return true

	case "SPL":
		return true

	case "SRD":
		return true

	case "SRG":
		return true

	case "SSP":
		return true

	case "STD":
		return true

	case "STG":
		return true

	case "STN":
		return true

	case "STORJ":
		return true

	case "STX":
		return true

	case "SUI":
		return true

	case "SUKU":
		return true

	case "SUPER":
		return true

	case "SUSHI":
		return true

	case "SVC":
		return true

	case "SWFTC":
		return true

	case "SYLO":
		return true

	case "SYN":
		return true

	case "SYP":
		return true

	case "SZL":
		return true

	case "T":
		return true

	case "THB":
		return true

	case "THETA":
		return true

	case "TIA":
		return true

	case "TIME":
		return true

	case "TJS":
		return true

	case "TMM":
		return true

	case "TMT":
		return true

	case "TND":
		return true

	case "TON":
		return true

	case "TONE":
		return true

	case "TOP":
		return true

	case "TRAC":
		return true

	case "TRB":
		return true

	case "TRIBE":
		return true

	case "TRL":
		return true

	case "TRU":
		return true

	case "TRX":
		return true

	case "TRY":
		return true

	case "TTD":
		return true

	case "TTT":
		return true

	case "TUSD":
		return true

	case "TVD":
		return true

	case "TVK":
		return true

	case "TWD":
		return true

	case "TWT":
		return true

	case "TZS":
		return true

	case "UAH":
		return true

	case "UGX":
		return true

	case "UMA":
		return true

	case "UNFI":
		return true

	case "UNI":
		return true

	case "UPI":
		return true

	case "USD":
		return true

	case "USDC":
		return true

	case "USDD":
		return true

	case "USDP":
		return true

	case "USDT":
		return true

	case "UST":
		return true

	case "UYU":
		return true

	case "UZS":
		return true

	case "VAL":
		return true

	case "VARA":
		return true

	case "VEB":
		return true

	case "VED":
		return true

	case "VEF":
		return true

	case "VES":
		return true

	case "VET":
		return true

	case "VGX":
		return true

	case "VND":
		return true

	case "VOXEL":
		return true

	case "VTHO":
		return true

	case "VUV":
		return true

	case "WAMPL":
		return true

	case "WAVES":
		return true

	case "WAXL":
		return true

	case "WBTC":
		return true

	case "WCFG":
		return true

	case "WEMIX":
		return true

	case "WLUNA":
		return true

	case "WOO":
		return true

	case "WST":
		return true

	case "XAF":
		return true

	case "XAG":
		return true

	case "XAU":
		return true

	case "XAUT":
		return true

	case "XBT":
		return true

	case "XCD":
		return true

	case "XCH":
		return true

	case "XCN":
		return true

	case "XDC":
		return true

	case "XDR":
		return true

	case "XEC":
		return true

	case "XEM":
		return true

	case "XLM":
		return true

	case "XMON":
		return true

	case "XMR":
		return true

	case "XOF":
		return true

	case "XPD":
		return true

	case "XPF":
		return true

	case "XPT":
		return true

	case "XRP":
		return true

	case "XTZ":
		return true

	case "XYO":
		return true

	case "YER":
		return true

	case "YFI":
		return true

	case "YFII":
		return true

	case "ZAR":
		return true

	case "ZEC":
		return true

	case "ZEN":
		return true

	case "ZIL":
		return true

	case "ZMK":
		return true

	case "ZMW":
		return true

	case "ZRX":
		return true

	case "ZWD":
		return true

	case "ZWL":
		return true

	default:
		return false
	}
}
