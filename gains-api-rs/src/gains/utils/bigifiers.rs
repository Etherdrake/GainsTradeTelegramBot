use ethers::types::Address;

use ethers::prelude::*;

use crate::gains::types::bigtypes::{
    CloseTradeMarketBig,
    OpenTradeBig,
    UpdateOpenLimitOrderBig,
    UpdateStopLossBig,
    UpdateTakeProfitBig,
};

use ethcontract::{H160, U256};
use ethers::utils::{parse_units, ParseUnits};
use std::error::Error;

// pub fn transform_to_big(open_trade_json: &OpenTradeJSON) -> Result<OpenTradeBig, Box<dyn Error>> {
//     let pos_size_dai: U256;
//     if open_trade_json.collateral_token == "USDC" {
//         let pu_pos_size_dai: ParseUnits =
//             parse_units(open_trade_json.position_size_dai.clone(), 6).unwrap();
//         pos_size_dai = U256::from(pu_pos_size_dai)
//     } else {
//         let pu_pos_size_dai: ParseUnits =
//             parse_units(open_trade_json.position_size_dai.clone(), 18).unwrap();
//         pos_size_dai = U256::from(pu_pos_size_dai)
//     }
//
//     let pu_open: ParseUnits = parse_units(open_trade_json.open_price.clone(), 10).unwrap();
//     let pu_tp: ParseUnits = parse_units(open_trade_json.tp.clone(), 10).unwrap();
//     let pu_sl: ParseUnits = parse_units(open_trade_json.sl.clone(), 10).unwrap();
//
//     let open_parsed = U256::from(pu_open);
//     let tp_parsed = U256::from(pu_tp);
//     let sl_parsed = U256::from(pu_sl);
//
//     // Return an Ok variant with OpenTradeBig directly
//     Ok(OpenTradeBig {
//         _id: open_trade_json._id,
//         trader: Address::default(),
//         trade_id: open_trade_json.trade_id.clone(),
//         paper: true,
//         pair_index: U256::from_dec_str(&open_trade_json.pair_index).map_err(
//             |_| -> Box<dyn Error> {
//                 format!("error parsing pair_index: {}", &open_trade_json.pair_index).into()
//             },
//         )?,
//         index: U256::from_dec_str(&open_trade_json.index).map_err(|_| -> Box<dyn Error> {
//             format!("error parsing index: {}", &open_trade_json.index).into()
//         })?,
//         initial_pos_token: U256::from_dec_str(&open_trade_json.initial_pos_token).map_err(
//             |_| -> Box<dyn Error> {
//                 format!(
//                     "error parsing initial_pos_token: {}",
//                     &open_trade_json.initial_pos_token
//                 )
//                     .into()
//             },
//         )?,
//         position_size_dai: pos_size_dai,
//         open_price: open_parsed,
//         close_price: U256::from_dec_str("420").map_err(|_| -> Box<dyn Error> {
//             format!(
//                 "error parsing close_price: {}",
//                 &open_trade_json.close_price
//             )
//                 .into()
//         })?,
//         buy: open_trade_json.buy,
//         leverage: U256::from_dec_str(&open_trade_json.leverage).map_err(|_| -> Box<dyn Error> {
//             format!("error parsing leverage: {}", &open_trade_json.leverage).into()
//         })?,
//         tp: tp_parsed,
//         sl: sl_parsed,
//         liq: U256::from_dec_str(&open_trade_json.liq).map_err(|_| -> Box<dyn Error> {
//             format!("error parsing liq: {}", &open_trade_json.liq).into()
//         })?,
//         percentage_tp: U256::from_dec_str("420").map_err(|_| -> Box<dyn Error> {
//             format!(
//                 "error parsing percentage_tp: {}",
//                 &open_trade_json.percentage_tp
//             )
//                 .into()
//         })?,
//         percentage_sl: U256::from_dec_str("420").map_err(|_| -> Box<dyn Error> {
//             format!(
//                 "error parsing percentage_sl: {}",
//                 &open_trade_json.percentage_sl
//             )
//                 .into()
//         })?,
//         order_type: U256::from(open_trade_json.order_type),
//         order_status: U256::from(open_trade_json.order_status),
//         pnl: U256::from_dec_str("420").map_err(|_| -> Box<dyn Error> {
//             format!("error parsing pnl: {}", &open_trade_json.pnl).into()
//         })?,
//         chain: open_trade_json.chain.clone(),
//         timestamp: U256::from(open_trade_json.timestamp),
//     })
// }
//
// pub fn bigify_close_trade_market(
//     close_trade_market_struct: CloseTradeMarketJSON,
// ) -> Result<CloseTradeMarketBig, Box<dyn Error>> {
//     let id = close_trade_market_struct._id;
//     let pair_index = close_trade_market_struct.pair_index;
//     let index = close_trade_market_struct.index;
//
//     let close_trade_market_big = CloseTradeMarketBig {
//         _id: id,
//         pair_index: U256::from_dec_str(&pair_index).map_err(|_| "error parsing pair_index")?,
//         index: U256::from_dec_str(&index).map_err(|_| "error parsing index")?,
//     };
//
//     Ok(close_trade_market_big)
// }
//
// pub fn bigify_update_sl(
//     update_sl_struct: UpdateStopLossJSON,
// ) -> Result<UpdateStopLossBig, Box<dyn Error>> {
//     let id = update_sl_struct._id;
//     let pair_index = update_sl_struct.pair_index;
//     let index = update_sl_struct.index;
//
//     let pu_sl: ParseUnits = parse_units(update_sl_struct.new_sl.clone(), 10).unwrap();
//     let sl_parsed = U256::from(pu_sl);
//
//     let update_sl_big = UpdateStopLossBig {
//         _id: id,
//         pair_index: U256::from_dec_str(&pair_index).map_err(|_| "error parsing pair_index")?,
//         index: U256::from_dec_str(&index).map_err(|_| "error parsing index")?,
//         new_sl: sl_parsed,
//     };
//     Ok(update_sl_big)
// }
//
// pub fn bigify_update_tp(
//     update_tp_struct: UpdateTakeProfitJSON,
// ) -> Result<UpdateTakeProfitBig, Box<dyn Error>> {
//     let id = update_tp_struct._id;
//     let pair_index = update_tp_struct.pair_index;
//     let index = update_tp_struct.index;
//
//     let tp_sl: ParseUnits = parse_units(update_tp_struct.new_tp.clone(), 10).unwrap();
//     let tp_parsed = U256::from(tp_sl);
//
//     let update_tp_big = UpdateTakeProfitBig {
//         _id: id,
//         pair_index: U256::from_dec_str(&pair_index).map_err(|_| "error parsing pair_index")?,
//         index: U256::from_dec_str(&index).map_err(|_| "error parsing index")?,
//         new_tp: tp_parsed,
//     };
//
//     Ok(update_tp_big)
// }
//
// pub fn bigify_update_open_limit_order(
//     update_open_limit_order_struct: UpdateOpenLimitOrderJSON,
// ) -> Result<UpdateOpenLimitOrderBig, Box<dyn Error>> {
//     let id = update_open_limit_order_struct._id;
//     let index = update_open_limit_order_struct.index;
//     let pair_index = update_open_limit_order_struct.pair_index;
//
//     let pu_price: ParseUnits =
//         parse_units(update_open_limit_order_struct.price.clone(), 10).unwrap();
//     let pu_tp: ParseUnits = parse_units(update_open_limit_order_struct.tp.clone(), 10).unwrap();
//     let pu_sl: ParseUnits = parse_units(update_open_limit_order_struct.sl.clone(), 10).unwrap();
//
//     let price_parsed = U256::from(pu_price);
//     let tp_parsed = U256::from(pu_sl);
//     let sl_parsed = U256::from(pu_tp);
//
//     let update_open_limit_order_big = UpdateOpenLimitOrderBig {
//         _id: id,
//         pair_index: U256::from_dec_str(&pair_index).map_err(|_| "error parsing pair_index")?,
//         index: U256::from_dec_str(&index).map_err(|_| "error parsing index")?,
//         price: price_parsed,
//         tp: tp_parsed,
//         sl: sl_parsed,
//     };
//     Ok(update_open_limit_order_big)
// }
