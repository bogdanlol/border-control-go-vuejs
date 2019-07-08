<template lang='pug'>

	table.uk-table-small.uk-table-condensed.uk-table.uk-table-divider.uk-table-striped
		caption
			h4.uk-text-center Verificari
		thead
			tr
				th.uk-table-shrink Persoana 
				th.uk-table-shrink Masina
				th.uk-table-shrink Ofiter
				th Ora Venirii
				th Ora Plecarii
				th
		tbody
			tr(v-for='verificare in verificari')
				td(v-for='pers in persoane' v-if='verificare.codp==pers.id')
					router-link(:to="'/persoane/' + verificare.codp + '/edit'") {{ pers.NUME }} {{pers.PREN}}
				td(v-for="masina in masini" v-if="verificare.codm==masina.id")
					router-link(:to="'/masini/' + verificare.codm + '/edit'") {{masina.numar}} {{masina.culoare}}
				td(v-for="pers in personal" v-if="verificare.codo==pers.id") 
					router-link(:to="'/personal/' + verificare.codo + '/edit'") {{pers.NUME}} {{pers.PREN}}
				td {{verificare.orav | moment}}
				td {{verificare.orap | moment}}
				td
					


</template>

<script>
import { mapState } from 'vuex'
import moment from 'moment'
export default {
	name:'table-verificari',
	computed:{
		...mapState({
			verificari:state => state.verificari.items,
			masini:state => state.masini.items,
			persoane:state => state.persoane.items,
			personal:state => state.personal.items
		})
	},
	mounted() {
		let vm = this
		vm.$store.dispatch('verificari/index')
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
				})
		}
	}
}
</script>