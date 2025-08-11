use axum::http::StatusCode;
use axum::Json;
use axum::response::IntoResponse;
use serde::Deserialize;
use crate::errors::ErrorMessageResponse;
use crate::gains::errors::{ChainError, CollateralError};
use crate::gains::tradingstorage::read::getalltradeinfos::get_all_trade_infos;



pub async fn handle_get_all_trade_infos(
    Json(payload): Json<crate::routes::tradingstorage::payloads::GetAllTradeInfosPayload>,
) -> impl IntoResponse {
    match get_all_trade_infos(
        payload.offset,
        payload.limit,
        payload.collateral_token,
        payload.chain,
    ).await {
        Ok(tradeinfos) => (
            StatusCode::OK,
            Json(Ok(tradeinfos)),
        ),
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
