import Vue from 'vue'
import Router from 'vue-router'

let routes =[]
Vue.use(Router)

const Dashboard = () => import('@/components/dashboard/index.vue')
routes.push({ path:'/',component:Dashboard,meta:{ auth:true } })





const Persoane = () => import('@/components/persoane/index.vue')
const PersoaneCreate = () => import('@/components/persoane/create.vue')
const PersoaneEdit = () => import('@/components/persoane/edit.vue')


routes.push({ path:'/persoane',component:Persoane,meta:{ auth:true } })
routes.push({ path:'/persoane/create',component:PersoaneCreate,meta:{ auth:true } })
routes.push({ path:'/persoane/:id/edit',component:PersoaneEdit,meta:{ auth:true } })


const Masini = () => import('@/components/masini/index.vue')
const MasiniCreate = () => import('@/components/masini/create.vue')
const MasiniEdit = () => import('@/components/masini/edit.vue')


routes.push({ path:'/masini',component:Masini,meta:{ auth:true } })
routes.push({ path:'/masini/create',component:MasiniCreate,meta:{ auth:true } })
routes.push({ path:'/masini/:id/edit',component:MasiniEdit,meta:{ auth:true } })


const Verificari = () => import('@/components/verificari/index.vue')
const VerificariCreate = () => import('@/components/verificari/create.vue')
const VerificariEdit = () => import('@/components/verificari/edit.vue')


routes.push({ path:'/verificari',component:Verificari,meta:{ auth:true } })
routes.push({ path:'/verificari/create',component:VerificariCreate,meta:{ auth:true } })
routes.push({ path:'/verificari/:id/edit',component:VerificariEdit,meta:{ auth:true } })


const Personal = () => import('@/components/personal/index.vue')
const PersonalCreate = () => import('@/components/personal/create.vue')
const PersonalEdit = () => import('@/components/personal/edit.vue')


routes.push({ path:'/personal',component:Personal,meta:{ auth:true } })
routes.push({ path:'/personal/create',component:PersonalCreate,meta:{ auth:true } })
routes.push({ path:'/personal/:id/edit',component:PersonalEdit,meta:{ auth:true } })



const Users = () => import('@/components/users/index.vue')
const UsersCreate = () => import('@/components/users/create.vue')
const UsersEdit = () => import('@/components/users/edit.vue')


routes.push({ path:'/users',component:Users,meta:{ auth:true } })
routes.push({ path:'/users/create',component:UsersCreate,meta:{ auth:true } })
routes.push({ path:'/users/:id/edit',component:UsersEdit,meta:{ auth:true } })


const Cautare= () => import('@/components/cautare/index.vue')
routes.push({path:'/cautare',component:Cautare,meta:{auth:true}})

const Stats= () => import('@/components/stats/index.vue')
routes.push({path:'/stats',component:Stats,meta:{auth:true}})

const Session = () => import('@/components/session/index.vue')
const SessionCreate = () => import('@/components/session/create.vue')
const SessionDestroy = () => import('@/components/session/destroy.vue')
routes.push({path: '/session', component: Session,children:[
  { path : 'create' , component: SessionCreate,meta:{auth:false}},
  {path : 'destroy', component: SessionDestroy , meta:{auth:false}}
],meta:{auth:false} })


export default new Router({
  mode: 'history',
  routes

    
  
})
