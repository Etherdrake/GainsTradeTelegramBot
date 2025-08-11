use lazy_static::lazy_static;
use std::str::FromStr;
use primitive_types::H160;

lazy_static! {
    pub static ref DAI_ADDRESS_ETHEREUM: H160 =
        H160::from_str("0x6B175474E89094C44Da98b954EedeAC495271d0F").unwrap();
    pub static ref DAI_ADDRESS_BINANCE_SMART_CHAIN: H160 =
        H160::from_str("0x1AF3F329e8BE154074D8769D1FFa4eE058B1DBc3").unwrap();

    // Polygon
    pub static ref DAI_ADDRESS_POLYGON: H160 =
        H160::from_str("0x8f3cf7ad23cd3cadbd9735aff958023239c6a063").unwrap();
    pub static ref WETH_ADDRESS_POLYGON: H160 =
        H160::from_str("0x7ceb23fd6bc0add59e62ac25578270cff1b9f619").unwrap();
    pub static ref USDC_ADDRESS_POLYGON: H160 =
        H160::from_str("0x2791bca1f2de4661ed88a30c99a7a9449aa84174").unwrap();

    // Arbitrum
    pub static ref DAI_ADDRESS_ARBITRUM: H160 =
        H160::from_str("0xda10009cbd5d07dd0cecc66161fc93d7c9000da1").unwrap();
    pub static ref WETH_ADDRESS_ARBITRUM: H160 =
        H160::from_str("0x82af49447d8a07e3bd95bd0d56f35241523fbab1").unwrap();
    pub static ref USDC_ADDRESS_ARBITRUM: H160 =
        H160::from_str("0xaf88d065e77c8cc2239327c5edb3a432268e5831").unwrap();

    // Sepolia
    pub static ref DAI_ADDRESS_SEPOLIA: H160 =
        H160::from_str("0xfbb7e7fee1525958bf5a4f04ed8d7be547ab6d27").unwrap();
    pub static ref WETH_ADDRESS_SEPOLIA: H160 =
        H160::from_str("0x4200000000000000000000000000000000000006").unwrap();
    pub static ref USDC_ADDRESS_SEPOLIA: H160 =
        H160::from_str("0x4cc7ebeed5ea3adf3978f19833d2e1f3e8980cd6").unwrap();
}
