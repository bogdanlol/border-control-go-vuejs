<template lang='pug'>

	div
		router-link(to='/verificari') Back To Verificare
		form.uk-form-stacked
			div
				label.uk-form-label Persoana
				div.uk-margin
					select.uk-select(v-model="persoana" )
						option(v-for="pers in persoane") {{pers.NUME}} {{pers.PREN}}
			div
				label.uk-form-label Masina
				div.uk-margin
					select.uk-select(v-model="masina")
						option(v-for="masina in masini") {{masina.numar}} {{masina.culoare}}		
			div
				label.uk-form-label  Ofiter
				div.uk-margin
					select.uk-select(v-model="ofiter")
						option(v-for="persona in personal") {{persona.NUME}} {{persona.PREN}}
			div
				label.uk-form-label Ora Venirii
				div.uk-form-controls
					input.uk-input(v-model='orav',placeholder='Ora venirii')
			div
				label.uk-form-label Ora Plecarii
				div.uk-form-controls
					input.uk-input(v-model='orap',placeholder='Ora Plecarii')
			hr
			hr
			div
				button.uk-button.uk-button-primary(type='button',@click.prevent='create') Create
			

</template> 

<script>
import { mapState } from 'vuex'

export default {
	name:'verificari-create',
	computed:{
		...mapState({
			verificari:state => state.verificari.items,
			persoane:state=> state.persoane.items,
			masini:state=> state.masini.items,
			personal:state => state.personal.items
		}),

	},
	data() {
		return {
			persoana:0,
			masina:0,
			ofiter:0,
			orav:new Date(),
			orap: new Date()
			
		}
	},
	methods:{
		create() {
			let vm = this
			// console.log(vm.persoana,vm.masina,vm.ofiter)
			// console.log("0",vm.persoana.split(" ",0))
			// console.log("1",vm.persoana.split(" ",4)[0])
			// console.log("2",vm.persoana.split(" ",4)[1])
				let codp=0
				let codm=0
				let codo=0
				this.persoane.forEach(function(persoane){
						
                        if(vm.persoana.split(" ",4)[0]==persoane.NUME || vm.persoana.split(" ",4)[1]==persoane.PREN){
							codp=persoane.id
							
						}
						
                });
               this.masini.forEach(function(masini){
						
                        if(vm.masina.split(" ",4)[0]==masini.numar || vm.masina.split(" ",4)[1]==masini.culoare){
							codm=masini.id
							
						}
						
				});
				this.personal.forEach(function(personal){
						
                        if(vm.ofiter.split(" ",4)[0]==personal.NUME || vm.ofiter.split(" ",4)[1]==personal.PREN){
							codo=personal.id
							
						}
						
                });
		 	let new_verificare = {
		 		codp:codp,
				codm:codm,
				codo:codo,
				orav:vm.orav,
				orap:vm.orap,
		 	}

		 	vm.$store
		 		.dispatch('verificari/store',new_verificare)
		 		.then(newly_created_verificare => {
		 			vm.$router.push({ path:'/verificare' })
		 		})
		 		.catch(err => vm.$root.error(err))
		}
	}, 
	
	
	
}
</script>