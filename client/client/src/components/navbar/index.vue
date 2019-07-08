
<template lang="pug">
    


nav.uk-navbar-container(uk-navbar)
    div.uk-navbar-left
        ul.uk-navbar-nav
            li
                router-link(to='/') Dashboard
            li
                router-link(to='/persoane') Persoane
            li
                router-link(to='/masini') Masini
            li
                router-link(to='/personal') Personal
            li
                router-link(to='/verificari') Verificari
            li
                router-link(to='/cautare') Cautare
          
                

    div.uk-navbar-right
        ul.uk-navbar-nav
            li 
                router-link(to='/stats') Stats
            li
                router-link(to='/session/destroy') Logout
       
</nav>
</template>

<script>
export default {
    name:'navbar'
}
</script>

<style>
.uk-navbar-container:not(.uk-navbar-transparent) {

    background: black;
    border: 2px transparent red;
    overflow : hidden;
}
.uk-navbar-nav > li > a {

    color: #93ddf8;
    
    }
.uk-navbar-nav > li:hover > a, .uk-navbar-nav > li > a.uk-open, .uk-navbar-nav > li > a:focus {

    color: #9b1414;
    outline: 0;

}
</style>
