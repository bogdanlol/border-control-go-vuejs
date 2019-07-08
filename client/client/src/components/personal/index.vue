<template lang='pug'>

	table.uk-table.uk-table-striped
		caption
			router-link(to='/personal/create') Create New Personal
		thead
			tr
				th Nume
				th Prenume
				th Rang
				th
		tbody
			tr(v-for='pers in personal')
				td
					router-link(:to="'/personal/' + pers.id + '/edit'") {{ pers.NUME }}
				td {{ pers.PREN }}
				td {{ pers.RANG }}
				td
					button.uk-button.uk-button-primary(@click.prevent='destroy(pers.id)') Delete


</template>

<script>
import { mapState } from 'vuex'
export default {
	name:'personal-index',
	computed:{
		...mapState({
			personal:state => state.personal.items
		})
	},
	mounted() {
		let vm = this 
		vm.$store.dispatch('personal/index')
	},
	methods:{
		destroy(id) {
			let vm = this
			
			vm.$store.dispatch('personal/destroy',id)
				.then(res => {
					vm.$store.dispatch('personal/index')
				})
		}
	}
}
</script>