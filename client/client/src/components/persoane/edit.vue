<template lang='pug'>

	div
		router-link(to='/persoane') Back To Persoane
		form.uk-form-stacked
			div
				label.uk-form-label Nume
				div.uk-form-controls
					input.uk-input(v-model='NUME',placeholder='NUME')
			div
				label.uk-form-label Prenume
				div.uk-form-controls
					input.uk-input(v-model='PREN',placeholder='PRENUME ')
			div
				label.uk-form-label CNP
				div.uk-form-controls
					input.uk-input(v-model='CNP',placeholder='CNP')
			div
				label.uk-form-label Varsta
				div.uk-form-controls
					input.uk-input(v-model='varsta',placeholder='Varsta')
			div
				label.uk-form-label Cod Masina
				div.uk-form-controls
					input.uk-input(v-model='codm',placeholder='codm')
			hr
			div
				button.uk-button.uk-button-primary(type='button',@click.prevent='save') Save

</template> 

<script>
export default {
	name:'persoane-edit',
	data() {
		return {
			NUME:'',
			PREN:'',
			CNP:'',
			varsta:0,
			codm:0
		}
	},
	computed: {
		id() { return this.$route.params.id }
	},
	methods:{
		save() {
			let vm = this

		 	let edited_persoana = {
		 		id:vm.id,
		 		nume:vm.NUME,
		 		pren:vm.PREN,
				cnp:vm.CNP,
				varsta:vm.varsta,
				codm:vm.codm
				 
		 	}

		 	vm.$store.dispatch('persoane/update',edited_persoana)
		 		.then(newly_edited_persoana => {
		 			vm.$router.push({ path:'/persoane' })
		 		})
		 		.catch(err => vm.$root.error(err))
		}
	},
	mounted() {
		let vm = this

		vm.$store.dispatch(`persoane/edit`,vm.id)
			.then(persoana => {
				vm.NUME = persoana.NUME
				vm.PREN = persoana.PREN
				vm.CNP = persoana.CNP
				vm.varsta = persoana.varsta
				vm.codm = persoana.codm
			})
		 	.catch(err => vm.$root.error(err))
	}
}
</script>