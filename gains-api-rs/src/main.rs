#![allow(warnings)]

use std::{env, fmt};
use std::convert::Infallible;
use std::error::Error;
use ethers::prelude::{Http, Provider};
use serde_json::json;

use hoot_core_api::routes::tradingstorage::*;

use axum::{routing::get, Router, Json};
use axum::extract::Path;
use axum::response::IntoResponse;
use axum::routing::post;
use dotenv::dotenv;

use hoot_core_api::routes::native::{
    handlebalanceof::handle_balance_of,
    handlegetallowance::handle_get_allowance,
    handlegetgasbalance::handle_get_gas_balance,
};
use hoot_core_api::routes::testnet::{
    handlegetfreedai::handle_get_free_dai,
    handlegetfreeusdc::handle_get_free_usdc,
};
use hoot_core_api::routes::tradinginteractions::{
    handlecancelopenorder::handle_cancel_open_order,
    handleclosetrademarket::handle_close_trade_market,
    handleopentrade::handle_open_trade,
    handleupdateopenorder::handle_update_open_order,
};
use hoot_core_api::routes::tradingstorage::{
    handlegetallpendingorders::handle_get_all_pending_orders,
    handlegetalltradeinfos::handle_get_all_trade_infos,
    handlegetalltrades::handle_get_all_trades,
    handlegetcollaterals::handle_get_collaterals,
    handlegetpendingorder::handle_get_pending_order,
    handlegetpnlpercent::handle_get_pnl_percent,
    handlegettradeinfo::handle_get_trade_info,
    handlegettradeinfos::handle_get_trade_infos,
    handlegettraders::handle_get_traders,
};

use hoot_core_api::routes::tradinginteractions::handleupdatesl::handle_update_sl;
use hoot_core_api::routes::tradinginteractions::handleupdatetp::handle_update_tp;

use env_logger;
use log::info;
use hoot_core_api::routes::native::handleapprove::handle_approve;
use hoot_core_api::routes::pairsstorage::handlegetpairmaxleverage::handle_get_pair_max_leverage;
use hoot_core_api::routes::pairsstorage::handlegetpairminleverage::handle_get_pair_min_leverage;
use hoot_core_api::routes::pairsstorage::handlegetpairminpositionsizeusd::handle_get_pair_min_position_size_usd;
use hoot_core_api::routes::proprietary::handlegetethsepolia::handle_get_eth_sepolia;
use hoot_core_api::routes::proprietary::handlegetfunds::handle_get_funds;

use hoot_core_api::routes::native::handlegetblocktime::handle_get_block_time;

const LOCAL_URI: &str = "http://localhost:8545";
const MONGO_URI: &str = "mongodb://localhost:55000/";

async fn health_check() -> Result<Json<&'static str>, Infallible> {
    Ok(Json("Service is up and running"))
}

async fn get_all_pending_orders_test(Path((offset, limit, collateral, chain)): Path<(u64, u64, String, String)>) -> Result<Json<String>, Infallible> {
    // Your implementation here
    Ok(Json(format!("Offset: {}, Limit: {}, Collateral: {}, Chain: {}", offset, limit, collateral, chain)))
}

#[tokio::main]
async fn main() {
    // Load environment variables from .env file
    dotenv().ok();

    env::set_var("RUST_BACKTRACE", "1");

    // Initialize the logger
    env_logger::init();

    let app = create_routers();

    // Retrieve the URIs from environment variables
    let polygon_uri = env::var("POLYGON_URI").expect("POLYGON_URI not set");
    let arbitrum_uri = env::var("ARBITRUM_URI").expect("ARBITRUM_URI not ");

    let polygon_provider = Provider::<Http>::try_from(polygon_uri)
        .expect("Failed to create Polygon provider");

    let arbitrum_provider = Provider::<Http>::try_from(arbitrum_uri)
        .expect("Failed to create Arbitrum provider");

    info!("Server is up and running on port 3030");

    let listener = tokio::net::TcpListener::bind("0.0.0.0:3030").await.unwrap();
    axum::serve(listener, app).await.unwrap();
}

pub fn create_routers() -> Router {
    info!("Creating routers...");
    Router::new()
        .route("/", get(|| async { "Hello, Scotland!" }))
        .route("/health", get(health_check))

        // Trading Callbacks Read Methods
        .route("/v2/getallpendingorders", get(handle_get_all_pending_orders))
        .route("/v2/getalltradesinfos", get(handle_get_all_trade_infos))
        .route("/v2/getalltrades", get(handle_get_all_trades))
        .route("/v2/getcollaterals", get(handle_get_collaterals))
        .route("/v2/getpendingorder", get(handle_get_pending_order))
        .route("/v2/getpendingorders", get(handle_get_all_pending_orders))
        .route("/v2/getpnlpercent", get(handle_get_pnl_percent))
        .route("/v2/gettradeinfo", get(handle_get_trade_info))
        .route("/v2/gettradesinfos", get(handle_get_trade_infos))
        .route("/v2/gettraders", get(handle_get_traders))

        // Trading Interactions Write Methods
        .route("/v2/cancelopenorder", post(handle_cancel_open_order))
        .route("/v2/closetrademarket", post(handle_close_trade_market))
        .route("/v2/opentrade", post(handle_open_trade))
        .route("/v2/updateopenorder", post(handle_update_open_order))
        .route("/v2/updatesl", post(handle_update_sl))
        .route("/v2/updatetp", post(handle_update_tp))

        // Pairs Storage Method
        .route("/v2/getpairmaxleverage", post(handle_get_pair_max_leverage))
        .route("/v2/getpairminleverage", post(handle_get_pair_min_leverage))
        .route("/v2/getpairminpossizeusd", post(handle_get_pair_min_position_size_usd))

        // Blockchain Native Methods
        .route("/v2/approve", post(handle_approve))
        .route("/v2/getallowance", post(handle_get_allowance))
        .route("/v2/getgasbalance", post(handle_get_gas_balance))
        .route("/v2/balanceof", post(handle_balance_of))
        .route("/v2/getblocktime", get(handle_get_block_time))

        // Testnet methods
        .route("/v2/getfreedai", post(handle_get_free_dai))
        .route("/v2/getfreeusdc", post(handle_get_free_usdc))

        // .route("/v2/getfunds", post(handle_get_funds))
        .route("/v2/getethsepolia", post(handle_get_eth_sepolia))
}