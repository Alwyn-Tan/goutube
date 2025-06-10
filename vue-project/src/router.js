import {createRouter, createWebHistory} from 'vue-router';
import Home from '@/views/Home.vue';

export default createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: '/',
            name: 'home',
            component: Home,
        },
        {
            path: '/upload',
            name: 'upload',
            component: () => import('./views/UploadVideo.vue'),
        },
        {
            path: '/about',
            name: 'about',
            component: () => import('./views/About.vue'),
        }
    ]
})
