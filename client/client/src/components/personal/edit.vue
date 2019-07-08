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
				button.uk-button.uk-button-primary(type='button',@click.prevent='save') Save

</template> 

<script>
export default {
	name:'personal-edit',
	data() {
		return {
			NUME:'',
			PREN:'',
			RANG:''
			
		}
	},
	computed: {
		id() { return this.$route.params.id }
	},
	methods:{
		save() {
			let vm = this

		 	let updated_personal = {
		 		id:vm.id,
		 		nume:vm.NUME,
				pren:vm.PREN,
				rang:vm.RANG
		 		
		 	}

		 	vm.$store.dispatch('personal/update',updated_personal)
		 		.then(newly_edited_personal => {
		 			vm.$router.push({ path:'/personal' })
		 		})
		 		.catch(err => vm.$root.error(err))
		}
	},
	mounted() {
		let vm = this

		vm.$store.dispatch(`personal/edit`,vm.id)
			.then(personal => {
				vm.NUME = personal.NUME
				vm.PREN = personal.PREN
				vm.RANG = personal.RANG
				
			})
		 	.catch(err => vm.$root.error(err))
	}
}
</script>