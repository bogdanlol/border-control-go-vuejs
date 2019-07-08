<template lang="pug">
    section.uk-section.uk-section-default
        form.uk-form-stacked.uk-list-large
            h3 Enter person name : {{name}}
            input.uk-input(v-model="name") 

            h3 Enter age : {{age}}
            input.uk-range(type ="range" value=15 min=0 max=100 step="1" v-model="age")
            
            ul.uk-list-divider
                    li(v-for="pers in persoane" v-if=("name == pers.NUME || name==pers.PREN || age == pers.varsta")) 
                        router-link(:to="'/persoane/' + pers.id + '/edit'") Nume:  {{ pers.NUME }} {{pers.PREN}}
                            li(v-for="masina in masini" v-if=("pers.codm==masina.id")) Masina:  {{masina.numar}} 
                            li(v-for="verificare in verificari" v-if=("pers.id==verificare.codp")) Ora Venirii : {{verificare.orav}}

               
                     
</template>

<script>

import { mapState } from 'vuex'
export default{
name:'cautare',
 computed:{
		...mapState({
			persoane:state => state.persoane.items,
            masini:state => state.masini.items,
            verificari:state => state.verificari.items,
            personal:state => state.personal.items

        }),
        //computed methods for filtering search
        
        
	},
data(){
    return{
        name:'',
        age:0
    }
},mounted(){
        let vm = this
		vm.$store.dispatch('masini/index')
        vm.$store.dispatch('persoane/index')
        vm.$store.dispatch('verificari/index')	
        vm.$store.dispatch('personal/index')
}   

//Cautare dupa nume
//Cautare cu filtre
//Comboboxes

//cautare 
}
</script>