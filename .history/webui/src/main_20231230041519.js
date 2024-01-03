import {createApp, reactive} from 'vue'
import App from './App.vue'
import router from './router'
import axios from './services/axios.js';
import ErrorMsg from './components/ErrorMsg.vue'
import LoadingSpinner from './components/LoadingSpinner.vue'
import Photo from './components/Photo.vue'
import LikeModal from './components/LikeModel.vue'
import CommentModal from './components/CommentModel.vue'
import PhotoComment from './components/PhotoComment.vue'
import UserMini from './components/UserMini.vue'


import './assets/dashboard.css'
import './assets/main.css'

import { library } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import { fas } from '@fortawesome/free-solid-svg-icons'
import { far } from '@fortawesome/free-regular-svg-icons'

const app = createApp(App)
app.config.globalProperties.$axios = axios;

app.component("ErrorMsg", ErrorMsg);
app.component("LoadingSpinner", LoadingSpinner);
app.component("Photo", Photo);
app.component("UserMini", UserMini);
app.component("LikeModel", LikeModel);
app.component("PhotoComment", PhotoComment);
app.component("CommentModal", CommentModal);

library.add(fas, far);

app.component('font-awesome-icon', FontAwesomeIcon)

app.use(router)
app.mount('#app')
