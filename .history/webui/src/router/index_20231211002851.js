import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'
import ProfileView from '../views/ProfileView.vue'


const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [


		{
			path: '/login',
		    component: LoginView
		},

		{
			path: '/home',
		    component: HomeView
		},
		{
			path: '/users/:id',
			
		}

		{
			path: '/',
		    component: LoginView
		},


	]
})
router.beforeEach((to, from, next) => {
	if (to.path !== '/login' && !localStorage.token) next({ path: '/login' })
	else next()
  })


export default router
