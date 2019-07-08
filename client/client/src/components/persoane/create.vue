<template lang='pug'>

	div
		router-link(to='/persoane') Back To Persoane
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
				label.uk-form-label CNP
				div.uk-form-controls
					input.uk-input(v-model='CNP',placeholder='CNP')
			div
				label.uk-form-label Varsta
				div.uk-form-controls
					input.uk-input(v-model='varsta',placeholder='Varsta')
			div
				label.uk-form-label Masina
				div.uk-margin
					select.uk-select(v-model="masina")
						option(v-for="masina in masini") {{masina.numar}} {{masina.culoare}}
			hr
			div
				button.uk-button.uk-button-primary(type='button',@click.prevent='create') Create
	

</template> 

<script>
import { mapState } from 'vuex'
export default {
	name:'persoane-create',
	computed:{
		...mapState({
			persoane:state => state.persoane.items,
			masini:state => state.masini.items
		})
	},
	data() {
		return {
			NUME:'',
			PREN:'',
			CNP:'',
			varsta:0,
			masina:0
		}
	},
	methods:{
		create() {
			 
			let vm = this
			let codm=0
			this.masini.forEach(function(masini){
						
                        if(vm.masina.split(" ",4)[0]==masini.numar || vm.masina.split(" ",4)[1]==masini.culoare){
							codm=masini.id
							
						}
						});

		 	let new_persoana = {
				
		 		nume:vm.NUME,
		 		pren:vm.PREN,
		 		cnp:vm.CNP,
				varsta:vm.varsta,
				codm:codm
				
		 	}

		 	vm.$store
		 		.dispatch('persoane/store',new_persoana)
		 		.then(newly_created_persoana => {
					
		 			vm.$router.push({ path:'/persoane' })
		 		})
				 .catch(err => vm.$root.error(err))
				
		}
	}
	
}
</script>