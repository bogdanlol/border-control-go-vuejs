<template lang="pug">

	div(class=hi).uk-flex.uk-flex-wrap
		
		tablemasini
	
		tablepersoane

		tableverificari
		
	
</template>

<script>

import { mapState } from 'vuex'

import tablemasini from './tablemasini.vue'
import tableusers from './tableusers.vue'
import tabelpersonal from './tabelpersonal.vue'
import tableverificari from './tableverificari.vue'
import tablepersoane from './tablepersoana.vue'
export default{
	name:'dashboard',
	
    computed:{
		...mapState({
			users:state => state.users.items,
			masini:state => state.masini.items
		})
	},
		mounted(){
		let vm = this
		vm.$store.dispatch('masini/index')
		vm.$store.dispatch('users/index')	
	},
	components:
	{
		tablemasini,
		tableusers,
		tabelpersonal,
		tableverificari,
		tablepersoane
	}
      

//Persoane page | link cars
//Create Persoane | Responsive codm selector and varsta picker
//Verificari | Format  the date 
//Create new verificare | date & empty checkers
//cautare | event target by name, cauta as ofiter, cauta by date, cauta by multiple
//random stats add
//Interfata grafica ++ final animatations,transitions
//etc
//Verificari easy add
//add restriction to cnp EDIT AND CREATE@@
}
</script>
<style scoped>

.uk-flex{
   
margin:25px;
border:1px solid;
border-color: rgba(0, 0, 0, 0.055);
background-color:rgba(0, 0, 0, 0.075);
font-family: 'Courier New', Courier, monospace, Geneva, Tahoma, sans-serif; 
padding-top:50px;

}
	
</style>