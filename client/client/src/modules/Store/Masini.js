export default {
	namespaced:true,
	state:{
		items:[],
		version:'3'
	},
	actions:{
		index({ dispatch,commit,getters,state,rootGetters,rootState },data) {
			return new Promise(async (resolve,reject) => {

				let res = await rootState.$root.$server.request(`/v${ state.version }/masini`)
				
				commit('setItems',res || [])

				return resolve()
			})
		},
		store({ dispatch,commit,getters,state,rootGetters,rootState },new_masina) {
			return new Promise(async (resolve,reject) => {
				
				let res = await rootState.$root.$server.request(`/v${ state.version }/masini`,new_masina,'POST')

				return resolve(res)
			})
		},
		edit({ dispatch,commit,getters,state,rootGetters,rootState },id) {
			return new Promise(async (resolve,reject) => {
				
				let res = await rootState.$root.$server.request(`/v${ state.version }/masini/${ id }/edit`)

				return resolve(res)
			})
		},
		update({ dispatch,commit,getters,state,rootGetters,rootState },edited_masina) {
			return new Promise(async (resolve,reject) => {
				
				let res = await rootState.$root.$server.request(`/v${ state.version }/masini/${ edited_masina.id }`,edited_masina,'PUT')

				return resolve(res)
			})
		},
		destroy({ dispatch,commit,getters,state,rootGetters,rootState },id) {
			return new Promise(async (resolve,reject) => {
				
				let res = await rootState.$root.$server.request(`/v${ state.version }/masini/${ id }`,{},'DELETE')

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