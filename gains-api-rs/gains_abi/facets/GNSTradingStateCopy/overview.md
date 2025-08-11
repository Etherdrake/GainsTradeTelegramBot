1. ## Read Functions

1. ### getAddresses
   #### Description
   Returns the addresses stored in the contract.
   #### Parameters
   None
   #### Returns
    - **gns**: Address of the GNS contract.
    - **gnsStaking**: Address of the GNS Staking contract.
   #### State Mutability
   View

2. ### getCollateralStageState
   #### Description
   Returns the stage state for a specific collateral index.
   #### Parameters
    - **_collateralIndex**: Index of the collateral.
   #### Returns
    - **bool**: Stage state.
   #### State Mutability
   View

3. ### getCollateralState
   #### Description
   Returns the state of a specific collateral.
   #### Parameters
    - **_collateralIndex**: Index of the collateral.
   #### Returns
    - **COPY_STATE**: State of the collateral.
    - **uint256**: Balance of the collateral.
    - **uint16**: Index of the collateral.
   #### State Mutability
   View

4. ### hasRole
   #### Description
   Checks if an account has a specific role.
   #### Parameters
    - **_account**: Address of the account.
    - **_role**: Role to check.
   #### Returns
    - **bool**: Whether the account has the role.
   #### State Mutability
   View

2. ## Write Functions

1. ### initialize
   #### Description
   Initializes the contract.
   #### Parameters
    - **_rolesManager**: Address of the roles manager contract.
   #### Returns
   None
   #### State Mutability
   Non-payable

2. ### markAsDone
   #### Description
   Marks a specific collateral as done.
   #### Parameters
    - **_collateralIndex**: Index of the collateral to mark.
   #### Returns
   None
   #### State Mutability
   Non-payable

3. ### setRoles
   #### Description
   Sets roles for multiple accounts.
   #### Parameters
    - **_accounts**: Array of account addresses.
    - **_roles**: Array of role values.
    - **_values**: Array of boolean values.
   #### Returns
   None
   #### State Mutability
   Non-payable

4. ### transferBalance
   #### Description
   Transfers the balance of a specific collateral.
   #### Parameters
    - **_collateralIndex**: Index of the collateral to transfer the balance from.
   #### Returns
   None
   #### State Mutability
   Non-payable
