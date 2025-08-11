use std::str::FromStr;
use ethcontract::private::lazy_static;
use primitive_types::H160;

lazy_static! {
    pub static ref GNS_DIAMOND_CONTRACT_ARB: H160 =
        H160::from_str("0xFF162c694eAA571f685030649814282eA457f169").unwrap();
    pub static ref GNS_STORAGE_ARB_DAI: H160 =
        H160::from_str("0xcFa6ebD475d89dB04cAd5A756fff1cb2BC5bE33c").unwrap();
    pub static ref GNS_STORAGE_ARB_WETH: H160 =
        H160::from_str("0xFe54a9A1C2C276cf37C56CeeE30737FDc6dA4d27").unwrap();
    pub static ref GNS_STORAGE_ARB_USDC: H160 =
        H160::from_str("0x3B09fCa4cC6b140fDd364f28db830ccE01Fd60fD").unwrap();
    pub static ref GNS_TRADING_CONTRACT_ARB_DAI: H160 =
        H160::from_str("0x2c7e82641f03Fa077F88833213210A86027f15dc").unwrap();
    pub static ref GNS_TRADING_CONTRACT_ARB_WETH: H160 =
        H160::from_str("0x48B07695c41AaC54CC35F56AF25573dd19235c6f").unwrap();
    pub static ref GNS_TRADING_CONTRACT_ARB_USDC: H160 =
        H160::from_str("0x2FE799d81FDfCC441093eaB52Af788d4Cc6Ff650").unwrap();
    pub static ref GFARM_V5_CONTRACT_ARB_DAI: H160 =
        H160::from_str("0xcFa6ebD475d89dB04cAd5A756fff1cb2BC5bE33c").unwrap();
    pub static ref GFARM_V5_CONTRACT_ARB_WETH: H160 =
        H160::from_str("0xFe54a9A1C2C276cf37C56CeeE30737FDc6dA4d27").unwrap();
    pub static ref GFARM_V5_CONTRACT_ARB_USDC: H160 =
        H160::from_str("0x3B09fCa4cC6b140fDd364f28db830ccE01Fd60fD").unwrap();
}