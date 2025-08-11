### Read Functions

1. **getAddresses**
   - **Description:** Retrieves the addresses of certain contracts from the trading contract.
   - **Parameters:** None
   - **Returns:** A tuple containing addresses of certain contracts.

2. **getByPassTriggerLink**
   - **Description:** Retrieves whether a specific user has bypassed the trigger link.
   - **Parameters:** `_user` (address): The address of the user.
   - **Returns:** A boolean indicating whether the user has bypassed the trigger link.

3. **getMarketOrdersTimeoutBlocks**
   - **Description:** Retrieves the number of blocks after which market orders timeout.
   - **Parameters:** None
   - **Returns:** The number of blocks after which market orders timeout.

4. **getTradingDelegate**
   - **Description:** Retrieves the trading delegate assigned to a specific trader.
   - **Parameters:** `_trader` (address): The address of the trader.
   - **Returns:** The address of the trading delegate assigned to the trader.

5. **getWrappedNativeToken**
   - **Description:** Retrieves the address of the wrapped native token.
   - **Parameters:** None
   - **Returns:** The address of the wrapped native token.

6. **hasRole**
   - **Description:** Checks whether an account has a specific role.
   - **Parameters:**
     - `_account` (address): The address of the account to check.
     - `_role` (enum IAddressStore.Role): The role to check.
   - **Returns:** A boolean indicating whether the account has the specified role.

### Write Functions

1. **initialize**
   - **Description:** Initializes the trading contract.
   - **Parameters:** None
   - **Returns:** None

2. **initializeTrading**
   - **Description:** Initializes trading with specific parameters.
   - **Parameters:**
     - `_marketOrdersTimeoutBlocks` (uint16): Number of blocks after which market orders timeout.
     - `_usersByPassTriggerLink` (address[]): List of user addresses that bypass the trigger link.
   - **Returns:** None

3. **removeTradingDelegate**
   - **Description:** Removes the trading delegate.
   - **Parameters:** None
   - **Returns:** None

4. **setRoles**
   - **Description:** Sets roles for multiple accounts.
   - **Parameters:**
     - `_accounts` (address[]): List of account addresses.
     - `_roles` (enum IAddressStore.Role[]): List of role values.
     - `_values` (bool[]): List of boolean values indicating whether the role is assigned.
   - **Returns:** None

5. **setTradingDelegate**
   - **Description:** Sets the trading delegate for a specific trader.
   - **Parameters:** `_delegate` (address): The address of the trading delegate.
   - **Returns:** None

6. **triggerOrder**
   - **Description:** Triggers an order.
   - **Parameters:** `_packed` (uint256): Packed order data.
   - **Returns:** None

7. **updateByPassTriggerLink**
   - **Description:** Updates the bypass trigger link status for multiple users.
   - **Parameters:**
     - `_users` (address[]): List of user addresses.
     - `_shouldByPass` (bool[]): List of boolean values indicating whether to bypass the trigger link.
   - **Returns:** None

8. **updateMarketOrdersTimeoutBlocks**
   - **Description:** Updates the number of blocks after which market orders timeout.
   - **Parameters:** `_valueBlocks` (uint16): The new value for market orders timeout in blocks.
   - **Returns:** None

9. **updateOpenOrder**
   - **Description:** Updates an open order.
   - **Parameters:**
     - `_index` (uint32): The index of the open order.
     - `_triggerPrice` (uint64): The new trigger price.
     - `_tp` (uint64): The new take profit price.
     - `_sl` (uint64): The new stop-loss price.
     - `_maxSlippageP` (uint16): The new maximum slippage percentage.
   - **Returns:** None

10. **updateSl**
    - **Description:** Updates the stop-loss price of an open order.
    - **Parameters:**
      - `_index` (uint32): The index of the open order.
      - `_newSl` (uint64): The new stop-loss price.
    - **Returns:** None

11. **updateTp**
    - **Description:** Updates the take profit price of an open order.
    - **Parameters:**
      - `_index` (uint32): The index of the open order.
      - `_newTp` (uint64): The new take profit price.
    - **Returns:** None
```****