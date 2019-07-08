<template lang="pug">
    h3 logging out!
</template>

<script>
import Cookies from 'js-cookie'

export default {
    name:'session-destroy',
    beforeRouteEnter(to,from,next){

        Cookies.remove('proj1')

        location.reload(true)
    }
}
</script>
