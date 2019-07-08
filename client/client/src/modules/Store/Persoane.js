export default {
	namespaced:true,
	state:{
		items:[],
		version:'2'
	},
	actions:{
		index({ dispatch,commit,getters,state,rootGetters,rootState },data) {
			return new Promise(async (resolve,reject) => {

				let res = await rootState.$root.$server.request(`/v${ state.version }/persoane`)
				
				commit('setItems',res || [])

				return resolve()
			})
		},
		store({ dispatch,commit,getters,state,rootGetters,rootState },new_persoana) {
			return new Promise(async (resolve,reject) => {
				
				let res = await rootState.$root.$server.request(`/v${ state.version }/persoane`,new_persoana,'POST')

				return resolve(res)
			})
		},
		edit({ dispatch,commit,getters,state,rootGetters,rootState },id) {
			return new Promise(async (resolve,reject) => {
				
				let res = await rootState.$root.$server.request(`/v${ state.version }/persoane/${ id }/edit`)

				return resolve(res)
			})
		},
		update({ dispatch,commit,getters,state,rootGetters,rootState },edited_persoana) {
			return new Promise(async (resolve,reject) => {
				
				let res = await rootState.$root.$server.request(`/v${ state.version }/persoane/${ edited_persoana.id }`,edited_persoana,'PUT')

				return resolve(res)
			})
		},
		destroy({ dispatch,commit,getters,state,rootGetters,rootState },id) {
			return new Promise(async (resolve,reject) => {
				
				let res = await rootState.$root.$server.request(`/v${ state.version }/persoane/${ id }`,{},'DELETE')

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