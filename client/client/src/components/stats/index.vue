<template lang="pug">
h5 Current Year : {{CurrentYear}} Current Time : {{CurrentTime}}
    table.uk-table.uk-table-divider 
        caption Database Status          
        tbody 
        tr Numar de persoane in database : {{persoane.length}}
        tr Numar de masini in database : {{masini.length}}
        tr Numar de ofiteri in database : {{personal.length}}
        tr Numar de verificari in database :{{verificari.length}}
        
        p
        table.uk-table.uk-table-divider 
            caption Statistici
            tr
            tr Numar statistic de persoane per masina : {{parseFloat(persoane.length/masini.length).toFixed(2)}}
            tr Numar statistic de ofiteri per verificare : {{parseFloat(personal.length/verificari.length).toFixed(2)}}
            tr(v-for="verificare,key in verificari" v-if="key==verificari.length-1") Numar de verificari per an: {{key+1}} - {{verificare.orav.substring(0,4)}}   
            tr Varsta Medie Pe Persoana : {{parseFloat(varstaAvg).toFixed(2)}}

            //- //TODO
        
            //- // etc
</template>

<script>
import moment from 'moment'
import { mapState } from 'vuex'

export default{
    name:'stats',
    computed:{
		...mapState({
			persoane:state => state.persoane.items,
            masini:state => state.masini.items,
            personal:state => state.personal.items,
            verificari:state => state.verificari.items
            }),
            varstaAvg:function(){
                let sum =0
                this.persoane.forEach(function(persoane){
                        sum+= persoane.varsta;
                });
                return sum/this.persoane.length;
            }
        
	},
    data(){
        return{
         CurrentYear : new Date().getFullYear(),
         CurrentTime : moment().format("HH:mm"),
         varsta : 0,
        
         }
    },mounted(){
        let vm = this
		vm.$store.dispatch('masini/index')
        vm.$store.dispatch('persoane/index')
        vm.$store.dispatch('verificari/index')	
        vm.$store.dispatch('personal/index')
},
}
</script>
