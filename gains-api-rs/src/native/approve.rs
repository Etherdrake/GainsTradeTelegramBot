use std::error::Error;
use std::str::FromStr;
use std::sync::Arc;
use ethcontract::common::abi::ethereum_types::U64;
use ethers::abi::AbiEncode;
use ethers::contract::abigen;
use ethers::middleware::Middleware;
use ethers::prelude::{Address, Eip1559TransactionRequest, Http, LocalWallet, NameOrAddress, Provider, Signer, SignerMiddleware, TransactionReceipt};
use ethers::prelude::transaction::eip2718::TypedTransaction;
use ethers::prelude::transaction::eip2930::AccessList;
use ethers::utils::parse_units;
use log::error;
use primitive_types::{H160, U256};
use crate::database::getkeys::get_priv_key_by_guid;
use crate::gains::utils::getchainparams::get_chain_params;
use crate::gains::utils::getgasparams::get_gas_params;

pub async fn approve(
    guid: i64,
    owner: String,
    token: String,
    spender: String,
    amount: String,
    collateral: u8,
    chain: u64,
) -> Result<String, Box<dyn Error>> {
    let (trading_contract, chain_id, uri) = match get_chain_params(chain, collateral) {
        Ok((contract, id, uri_str)) => (contract, id, uri_str),
        Err(e) => return Err(e.into()),
    };

    let provider = Provider::<Http>::try_from(uri)?;

    let priv_key_from_db = get_priv_key_by_guid(&guid).await?;
    let priv_key_str = priv_key_from_db.ok_or_else(|| {
        let error_msg = "No private key found for guid.";
        error!("{}", error_msg);
        std::io::Error::new(std::io::ErrorKind::Other, error_msg)
    })?;

    let wallet: LocalWallet = priv_key_str.parse()?;
    let client = Arc::new(SignerMiddleware::new(provider, wallet.clone().with_chain_id(chain_id)));

    abigen!(
        ERC20,
        "./abi/ERC20.abi.json",
        event_derives(serde::Deserialize, serde::Serialize)
    );

    let spender: H160 = Address::from_str(&*spender).unwrap();
    let token: H160 = Address::from_str(&*token).unwrap();
    let amount: U256 = U256::from_dec_str(&amount).unwrap();


    let call: ApproveCall = ApproveCall {
        spender,
        value: amount,
    };

    let nonce: U256 = client.get_transaction_count(client.address(), None).await?;

    let (
        gas_fee,
        max_prio_fee,
        max_fee_per_gas) = get_gas_params(chain).await.unwrap();

    let tx: Eip1559TransactionRequest = Eip1559TransactionRequest {
        from: Some(client.address()),
        to: Some(NameOrAddress::Address(token)),
        gas: Some(gas_fee),
        value: None,
        data: Some(call.encode().into()),
        nonce: Some(nonce),
        access_list: AccessList::default(),
        max_priority_fee_per_gas: Some(max_prio_fee),
        max_fee_per_gas: Some(max_fee_per_gas),
        chain_id: Some(U64::from(chain_id)),
    };

    let tx_typed: TypedTransaction = TypedTransaction::Eip1559(tx);

    let pending_tx: Option<TransactionReceipt> = client
        .send_transaction(tx_typed, None)
        .await
        .map_err(|err| {
            error!("Error sending transaction: {}", err);
            err
        })?
        .await
        .map_err(|err| {
            error!("Error sending transaction: {}", err);
            err
        })?;

    let receipt = pending_tx.unwrap();
    let tx_receipt = format!("{:?}", receipt.transaction_hash);
    let tx_status = receipt.status.unwrap();

    if tx_status == U64::from(1) {
        println!(
            "Approve Successful - Receipt: {}", tx_receipt
        );
        Ok(tx_receipt)
    } else {
        println!("Approve Failed  - Receipt: {}", tx_receipt);
        Err(Box::from(tx_receipt))
    }
}
