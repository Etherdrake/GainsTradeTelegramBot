use std::str::FromStr;
use ethcontract::private::lazy_static;
use primitive_types::H160;


lazy_static! {
    pub static ref GNS_TRADING_CONTRACT_SEPOLIA_DAI: H160 =
        H160::from_str("0x1D29c95Fa9F47987ede5121700881ddaa9116B29").unwrap();
    pub static ref GNS_TRADING_CONTRACT_SEPOLIA_WETH: H160 =
        H160::from_str("0xb0901FEaD3112f6CaF9353ec5c36DC3DdE111F61").unwrap();
    pub static ref GNS_TRADING_CONTRACT_SEPOLIA_USDC: H160 =
        H160::from_str("0xb0901FEaD3112f6CaF9353ec5c36DC3DdE111F61").unwrap();
    pub static ref GNS_DIAMOND_CONTRACT_SEPOLIA: H160 =
        H160::from_str("0x209A9A01980377916851af2cA075C2b170452018").unwrap();
    pub static ref GNS_STORAGE_SEPOLIA_DAI: H160 =
        H160::from_str("0xaee4d11a16B2bc65EDD6416Fb626EB404a6D65BD").unwrap();
    pub static ref GNS_STORAGE_SEPOLIA_WETH: H160 =
        H160::from_str("0xE7712ebcd451919B38Be8fD102800A496C5BeD4E").unwrap();
    pub static ref GNS_STORAGE_SEPOLIA_USDC: H160 =
        H160::from_str("0xFe54a9A1C2C276cf37C56CeeE30737FDc6dA4d27").unwrap();
    pub static ref GFARM_V5_CONTRACT_SEPOLIA_DAI: H160 =
        H160::from_str("0xaee4d11a16B2bc65EDD6416Fb626EB404a6D65BD").unwrap();
    pub static ref GFARM_V5_CONTRACT_SEPOLIA_WETH: H160 =
        H160::from_str("0xE7712ebcd451919B38Be8fD102800A496C5BeD4E").unwrap();
    pub static ref GFARM_V5_CONTRACT_SEPOLIA_USDC: H160 =
        H160::from_str("0xC504C9C30B9d88cBc9704Fc2d06a08A4c7bE9378").unwrap();
}