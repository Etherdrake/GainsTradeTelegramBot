use std::error::Error;
use std::fmt;
use std::fmt::{Display, Formatter};
use std::str::FromStr;
use std::sync::Arc;

#[derive(Debug)]
pub struct ChainError {
    pub(crate) message: String,
}
impl Error for ChainError {}
impl Display for ChainError {
    fn fmt(&self, f: &mut self::Formatter<'_>) -> std::fmt::Result {
        write!(f, "Chain Incorrect")
    }
}

#[derive(Debug)]
pub struct CollateralError {
    pub(crate) message: String,
}
impl Error for CollateralError {}
impl Display for CollateralError {
    fn fmt(&self, f: &mut self::Formatter<'_>) -> std::fmt::Result {
        write!(f, "Collateral Incorrect")
    }
}

#[derive(Debug)]
pub struct TradeIndexError {
    pub(crate) message: String,
}

impl Error for TradeIndexError {}
impl Display for TradeIndexError {
    fn fmt(&self, f: &mut self::Formatter<'_>) -> std::fmt::Result {
        write!(f, "Trade index incorrect")
    }
}

#[derive(Debug)]
pub struct ConvertError {
    pub(crate) message: String,
}
impl Error for crate::gains::errors::ConvertError {}
impl Display for crate::gains::errors::ConvertError {
    fn fmt(&self, f: &mut self::Formatter<'_>) -> std::fmt::Result {
        write!(f, "Error dividing U256")
    }
}

#[derive(Debug)]
pub struct WrappedError(pub Box<dyn std::error::Error + Send + Sync>);

impl fmt::Display for WrappedError {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "Error: {}", self.0)
    }
}

#[derive(Debug)]
pub struct WrappedContractError(pub Box<dyn std::error::Error + Send + Sync>);

impl fmt::Display for WrappedContractError {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "Error: {}", self.0)
    }
}
