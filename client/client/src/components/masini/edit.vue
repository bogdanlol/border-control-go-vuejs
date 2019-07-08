<template lang='pug'>

	div
		router-link(to='/masini') Back To Masini
		form.uk-form-stacked
			div
				label.uk-form-label Numar
				div.uk-form-controls
					input.uk-input(v-model='numar',placeholder='Numar')
			div
				label.uk-form-label Culoare
				div.uk-form-controls
					input.uk-input(v-model='culoare',placeholder='Culoare')
			hr
			div
				button.uk-button.uk-button-primary(type='button',@click.prevent='save') Save

</template> 

<script>
export default {
	name:'masini-edit',
	data() {
		return {
			numar:'',
			culoare:''
			
		}
	},
	computed: {
		id() { return this.$route.params.id }
	},
	methods:{
		save() {
			let vm = this

		 	let updated_masina = {
		 		id:vm.id,
		 		numar:vm.numar,
		 		culoare:vm.culoare
		 		
		 	}

		 	vm.$store.dispatch('masini/update',updated_masina)
		 		.then(newly_edited_masina => {
		 			vm.$router.push({ path:'/masini' })
		 		})
		 		.catch(err => vm.$root.error(err))
		}
	},
	mounted() {
		let vm = this

		vm.$store.dispatch(`masini/edit`,vm.id)
			.then(masina => {
				vm.numar = masina.numar
				vm.culoare = masina.culoare
				
			})
		 	.catch(err => vm.$root.error(err))
	}
}
</script>