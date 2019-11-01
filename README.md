# ENS Go

A simple gadget of ens.

## Prepare

A web3 API url is required to communicate with ens contracts.
If you don't have one, consider applying to [infura.io](https://infura.io/) for a free one.

An infura api url looks like: `https://mainnet.infura.io/v3/{PROJECT_ID}`.

## Usage

* Query a batch of ens domains info:

    `ens-go query --api {API_URL} hello blockchain google ok`

* Query a batch of ens domains info from file:

    `ens-go query --api {API_URL} -f names.txt`

* Run a telegram robot with the ability to query ens domains info:

    `ens-go robot --api {API_URL} --token {ROBOT_TOKEN} -u user1 -u user2`

## Links

* [ens docs](https://docs.ens.domains)
* [ens github](https://github.com/ensdomains/ens)
* [wealdtech's go-ens](https://github.com/wealdtech/go-ens)