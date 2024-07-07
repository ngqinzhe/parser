# Eth Parser

Fetches transactions related to subscribed addresses from the most recent 1000 blocks

## How to run
Build and run go program. HTTP server will run on `localhost:3000`

User need to subscribe their addresses to `parser` service before they can try to get transactions. Trying to fetch transactions using unsubscribed addresses will return an error.

### Endpoints
#### Subscribe
`curl --header "Content-Type: application/json" --request POST --data '{"address":"{your addresss}"}' http://localhost:3000/subscribe`

#### Get Transactions
`curl --header "Content-Type: application/json" --request POST --data '{"address":"{your address}"}' http://localhost:3000/getTransactions`


## Improvements
##### TODO: Cannot find any implementation in Eth golib to support query `fromAddresses` transactions.
##### TODO: Hook up message queue to push notifications
##### TODO: Store requests/responses and queried transactions into db
##### TODO: Improve Unit Tests coverage