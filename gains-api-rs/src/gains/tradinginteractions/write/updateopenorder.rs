use crate::database::getkeys::get_priv_key_by_guid;
use crate::gains::errors::ChainError;
use ethcontract::common::abi::ethereum_types::U64;
use ethers::abi::AbiEncode;
use ethers::contract::abigen;
use ethers::middleware::SignerMiddleware;
use ethers::prelude::transaction::eip2718::TypedTransaction;
use ethers::prelude::transaction::eip2930::AccessList;
use ethers::prelude::{
    Eip1559TransactionRequest, Http, LocalWallet, Middleware, NameOrAddress, Provider, Signer,
    TransactionReceipt,
};
use ethers::utils::parse_units;
use primitive_types::{H160, U256};
use std::sync::Arc;
use crate::gains::utils::getchainparams::get_chain_params;
use crate::gains::utils::getgasparams::get_gas_params;

pub async fn update_open_order(
    guid: i64,
    index: u32,
    trigger_price: u64,
    tp: u64,
    sl: u64,
    max_slippage_p: u16,
    collateral_token: u8,
    chain: u64,
) -> anyhow::Result<String, Box<dyn std::error::Error>> {
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
            let client = Arc::new(SignerMiddleware::new(
                provider,
                wallet.clone().with_chain_id(chain_id),
            ));

            abigen!(GNSTradingInteractions, "./gains_abi/facets/GNSTradingInteractions/GNSTradingInteractions_abi.json");

            let call: UpdateOpenOrderCall = UpdateOpenOrderCall {
                index,
                trigger_price,
                tp,
                sl,
                max_slippage_p,

            };

            let nonce: U256 = client.get_transaction_count(client.address(), None).await?;

            let (
                gas_fee,
                max_prio_fee,
                max_fee_per_gas) = get_gas_params(chain).await.unwrap();

            let tx: Eip1559TransactionRequest = Eip1559TransactionRequest {
                from: Some(client.address()),
                to: Some(NameOrAddress::Address(trading_contract)),
                gas: Some(gas_fee),
                value: None,
                data: Some(call.encode().into()),
                nonce: Some(nonce),
                access_list: AccessList::default(),
                max_priority_fee_per_gas: Some(max_prio_fee),
                max_fee_per_gas: Some(max_fee_per_gas), // Some(gas_price)
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
                println!(
                    "Update Open Limit Order Successful - Receipt: {}",
                    tx_receipt
                );
                Ok(tx_receipt)
            } else {
                println!("Update Open Limit Order Failed - Receipt: {}", tx_receipt);
                Err(Box::from(tx_receipt))
            }
        }

        None => Err(Box::new(std::io::Error::new(
            std::io::ErrorKind::Other,
            format!("No private key found for guid."),
        ))),
    }
}
