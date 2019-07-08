<template lang='pug'>

	table.uk-table.uk-table-striped
		caption
			router-link(to='/verificari/create') Create New Verificare
		thead
			tr
				th Persoana 
				th Masina
				th Ofiter
				th Ora Venirii
				th Ora Plecarii
				th
		tbody
			tr(v-for='verificare in verificari')
				td(v-for='persoana in persoane' v-if=('verificare.codp==persoana.id'))
					router-link(:to="'/verificari/' + verificare.id + '/edit'") {{ persoana.NUME }} {{persoana.PREN}}
				td(v-for='masina in masini' v-if=('verificare.codm==masina.id')) 
					router-link(:to="'/verificari/' + verificare.id + '/edit'") {{ masina.numar }} 

				td(v-for='pers in personal' v-if=('verificare.codo==pers.id'))
					router-link(:to="'/verificari/' + verificare.id + '/edit'") {{ pers.NUME }}  {{pers.PREN}} 
				td {{verificare.orav | moment}}
				td {{verificare.orap | moment}}
				td
					button.uk-button.uk-button-primary(@click.prevent='destroy(verificare.id)') Delete


</template>

<script>
import { mapState } from 'vuex'
import moment from 'moment'
export default {
	name:'verificari-index',
	computed:{
		...mapState({
			verificari:state => state.verificari.items,
			persoane:state=> state.persoane.items,
			masini:state=> state.masini.items,
			personal:state => state.personal.items
		})
	},
	mounted() {
		let vm = this
		vm.$store.dispatch('verificari/index')
		vm.$store.dispatch('persoane/index')
		vm.$store.dispatch('masini/index')
		vm.$store.dispatch('personal/index')
	},filters: {
  moment: function (date) {
    return moment(date).format('MMMM DDD YYYY, h:mm a');
  }
},
	
	methods:{
		destroy(id) {
			let vm = this
			
			vm.$store.dispatch('verificari/destroy',id)
				.then(res => {
					vm.$store.dispatch('verificari/index')
					vm.$store.dispatch('persoane/index')
					vm.$store.dispatch('masini/index')
					vm.$store.dispatch('personal/index')

				})
		}
	}
}
</script>

