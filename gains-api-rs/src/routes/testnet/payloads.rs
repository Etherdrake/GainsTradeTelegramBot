use serde::{Deserialize, Serialize};

#[derive(Debug, Serialize, Deserialize)]
pub struct GetFreeDaiPayload {
    pub(crate) guid: i64,
    pub(crate) chain: u64,
}