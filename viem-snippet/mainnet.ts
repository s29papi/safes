import {createPublicClient, createWalletClient, http, Address, TransactionReceipt, Log, defineChain, encodeFunctionData} from 'viem'
import {privateKeyToAccount} from 'viem/accounts'
import * as dotenv from 'dotenv'
import {safeAbi} from './abis/safeAbi'
import {proxyFactoryAbi} from './abis/proxyFactoryAbi'

export const g7Network = defineChain({
  id: 2187,
  name: 'G7 Network',
  nativeCurrency: { name: 'Game7 Token', symbol: 'G7', decimals: 18 },
  rpcUrls: {
    default: {
      http: ['https://mainnet-rpc.game7.io'],
    },
  },
  blockExplorers: {
    default: {
      name: 'G7 Network Explorer',
      url: 'https://mainnet.game7.io',
      apiUrl: 'https://mainnet.game7.io/api',
    },
  },
  mainnet: true,
})

dotenv.config()

const SAFE_SINGLETON_ADDRESS = '0x29fcb43b46531bca003ddc8fcb67ffe91900c762' as Address
const DEFAULT_FALLBACK_HANDLER_ADDRESS = '0xfd0732Dc9E303f09fCEf3a7388Ad10A83459Ec99' as Address
const SAFE_PROXY_FACTORY_ADDRESS = '0x4e1DCf7AD4e460CfD30791CCC4F9c8a4f820ec67' as Address
const TOPIC_PROXY_CREATION = '0x4f51faf6c4561ff95f067657e43439f0f856d97c04d9ec9070a6199ad418e235' as `0x${string}`
const ZERO_ADDRESS = '0x0000000000000000000000000000000000000000' as Address
const EMPTY_DATA = '0x' as `0x${string}`
const SALT_NONCE = 4n

const rpcUrl = 'https://mainnet-rpc.game7.io'

const publicClient = createPublicClient({
  chain: g7Network,
  transport: http(rpcUrl),
})

const account1 = privateKeyToAccount(process.env.PRIVATE_KEY as `0x${string}`)

const walletClient = createWalletClient({
  account: account1,
  chain: g7Network,
  transport: http(rpcUrl),
})

const getSafeAddressFromDeploymentTx = (receipt: TransactionReceipt): Address => {
  const proxyCreationLog = receipt.logs.find((log: Log) => log.topics[0] === TOPIC_PROXY_CREATION)

  if (!proxyCreationLog) {
    throw new Error('ProxyCreation event not found in transaction logs')
  }

  const proxyAddress = `0x${proxyCreationLog.topics[1]?.slice(26) ?? ''}` as Address
  return proxyAddress
}

async function deploySafe() {
  console.log('Deploying a new Safe...')

  try {
    const setupData = {
      owners: [account1.address],
      threshold: 1n,
      to: ZERO_ADDRESS,
      data: EMPTY_DATA,
      fallbackHandler: DEFAULT_FALLBACK_HANDLER_ADDRESS,
      paymentToken: ZERO_ADDRESS, 
      payment: 0n,
      paymentReceiver: ZERO_ADDRESS,
    }
    // 2. Encode the setup call
    const setupCalldata = encodeFunctionData({
      abi: safeAbi,
      functionName: 'setup',
      args: [
        setupData.owners,
        setupData.threshold,
        setupData.to,
        setupData.data,
        setupData.fallbackHandler,
        setupData.paymentToken,
        setupData.payment,
        setupData.paymentReceiver,
      ],
    })

    console.log('Sending deployment transaction...')
    const txHash = await walletClient.writeContract({
      address: SAFE_PROXY_FACTORY_ADDRESS,
      abi: proxyFactoryAbi,
      functionName: 'createProxyWithNonce',
      args: [SAFE_SINGLETON_ADDRESS, setupCalldata, SALT_NONCE],
    })

    console.log(`Transaction sent: ${txHash}`)

    console.log('Waiting for transaction confirmation...')
    const txReceipt = await publicClient.waitForTransactionReceipt({hash: txHash})

    const safeAddress = getSafeAddressFromDeploymentTx(txReceipt)
    console.log(`Safe deployed at: ${safeAddress}`)

    return safeAddress
  } catch (error) {
    console.error('Error deploying Safe:', error)
    throw error
  }
}

// Execute the deployment
deploySafe()
  .then(() => {
    console.log('Safe deployment completed successfully')
    process.exit(0)
  })
  .catch((error) => {
    console.error('Safe deployment failed:', error)
    process.exit(1)
  })
