import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		
		{
			path: '/login',
		    component: LoginView
		}

		{
			path: '/home',
		    component: HomeView
		},


	]
})



export default router
