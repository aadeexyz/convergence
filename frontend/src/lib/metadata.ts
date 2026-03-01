import {getAddress} from "viem"

export const factoryOfMarketFactoriesAddress = getAddress(process.env.NEXT_PUBLIC_FACTORY_OF_MARKET_FACTORIES_ADDRESS!)
export const lensAddress = getAddress(process.env.NEXT_PUBLIC_LENS_ADDRESS!)
export const routerAddress = getAddress(process.env.NEXT_PUBLIC_ROUTER_ADDRESS!)
export const usdcAddress = getAddress(process.env.NEXT_PUBLIC_USDC_ADDRESS!)
