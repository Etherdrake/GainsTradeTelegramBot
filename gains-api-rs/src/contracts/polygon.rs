use std::str::FromStr;
use ethcontract::private::lazy_static;
use primitive_types::H160;


lazy_static! {
    pub static ref GNS_DIAMOND_CONTRACT_POLYGON: H160 =
        H160::from_str("0x209A9A01980377916851af2cA075C2b170452018").unwrap();
    pub static ref GNS_TRADING_CONTRACT_POLYGON_DAI: H160 =
        H160::from_str("0x1D29c95Fa9F47987ede5121700881ddaa9116B29").unwrap();
    pub static ref GNS_TRADING_CONTRACT_POLYGON_WETH: H160 =
        H160::from_str("0xb0901FEaD3112f6CaF9353ec5c36DC3DdE111F61").unwrap();
    pub static ref GNS_TRADING_CONTRACT_POLYGON_USDC: H160 =
        H160::from_str("0xb0901FEaD3112f6CaF9353ec5c36DC3DdE111F61").unwrap();
    pub static ref GNS_TRADING_STORAGE_POLYGON: H160 =
        H160::from_str("0xaee4d11a16B2bc65EDD6416Fb626EB404a6D65BD").unwrap();
}
