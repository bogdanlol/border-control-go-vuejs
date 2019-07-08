<template lang='pug'>

	table.uk-table.uk-table-striped
		caption
			router-link(to='/masini/create') Create New Masina
		thead
			tr
				th Numar
				th Culoare
				th
		tbody
			tr(v-for='masina in masini')
				td
					router-link(:to="'/masini/' + masina.id + '/edit'") {{ masina.numar }}
				td {{ masina.culoare }}
				td
					button.uk-button.uk-button-primary(@click.prevent='destroy(masina.id)') Delete


</template>

<script>
import { mapState } from 'vuex'
export default {
	name:'masini-index',
	computed:{
		...mapState({
			masini:state => state.masini.items
		})
	},
	mounted() {
		let vm = this
		vm.$store.dispatch('masini/index')
	},
	methods:{
		destroy(id) {
			let vm = this
			
			vm.$store.dispatch('masini/destroy',id)
				.then(res => {
					vm.$store.dispatch('masini/index')
				})
		}
	}
}
</script>