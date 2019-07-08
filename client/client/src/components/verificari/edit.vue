<template lang='pug'>

	div
		router-link(to='/verificari') Back To Verificare
		form.uk-form-stacked
			div
				label.uk-form-label Cod Persoana
				div.uk-form-controls
					input.uk-input(v-model='codp',placeholder='Cod Persoana')
			div
				label.uk-form-label Cod Masina
				div.uk-form-controls
					input.uk-input(v-model='codm',placeholder='Cod Masina')
			div
				label.uk-form-label Cod Ofiter
				div.uk-form-controls
					input.uk-input(v-model='codo',placeholder='Cod Ofiter')
			div
				label.uk-form-label Ora venirii
				div.uk-form-controls
					input.uk-input(v-model='orav',placeholder='Ora Venirii')
			
			div
				label.uk-form-label Ora plecarii
				div.uk-form-controls
					input.uk-input(v-model='orap',placeholder='Ora Plecarii')
			hr
			div
				button.uk-button.uk-button-primary(type='button',@click.prevent='save') Save

</template> 

<script>
export default {
	name:'verificari-edit',
	data() {
		return {
			codp:0,
			codm: 0,
			codo: 0,
			orav:new Date(""),
			orap:new Date("")
			
		}
	},
	computed: {
		id() { return this.$route.params.id }
	},
	methods:{
		save() {
			let vm = this

		 	let updated_verificare = {
		 		 id:vm.id,
		 		 codp:vm.codp,
				 codm:vm.codm,
				 codo:vm.codo,
				 orav:vm.orav,
				 orap:vm.orap
		 		
		 	}

		 	vm.$store.dispatch('verificari/update',updated_verificare)
		 		.then(newly_edited_verificare => {
		 			vm.$router.push({ path:'/verificari' })
		 		})
		 		.catch(err => vm.$root.error(err))
		}
	},
	mounted() {
		let vm = this

		vm.$store.dispatch(`verificari/edit`,vm.id)
			.then(verificare => {
				vm.codp = verificare.codp
				vm.codm = verificare.codm
				vm.codo = verificare.codo
				vm.orav = verificare.orav
				vm.orap = verificare.orap
			
				
			})
		 	.catch(err => vm.$root.error(err))
	}
}
</script>