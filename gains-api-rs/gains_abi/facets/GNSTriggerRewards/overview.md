### Read Functions:

1. **getAddresses**
    - **Description:** Retrieves stored addresses.
    - **Parameters:** None
    - **Returns:** Tuple containing addresses.
    - **State Mutability:** View

2. **getTriggerPendingRewardsGns**
    - **Description:** Retrieves pending trigger rewards for a specific oracle.
    - **Parameters:** `_oracle` (address)
    - **Returns:** Pending trigger rewards in GNS.
    - **State Mutability:** View

3. **getTriggerTimeoutBlocks**
    - **Description:** Retrieves trigger timeout blocks.
    - **Parameters:** None
    - **Returns:** Timeout blocks.
    - **State Mutability:** View

4. **hasActiveOrder**
    - **Description:** Checks if there is an active order for a specific block.
    - **Parameters:** `_orderBlock` (uint256)
    - **Returns:** Boolean indicating whether there is an active order.
    - **State Mutability:** View

5. **hasRole**
    - **Description:** Checks if an account has a specific role.
    - **Parameters:** `_account` (address), `_role` (uint8)
    - **Returns:** Boolean indicating whether the account has the role.
    - **State Mutability:** View

### Write Functions:

1. **claimPendingTriggerRewards**
    - **Description:** Claims pending trigger rewards for the caller.
    - **Parameters:** `_oracle` (address)
    - **Returns:** None
    - **State Mutability:** Non-payable

2. **distributeTriggerReward**
    - **Description:** Distributes trigger rewards to the caller.
    - **Parameters:** `_rewardGns` (uint256)
    - **Returns:** None
    - **State Mutability:** Non-payable

3. **initialize**
    - **Description:** Initializes the contract.
    - **Parameters:** None
    - **Returns:** None
    - **State Mutability:** Non-payable

4. **initializeTriggerRewards**
    - **Description:** Initializes trigger rewards with a specific timeout block.
    - **Parameters:** `_timeoutBlocks` (uint16)
    - **Returns:** None
    - **State Mutability:** Non-payable

5. **setRoles**
    - **Description:** Sets roles for multiple accounts.
    - **Parameters:** `_accounts` (address[]), `_roles` (uint8[]), `_values` (bool[])
    - **Returns:** None
    - **State Mutability:** Non-payable

6. **updateTriggerTimeoutBlocks**
    - **Description:** Updates trigger timeout blocks.
    - **Parameters:** `_timeoutBlocks` (uint16)
    - **Returns:** None
    - **State Mutability:** Non-payable
