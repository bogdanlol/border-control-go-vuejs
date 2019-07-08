
export default {
	namespaced:true,
	state:{
		items:[],
		version:'4'
	},
	actions:{
		index({ dispatch,commit,getters,state,rootGetters,rootState },data) {
			return new Promise(async (resolve,reject) => {

				let res = await rootState.$root.$server.request(`/v${ state.version }/personal`)
				
				commit('setItems',res || [])

				return resolve()
			})
		},
		store({ dispatch,commit,getters,state,rootGetters,rootState },new_personal) {
			return new Promise(async (resolve,reject) => {
				
				let res = await rootState.$root.$server.request(`/v${ state.version }/personal`,new_personal,'POST')

				return resolve(res)
			})
		},
		edit({ dispatch,commit,getters,state,rootGetters,rootState },id) {
			return new Promise(async (resolve,reject) => {
				
				let res = await rootState.$root.$server.request(`/v${ state.version }/personal/${ id }/edit`)

				return resolve(res)
			})
		},
		update({ dispatch,commit,getters,state,rootGetters,rootState },edited_personal) {
			return new Promise(async (resolve,reject) => {
				
				let res = await rootState.$root.$server.request(`/v${ state.version }/personal/${ edited_personal.id }`,edited_personal,'PUT')

				return resolve(res)
			})
		},
		destroy({ dispatch,commit,getters,state,rootGetters,rootState },id) {
			return new Promise(async (resolve,reject) => {
				
				let res = await rootState.$root.$server.request(`/v${ state.version }/personal/${ id }`,{},'DELETE')

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