use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct GetFundsPayload {
    pub(crate) guid: i64,
}
