use std::str::FromStr;
use ethcontract::private::lazy_static;
use primitive_types::H160;

lazy_static! {
    pub static ref DIAMOND_ARBITRUM: H160 =
        H160::from_str("0x209A9A01980377916851af2cA075C2b170452018").unwrap();
    pub static ref DIAMOND_POLYGON: H160 =
        H160::from_str("0x1D29c95Fa9F47987ede5121700881ddaa9116B29").unwrap();
    pub static ref DIAMOND_SEPOLIA: H160 =
        H160::from_str("0xb0901FEaD3112f6CaF9353ec5c36DC3DdE111F61").unwrap();
}