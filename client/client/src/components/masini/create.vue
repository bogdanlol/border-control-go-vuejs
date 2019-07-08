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
				button.uk-button.uk-button-primary(type='button',@click.prevent='create') Create
			

</template> 

<script>
export default {
	name:'users-create',
	data() {
		return {
			numar:'',
			culoare:''
			
		}
	},
	methods:{
		create() {
			let vm = this

		 	let new_masina = {
		 		numar:vm.numar,
		 		culoare:vm.culoare,

		 	}

		 	vm.$store
		 		.dispatch('masini/store',new_masina)
		 		.then(newly_created_masina => {
		 			vm.$router.push({ path:'/masini' })
		 		})
		 		.catch(err => vm.$root.error(err))
		}
	},
	
}
</script>