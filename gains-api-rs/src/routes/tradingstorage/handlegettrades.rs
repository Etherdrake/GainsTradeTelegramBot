use axum::http::StatusCode;
use axum::Json;
use crate::errors::ErrorMessageResponse;
use crate::gains::errors::{ChainError, CollateralError};
use crate::gains::tradingstorage::read::gettrades::get_trades;
use crate::gains::tradingstorage::tradingstoragetypes::trading_storage::TradeStruct;

pub async fn handle_get_trades(
    Json(payload): Json<crate::routes::tradingstorage::payloads::GetTradesPayload>,
) -> (StatusCode, Json<Result<Vec<TradeStruct>, ErrorMessageResponse>>) {
    match get_trades(
        payload.trader,
        payload.collateral_token,
        payload.chain,
    ).await {
        Ok(trades) => (StatusCode::OK, Json(Ok(trades))),
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
