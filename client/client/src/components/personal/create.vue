<template lang='pug'>

	div
		router-link(to='/personal') Back To Personal
		form.uk-form-stacked
			div
				label.uk-form-label Nume
				div.uk-form-controls
					input.uk-input(v-model='NUME',placeholder='Nume')
			div
				label.uk-form-label Prenume
				div.uk-form-controls
					input.uk-input(v-model='PREN',placeholder='Prenume')
			div
				label.uk-form-label Rang
				div.uk-form-controls
					input.uk-input(v-model='RANG',placeholder='Rang')
			hr
			div
				button.uk-button.uk-button-primary(type='button',@click.prevent='create') Create
			

</template> 

<script>
export default {
	name:'personal-create',
	data() {
		return {
			NUME:'',
			PREN:'',
			RANG:''
			
		}
	},
	methods:{
		create() {
			let vm = this

		 	let new_personal = {
		 		nume:vm.NUME,
				pren:vm.PREN,
				rang: vm.RANG

		 	}

		 	vm.$store
		 		.dispatch('personal/store',new_personal)
		 		.then(newly_created_personal=> {
		 			vm.$router.push({ path:'/personal' })
		 		})
		 		.catch(err => vm.$root.error(err))
		}
	},
	
}
</script>