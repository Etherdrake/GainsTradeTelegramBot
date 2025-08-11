use std::str::FromStr;
use axum::http::StatusCode;
use axum::Json;
use axum::response::IntoResponse;
use ethers::types::Address;
use log::error;
use teloxide::Bot;
use teloxide::prelude::Requester;
use teloxide::types::{ChatId, Recipient};
use crate::errors::ErrorMessageResponse;
use crate::gains::errors::{ChainError, CollateralError};
use crate::gains::tradinginteractions::write::bigtypes::TradePayloadBig;
use crate::gains::tradinginteractions::write::closetrademarket::close_trade_market;
use crate::gains::tradinginteractions::write::opentrade::open_trade;
use crate::telegram_constants::constants::HOOT_API_GROUP_ID;

pub async fn handle_open_trade(
    Json(payload): Json<crate::routes::tradinginteractions::tradinginteractionspayloads::OpenTradePayload>,
) -> impl IntoResponse {
    let max_slippage_p: u16 = 0;
    let referrer: Address = Address::from_str("0x4A22a6B0a7aF226d4b2E8182C852550f3F53bbEf").unwrap();

    let payload_clone = payload.clone();
    let trade_big: TradePayloadBig = TradePayloadBig::from(payload.trade);

    let msg_success: String = payload_clone.format_success();

    let chat_id_from_i64: ChatId = ChatId(HOOT_API_GROUP_ID);
    // bot.send_message(Recipient::from(chat_id_from_i64), msg_success).await.expect("Error sending Teloxide message");

    match open_trade(
        payload.guid,
        trade_big,
        max_slippage_p,
        referrer,
        payload.collateral_token,
        payload.chain,
    ).await {
        Ok(trade) => (
            StatusCode::OK,
            Json(Ok(trade)),
        ),
        Err(e) => {
            let error_message = match e.downcast_ref::<CollateralError>() {
                Some(collateral_error) => format!("Collateral error: {}", collateral_error),
                None => match e.downcast_ref::<ChainError>() {
                    Some(chain_error) => format!("Chain error: {}", chain_error),
                    None => format!("An unknown error occurred: {}", e),
                },
            };
            error!("Error in handle_open_trade: {}", error_message);
            (
                StatusCode::BAD_REQUEST,
                Json(Err(ErrorMessageResponse { message: error_message })),
            )
        }
    }
}
