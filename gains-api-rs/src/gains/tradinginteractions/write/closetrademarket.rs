use anyhow::{Context, Result};
use clap::Parser;
use ethers::{prelude::*, types::transaction::eip2718::TypedTransaction};
use std::error::Error;
use std::sync::Arc;
use ethers::abi::AbiEncode;
use ethers::types::transaction::eip2930::AccessList;
use ethers::utils::parse_units;
use crate::gains::helpers::oracle::provider_oracle;
use crate::{
    database::getkeys::get_priv_key_by_guid,
    gains::errors::ChainError,
};
use crate::gains::utils::getchainparams::get_chain_params;
use crate::gains::utils::getgasparams::get_gas_params;

pub async fn close_trade_market(
    guid: i64,
    index: u32,
    collateral_token: u8,
    chain: u64,
) -> Result<String, Box<dyn Error>> {
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
            let client = Arc::new(SignerMiddleware::new(
                provider,
                wallet.clone().with_chain_id(chain_id),
            ));

            abigen!(GNSTradingInteractions, "./gains_abi/facets/GNSTradingInteractions/GNSTradingInteractions_abi.json");

            let call: CloseTradeMarketCall = CloseTradeMarketCall {
                index,
            };

            let nonce: U256 = client.get_transaction_count(client.address(), None).await?;
            let trading_contract_address: H160 = trading_contract;

            let block_number = clone_provider
                .get_block_number()
                .await
                .expect("Failed to get block number");

            let block: Option<Block<Transaction>> = third_provider
                .get_block_with_txs(block_number)
                .await
                .expect("Failed to get block");

            let mut base_fee_clone: U256;

            if let Some(block) = block {
                // Access the baseFeePerGas from the block
                if let Some(base_fee) = block.next_block_base_fee() {
                    base_fee_clone = base_fee.clone();
                    println!("Base Fee: {:?}", base_fee);
                }
            } else {
                base_fee_clone = U256::from(parse_units("0.1", "gwei")?);
                println!("Block not found");
            }

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
                println!("Close Trade Market Successful - Receipt: {}", tx_receipt);
                Ok(tx_receipt)
            } else {
                println!("Close Trade Market Failed - Receipt: {}", tx_receipt);
                Err(Box::from(tx_receipt))
            }
        }
        None => Err(format!(
            "No private key found for guid: {}",
            guid
        ).into()),
    }
}
