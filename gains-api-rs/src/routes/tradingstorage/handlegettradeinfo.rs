use axum::http::StatusCode;
use axum::Json;
use ethers::core::k256::U256;
use ethers::prelude::I256;
use num_format::Locale::pa;
use serde::Deserialize;
use crate::errors::ErrorMessageResponse;
use crate::gains::errors::{ChainError, CollateralError};
use crate::gains::tradingstorage::read::gettradeinfo::get_trade_info;
use crate::gains::tradingstorage::tradingstoragetypes::trading_storage::{PendingOrderStruct, TradeInfoStruct};

pub async fn handle_get_trade_info(
    Json(payload): Json<crate::routes::tradingstorage::payloads::GetTradeInfoPayload>,
) -> (StatusCode, Json<Result<TradeInfoStruct, ErrorMessageResponse>>) {
    match get_trade_info(
        payload.trader,
        payload.index,
        payload.collateral_token,
        payload.chain,
    ).await {
        Ok(trade_info) => (StatusCode::OK, Json(Ok(trade_info))),
        Err(e) => {
            let error_message = match e.downcast_ref::<CollateralError>() {
                Some(collateral_error) => format!("Collateral error: {}", collateral_error),
                None => match e.downcast_ref::<ChainError>() {
                    Some(chain_error) => format!("Chain error: {}", chain_error),
                    None => format!("An unknown error occurred: {}", e),
                },
            };
            (
                StatusCode::BAD_REQUEST,
                Json(Err(ErrorMessageResponse { message: error_message })),
            )
        }
    }
}
