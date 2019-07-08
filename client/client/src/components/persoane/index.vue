<template lang='pug'>

	table.uk-table.uk-table-striped
		caption
			router-link(to='/persoane/create') Create New Persoane
		thead
			tr
				th nume 
				th PRENUME
				th CNP
				th VARSTA
				th Masina
				th
		tbody
			tr(v-for='persoana in persoane')
				td
					router-link(:to="'/persoane/' + persoana.id + '/edit'") {{ persoana.NUME }}
				td {{ persoana.PREN}}
				td {{ persoana.CNP }}
				td {{persoana.varsta}}
				td(v-for="masina in masini" v-if="persoana.codm==masina.id") 
					router-link(:to="'/masini/' + masina.id+ '/edit'") {{masina.numar}} {{masina.culoare}}
				td
					button.uk-button.uk-button-primary(@click.prevent='destroy(persoana.id)') Delete


</template>

<script>
import { mapState } from 'vuex'
export default {
	name:'persoane-index',
	computed:{
		...mapState({
			persoane:state => state.persoane.items,
			masini:state => state.masini.items
		})
	},
	mounted() {
		let vm = this
		vm.$store.dispatch('persoane/index')
		vm.$store.dispatch('masini/index')
	},
	methods:{
		destroy(id) {
			let vm = this
			
			vm.$store.dispatch('persoane/destroy',id)
				.then(res => {
					vm.$store.dispatch('persoane/index')
				})
		}
	}
}
</script>