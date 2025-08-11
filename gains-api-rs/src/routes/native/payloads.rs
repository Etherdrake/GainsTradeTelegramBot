use serde::{Deserialize};

#[derive(Deserialize)]
pub struct ApprovePayload {
    pub guid: i64,
    pub owner: String,
    pub token: String,
    pub spender: String,
    pub amount: String,
    pub collateral :u8,
    pub chain: u64,
}

#[derive(Deserialize)]
pub struct BalanceOfPayload {
    pub public_key: String,
    pub collateral: u8,
    pub chain: u64,
}

#[derive(Deserialize)]
pub struct GetGasBalancePayload {
    pub public_key: String,
    pub collateral: u8,
    pub chain: u64,
}

#[derive(Deserialize)]
pub struct GetAllowancePayload {
    pub owner: String,
    pub token: String,
    pub spender: String,
    pub collateral :u8,
    pub chain: u64,
}

#[derive(Deserialize)]
pub struct GetBlockTimePayload {
    pub block_number: u64,
    pub chain: u64,
}

