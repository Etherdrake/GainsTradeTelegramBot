use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct GetPairMaxLeverage {
    pub(crate) index: u64,
    pub(crate) chain: u64,
}
#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct GetPairMinLeverage {
    pub(crate) index: u64,
    pub(crate) chain: u64,
}
#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct GetPairMinPosSizeUsd {
    pub(crate) index: u64,
    pub(crate) chain: u64,
}
