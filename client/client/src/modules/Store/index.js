import users from '@/modules/Store/Users.js'
import persoane from '@/modules/Store/Persoane.js'
import masini from '@/modules/Store/Masini.js'
import personal from '@/modules/Store/Personal.js'
import verificari from '@/modules/Store/Verificari.js'
export default {
	modules:{
		users,
		persoane,
		masini,
		verificari,
		personal,
	},
	state:{
		$root:{}
	},
	actions:{
		setRoot(context,data) {
			context.state.$root = data
		}
	}
}