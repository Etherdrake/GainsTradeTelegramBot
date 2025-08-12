use ethers::prelude;
use ethers::types::Address;
use std::collections::{HashMap, HashSet};
use std::sync::{Arc, RwLock};

pub fn remove_duplicates(input: &Vec<Vec<Address>>) -> Vec<Address> {
    let mut seen = HashSet::new();
    input.into_iter()
        .flatten()
        .filter(|addr| seen.insert(addr.clone())) // Double deref because of reference to a reference
        .cloned() // Clone the &H160 references to get H160 values
        .collect()
}

