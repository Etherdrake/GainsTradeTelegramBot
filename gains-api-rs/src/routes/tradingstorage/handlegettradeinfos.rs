use axum::http::StatusCode;
use axum::Json;
use crate::errors::ErrorMessageResponse;
use crate::gains::errors::{ChainError, CollateralError};
use crate::gains::tradingstorage::read::gettradeinfos::get_trades_infos;
use crate::gains::tradingstorage::tradingstoragetypes::trading_storage::{PendingOrderStruct, TradeInfoStruct};

pub async fn handle_get_trade_infos(
    Json(payload): Json<crate::routes::tradingstorage::payloads::GetTradeInfosPayload>,
) -> (StatusCode, Json<Result<Vec<TradeInfoStruct>, ErrorMessageResponse>>) {
    match get_trades_infos(
        payload.trader,
        payload.collateral_token,
        payload.chain,
    ).await {
        Ok(trades_infos) => (StatusCode::OK, Json(Ok(trades_infos))),
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
