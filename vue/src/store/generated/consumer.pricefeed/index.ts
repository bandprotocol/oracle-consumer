import { Client, registry, MissingWalletError } from 'consumer-client-ts'

import { RequestInterval } from "consumer-client-ts/consumer.pricefeed/types"
import { PriceFeed } from "consumer-client-ts/consumer.pricefeed/types"
import { UpdateSymbolRequestProposal } from "consumer-client-ts/consumer.pricefeed/types"
import { SymbolRequests } from "consumer-client-ts/consumer.pricefeed/types"
import { SymbolRequest } from "consumer-client-ts/consumer.pricefeed/types"
import { PriceFeedPacketData } from "consumer-client-ts/consumer.pricefeed/types"
import { NoData } from "consumer-client-ts/consumer.pricefeed/types"
import { Params } from "consumer-client-ts/consumer.pricefeed/types"


export { RequestInterval, PriceFeed, UpdateSymbolRequestProposal, SymbolRequests, SymbolRequest, PriceFeedPacketData, NoData, Params };

function initClient(vuexGetters) {
	return new Client(vuexGetters['common/env/getEnv'], vuexGetters['common/wallet/signer'])
}

function mergeResults(value, next_values) {
	for (let prop of Object.keys(next_values)) {
		if (Array.isArray(next_values[prop])) {
			value[prop]=[...value[prop], ...next_values[prop]]
		}else{
			value[prop]=next_values[prop]
		}
	}
	return value
}

type Field = {
	name: string;
	type: unknown;
}
function getStructure(template) {
	let structure: {fields: Field[]} = { fields: [] }
	for (const [key, value] of Object.entries(template)) {
		let field = { name: key, type: typeof value }
		structure.fields.push(field)
	}
	return structure
}
const getDefaultState = () => {
	return {
				Params: {},
				RequestInterval: {},
				Price: {},
				
				_Structure: {
						RequestInterval: getStructure(RequestInterval.fromPartial({})),
						PriceFeed: getStructure(PriceFeed.fromPartial({})),
						UpdateSymbolRequestProposal: getStructure(UpdateSymbolRequestProposal.fromPartial({})),
						SymbolRequests: getStructure(SymbolRequests.fromPartial({})),
						SymbolRequest: getStructure(SymbolRequest.fromPartial({})),
						PriceFeedPacketData: getStructure(PriceFeedPacketData.fromPartial({})),
						NoData: getStructure(NoData.fromPartial({})),
						Params: getStructure(Params.fromPartial({})),
						
		},
		_Registry: registry,
		_Subscriptions: new Set(),
	}
}

// initial state
const state = getDefaultState()

export default {
	namespaced: true,
	state,
	mutations: {
		RESET_STATE(state) {
			Object.assign(state, getDefaultState())
		},
		QUERY(state, { query, key, value }) {
			state[query][JSON.stringify(key)] = value
		},
		SUBSCRIBE(state, subscription) {
			state._Subscriptions.add(JSON.stringify(subscription))
		},
		UNSUBSCRIBE(state, subscription) {
			state._Subscriptions.delete(JSON.stringify(subscription))
		}
	},
	getters: {
				getParams: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Params[JSON.stringify(params)] ?? {}
		},
				getRequestInterval: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.RequestInterval[JSON.stringify(params)] ?? {}
		},
				getPrice: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Price[JSON.stringify(params)] ?? {}
		},
				
		getTypeStructure: (state) => (type) => {
			return state._Structure[type].fields
		},
		getRegistry: (state) => {
			return state._Registry
		}
	},
	actions: {
		init({ dispatch, rootGetters }) {
			console.log('Vuex module: consumer.pricefeed initialized!')
			if (rootGetters['common/env/client']) {
				rootGetters['common/env/client'].on('newblock', () => {
					dispatch('StoreUpdate')
				})
			}
		},
		resetState({ commit }) {
			commit('RESET_STATE')
		},
		unsubscribe({ commit }, subscription) {
			commit('UNSUBSCRIBE', subscription)
		},
		async StoreUpdate({ state, dispatch }) {
			state._Subscriptions.forEach(async (subscription) => {
				try {
					const sub=JSON.parse(subscription)
					await dispatch(sub.action, sub.payload)
				}catch(e) {
					throw new Error('Subscriptions: ' + e.message)
				}
			})
		},
		
		
		
		 		
		
		
		async QueryParams({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ConsumerPricefeed.query.queryParams()).data
				
					
				commit('QUERY', { query: 'Params', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryParams', payload: { options: { all }, params: {...key},query }})
				return getters['getParams']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryParams API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryRequestInterval({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ConsumerPricefeed.query.queryRequestInterval( key.symbol)).data
				
					
				commit('QUERY', { query: 'RequestInterval', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryRequestInterval', payload: { options: { all }, params: {...key},query }})
				return getters['getRequestInterval']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryRequestInterval API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryPrice({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ConsumerPricefeed.query.queryPrice( key.symbol)).data
				
					
				commit('QUERY', { query: 'Price', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryPrice', payload: { options: { all }, params: {...key},query }})
				return getters['getPrice']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryPrice API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgCreateRequestInterval({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.ConsumerPricefeed.tx.sendMsgCreateRequestInterval({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateRequestInterval:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateRequestInterval:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgCreateRequestInterval({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ConsumerPricefeed.tx.msgCreateRequestInterval({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateRequestInterval:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateRequestInterval:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}
