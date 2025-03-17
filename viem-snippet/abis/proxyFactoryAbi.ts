export const proxyFactoryAbi = [
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": true,
        "internalType": "address",
        "name": "proxy",
        "type": "address"
      },
      {
        "indexed": false,
        "internalType": "address",
        "name": "singleton",
        "type": "address"
      }
    ],
    "name": "ProxyCreation",
    "type": "event"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "singleton",
        "type": "address"
      },
      {
        "internalType": "bytes",
        "name": "data",
        "type": "bytes"
      }
    ],
    "name": "createProxy",
    "outputs": [
      {
        "internalType": "address",
        "name": "proxy",
        "type": "address"
      }
    ],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "_singleton",
        "type": "address"
      },
      {
        "internalType": "bytes",
        "name": "initializer",
        "type": "bytes"
      },
      {
        "internalType": "uint256",
        "name": "saltNonce",
        "type": "uint256"
      }
    ],
    "name": "createProxyWithNonce",
    "outputs": [
      {
        "internalType": "address",
        "name": "proxy",
        "type": "address"
      }
    ],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "_singleton",
        "type": "address"
      },
      {
        "internalType": "bytes",
        "name": "initializer",
        "type": "bytes"
      },
      {
        "internalType": "uint256",
        "name": "saltNonce",
        "type": "uint256"
      }
    ],
    "name": "calculateCreateProxyWithNonceAddress",
    "outputs": [
      {
        "internalType": "address",
        "name": "proxy",
        "type": "address"
      }
    ],
    "stateMutability": "nonpayable",
    "type": "function"
  }
] as const; 