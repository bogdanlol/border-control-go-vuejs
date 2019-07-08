export default {
	namespaced:true,
	state:{
		items:[],
		version:'5'
	},
	actions:{
		index({ dispatch,commit,getters,state,rootGetters,rootState },data) {
			return new Promise(async (resolve,reject) => {

				let res = await rootState.$root.$server.request(`/v${ state.version }/verificari`)
				
				commit('setItems',res || [])

				return resolve()
			})
		},
		store({ dispatch,commit,getters,state,rootGetters,rootState },new_verificare) {
			return new Promise(async (resolve,reject) => {
				
				let res = await rootState.$root.$server.request(`/v${ state.version }/verificari`,new_verificare,'POST')

				return resolve(res)
			})
		},
		edit({ dispatch,commit,getters,state,rootGetters,rootState },id) {
			return new Promise(async (resolve,reject) => {
				
				let res = await rootState.$root.$server.request(`/v${ state.version }/verificari/${ id }/edit`)

				return resolve(res)
			})
		},
		update({ dispatch,commit,getters,state,rootGetters,rootState },edited_verificare) {
			return new Promise(async (resolve,reject) => {
				
				let res = await rootState.$root.$server.request(`/v${ state.version }/verificari/${ edited_verificare.id }`,edited_verificare,'PUT')

				return resolve(res)
			})
		},
		destroy({ dispatch,commit,getters,state,rootGetters,rootState },id) {
			return new Promise(async (resolve,reject) => {
				
				let res = await rootState.$root.$server.request(`/v${ state.version }/verificari/${ id }`,{},'DELETE')

				return resolve(res)
			})
		}
	},
	mutations:{
		setItems(state,items) {
			state.items = items || []
		}
	}

}