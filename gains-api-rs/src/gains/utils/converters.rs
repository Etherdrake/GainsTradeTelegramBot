use ethers::types::U256;
use rust_decimal::Decimal;
use std::str::FromStr;
use web3_unit_converter::Unit;


// WARNING: Not all functions in this file are tested, some are deprecated or simpy incorrect.
pub fn to_ether_u128(wei: u128) -> f64 {
    wei as f64 / 1e18
}

pub fn wei_to_gwei(balance_wei: U256) -> U256 {
    // 1 gwei = 1e9 wei
    balance_wei / U256::from(1_000_000_000)
}

pub fn wei_to_ether(balance_wei: U256) -> U256 {
    // 1 gwei = 1e9 wei
    let ether = U256::from_dec_str("1000000000000000000").unwrap();
    balance_wei / ether
}

pub fn wei_to_ether_string(wei: U256) -> String {
    // Hardcode the ether factor as U256
    let ether_factor: U256 = U256::from_dec_str("1000000000000000000").unwrap();

    // Perform the division using U256
    let ether = wei / ether_factor;

    // Convert U256 to string
    ether.to_string()
}

pub fn wei_to_ether_decimals(wei: U256, decimals: usize) -> String {
    // Hardcode the ether factor as U256
    let ether_factor: U256 = U256::exp10(decimals);

    // Perform the division using U256
    let ether = wei / ether_factor;

    // Convert U256 to string
    ether.to_string()
}

pub fn wei_to_ether_string_convert(wei: U256) -> String {
    // Hardcode the ether factor as U256
    let wei_string: String = wei.to_string();
    let ether = Unit::Wei(&*wei_string).to_eth_str().unwrap();
    // Convert U256 to string
    ether
}

// Keeps two numbers as decimals
pub fn wei_to_float(balance_wei: U256) -> f64 {
    // 1 gwei = 1e9 wei
    let ether = U256::from_dec_str("10000000000000000").unwrap();
    let twodecimals = balance_wei / ether;
    twodecimals.to_f64_lossy() / 100.00
}

pub fn to_ether_u256(wei: U256) -> f64 {
    let ether_denominator: U256 = U256::exp10(18);
    let ether_value: U256 = wei / ether_denominator;

    // Convert the resulting U256 to f64
    // Note: This can lose precision for very large U256 values.
    ether_value.low_u64() as f64 + (ether_value >> 64).low_u64() as f64 * 1e-18
}

// Converts a u64 to U256
pub fn u64_to_u256(value: u64) -> U256 {
    U256::from(value)
}
