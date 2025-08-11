use std::str::FromStr;
use anyhow::{Context, Result};
use ethers::{prelude::*, types::transaction::eip2718::TypedTransaction};
use std::sync::Arc;
use ethers::abi::AbiEncode;
use ethers::types::transaction::eip2930::AccessList;
use crate::database::getkeys::get_pub_key_by_guid;
use crate::{
    database::getkeys::get_priv_key_by_guid
};
use crate::gains::helpers::oracle::provider_oracle;
use crate::gains::tradinginteractions::write::bigtypes::TradePayloadBig;
use crate::gains::utils::getchainparams::get_chain_params;
use crate::gains::utils::getgasparams::get_gas_params;

pub async fn open_trade(
    guid: i64,
    trade: TradePayloadBig,
    max_slippage_p: u16,
    referrer: Address,
    collateral_token: u8,
    chain: u64,
) -> Result<String, Box<dyn std::error::Error>> {
    let mut trading_contract: H160 = Default::default();
    let mut chain_id: u64 = 0;
    let mut uri: &str = "";

    match get_chain_params(chain, collateral_token) {
        Ok((contract, id, uri_str)) => {
            trading_contract = contract;
            chain_id = id;
            uri = uri_str;
        }
        Err(e) => return Err(e),
    }
    // This is the wallet PrivateKey
    let priv_key_from_db = get_priv_key_by_guid(&guid).await?;
    match priv_key_from_db {
        Some(priv_key_str) => {
            let wallet: LocalWallet = priv_key_str.parse()?;
            let provider = Provider::<Http>::try_from(uri)?;
            let clone_provider = provider.clone();
            let third_provider = clone_provider.clone();
            let balance = provider.get_balance(wallet.address(), None).await?;
            let client = Arc::new(SignerMiddleware::new(
                provider,
                wallet.clone().with_chain_id(chain_id),
            ));

            abigen!(GNSTradingInteractions, "./gains_abi/facets/GNSTradingInteractions/GNSTradingInteractions_abi.json");

            let order_type_cloned = trade.trade_type.clone();

            let pub_key = get_pub_key_by_guid(&guid).await?;
            match pub_key {
                Some(pub_key_str) => {
                    let zero_index = U256::from_dec_str("0").unwrap();
                    let call: OpenTradeCall = OpenTradeCall {
                        trade: Trade {
                            user: (trade.user),
                            pair_index: (trade.pair_index),
                            index: (trade.index),
                            open_price: (trade.open_price),
                            leverage: trade.leverage,
                            long: trade.long,
                            is_open: false,
                            collateral_index: trade.collateral_index,
                            trade_type: trade.trade_type,
                            tp: trade.tp,
                            sl: trade.sl,
                            collateral_amount: trade.collateral_amount,
                            placeholder: Default::default(),
                        },
                        max_slippage_p: u16::from_str("2000").unwrap(),
                        referrer: Address::from_str("0xBc9B019AA5885Ba50F3770B6Fd7E5DA981007068").unwrap(),
                    };

                    let nonce: U256 = client.get_transaction_count(client.address(), None).await?;

                    let trading_contract_address = trading_contract;

                    let gas_fee: U256;
                    let max_prio_fee: U256;
                    let max_fee_per_gas: U256;
                    let gas_price = provider_oracle(chain).await;

                    let (
                        gas_fee,
                        max_prio_fee,
                        max_fee_per_gas) = get_gas_params(chain).await.unwrap();

                    let tx: Eip1559TransactionRequest = Eip1559TransactionRequest {
                        from: Some(client.address()),
                        to: Some(NameOrAddress::Address(trading_contract_address)),
                        gas: Option::from(Some(U256::from(gas_fee)).unwrap()),
                        value: None,
                        data: Some(call.encode().into()),
                        nonce: Some(nonce),
                        access_list: AccessList::default(),
                        // max_priority_fee_per_gas: Some(max_prio_fee),
                        // max_fee_per_gas: Some(max_fee_per_gas),  // Some(gas_price)
                        max_priority_fee_per_gas: Option::from(max_prio_fee),
                        max_fee_per_gas: Option::from(max_fee_per_gas),
                        chain_id: Some(U64::from(chain_id)),
                    };

                    let tx_typed: TypedTransaction = TypedTransaction::Eip1559(tx);

                    let pending_tx: Option<TransactionReceipt> = client
                        .send_transaction(tx_typed, None)
                        .await
                        .map_err(|err| {
                            println!("Error sending update: {}", err);
                            err
                        })?
                        .await
                        .map_err(|err| {
                            println!("Error sending update: {}", err);
                            err
                        })?;

                    let receipt = pending_tx.unwrap();
                    let tx_receipt = format!("{:?}", receipt.transaction_hash);
                    let tx_status = receipt.status.unwrap();

                    if tx_status == U64::from(1) {
                        println!("OpenTrade Successful - Receipt: {}", tx_receipt);
                        Ok(tx_receipt)
                    } else {
                        println!("OpenTrade Failed - Receipt: {}", tx_receipt);
                        Err(Box::from(tx_receipt))
                    }
                }
                None => Err(format!("No publickey found for guid: {}", guid).into()),
            }
        }
        None => Err(format!("No private key found for guid: {}", guid).into()),
    }
}
