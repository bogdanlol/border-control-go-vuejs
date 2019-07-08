<template lang="pug">
table.uk-table.uk-table-small.uk-table-divider.uk-table-striped
        caption
            h4.uk-text-center Autoturisme 
        thead
            tr
                th Numar
                th Culoare
                th

        tbody
            tr(v-for='masina in masini')
                td
                    router-link(:to="'/masini/' + masina.id + '/edit'") {{ masina.numar }}
                td {{ masina.culoare }}
                td

</template>

<script>
import { mapState } from 'vuex'
export default{
name:'table-masini',
computed:{
		...mapState({
			masini:state => state.masini.items
		})
    },
    mounted(){
		let vm = this
		vm.$store.dispatch('masini/index')
        vm.$store.dispatch('users/index')
        }
    }
</script>
