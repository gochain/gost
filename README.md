# GOST Protocol

GOST is a protocol for executing cross-chain confirmations of events between
the GoChain and Ethereum blockchains. By leveraging the trust in GoChain's
network of reputable signers, contracts deployed on GoChain can securely 
and confidently confirm events from Ethereum logs, and vice-versa. This
simple primitive unlocks powerful interoperability which can be used to
build more complex solutions, like **cross-chain token transfers** of ERC20
and GO20 tokens.

https://medium.com/gochain/gost-protocol-cross-chain-transactions-between-ethereum-and-gochain-39180811301d

## Status

This project is still in development.

1. **Alpha (devnets)**
2. Beta (testnets)
3. Production (mainnets)

## How it Works

The GOST protocol relies on GoChain's network of reputable signers to 
monitor and manage a `Confirmations` contract on each network. Users
(humans or other contracts) request confirmation of events by 
specifying *block number*, *log index*, and an *event hash* - which 
uniquely identifies a particular event. The GoChain signers monitor
for requested events, and once the block has been confirmed as canonical
they vote to confirm an event as either *valid* or *invalid* with 
majority consensus.

### Cross-Chain Confirmations

The `IConfirmations` interface defines the public set of functions available
to users, and the `Confirmations` contract implements and extends the set
with administrative functions. Event hashes are content based identifiers 
derived from the source contract address which emitted the event and the
event parameters (topics and data). 

```solidity 
// Event hash pseudocode. Note that the first topic will be the event ID.
hash = keccak256(abi.encode(source, topics..., data))
```

This allows contracts to not only confirm a particular event, but also 
to operate on the *values of the event's parameters*. In the case of token
transfers, the event can specify the amount and recipient address for the
receiving contract code to use. 

### Cross-Chain Token Transfers

One application of cross-chain event confirmations is the ability to manage
cross-chain token transfers. By confirming events emitted when tokens are
frozen or burned on one network, we can securely unfreeze or mint tokens on
the other network. The `Transfers` contract defines a `GOSTTransfer` event, 
and provides a function to compute event hashes:
  
```solidity
event GOSTTransfer(address indexed addr, uint amount);

// Compute the event hash for a GOSTTransfer event from the source contract.
function transferEventHash(address addr, uint amount) public view returns (bytes32) {
    return keccak256(abi.encode(transferSourceContract, eventID, addr, amount));
}
```

The `Custody` and `Proxy` contracts use these building blocks to securely transfer
tokens back and forth between networks.
The `Custody` contract deployed to the origin network (where the `ERC20`/`GO20` token 
is deployed) and the `Proxy` contract deployed to the other network both emit and
confirm each others' `GOSTTransfer` events. The custody contract holds locked
tokens which have been sent to the proxy network, and only unlocks them when
return transfer events from the proxy network are confirmed. Similarly, the proxy 
contract mints tokens when incoming transfers are confirmed and burns them when
outgoing transfers are initiated.

#### Origin to Proxy

How to transfer tokens from their origin network to the proxy network:

1. `token.approve()`: Approve the custody contract as a spender for the amount 
being transferred.
2. `custody.proxyTransfer()`: Initiate the transfer by freezing tokens in the 
custody contract and emitting a `GOSTTransfer` event. 
3. `proxyConfirmations.request()`: Request confirmation of the event and 
deposit funds to cover confirmation gas fees. 
4. `proxyConfirmations.status()`: Wait for confirmation from the signers that 
the event is valid (status `Confirmed`).
5. `proxy.claimTransfer()`: Claim the confirmed transfer to receive minted tokens.

#### Proxy to Origin 

How to transfer tokens from the proxy network back to the origin network:

1. `proxy.originTransfer()`: Initiate the transfer by burning tokens in the proxy
contract and emitting a `GOSTTransfer` event.
2. `originConfirmations.request()`: Request confirmations of the event and 
deposit funds to cover confirmations gas fees.
3. `originConfirmations.status()`: Wait for confirmations from the signers that
the event is valid (status `Confirmed`).
4. `origin.claimTransfer()`: Claim the confirmed transfer to receive unlocked tokens. 

## Contracts

TODO confirmations
TODO proxy and custody

### Deployment

TODO deployment steps for proxy and custody contracts (just link command or script?)

### Confirmations

Contract addresses for the `Confirmations` contracts manages by the GoChain signers:

#### Mainnet

| Network  | Address |
| -------- | ------- |
| GoChain  | TBD |
| Ethereum | TBD |

#### Testnet

| Network  | Address |
| -------- | ------- |
| GoChain  | TBD |
| Ropsten | TBD |

### Tokens

Contract addresses for cross-chain tokens:

| Token  | Origin | Address | Custody | Proxy |
| ------ | ------ | ------- | ------- | ----- | 
| SampleToken  | Ethereum | 0x0 | 0x1 | 0x2 |
