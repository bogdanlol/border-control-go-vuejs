<template lang='pug'>

	table.uk-table-small.uk-table-condensed.uk-table.uk-table-divided.uk-table-striped
		caption
			h4.uk-text-center Persoane
		thead
			tr
				th Nume
				th Prenume
				th CNP
				th Varsta
				th Masina
				th
		tbody
			tr(v-for='persoana in persoane' )
				td
					router-link(:to="'/persoane/' + persoana.id + '/edit'") {{ persoana.NUME }}
				td {{ persoana.PREN}}
				td {{ persoana.CNP }}
				td {{ persoana.varsta }}
				td(v-for="masina in masini" v-if="persoana.codm==masina.id") 
					router-link(:to="'/masini/'+persoana.codm+'/edit'") {{masina.numar}}
				td
					


</template>

<script>
import { mapState } from 'vuex'
export default {
	name:'table-persoane',
	computed:{
		...mapState({
			persoane:state => state.persoane.items,
			masini:state => state.masini.items
		})
	},
	mounted() {
		let vm = this
		vm.$store.dispatch('persoane/index')
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