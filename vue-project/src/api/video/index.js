import axios from 'axios'

const api = axios.create({
    baseURL: `${import.meta.env.VUE_APP_API_BASE_URL}:${import.meta.env.VUE_APP_API_PORT}`,
    timeout: 5000,
})

const uploadVideo = form => api.post('/api/v1/videos', form)
    .then(res => res.data);

const getVideo = id => api.get(`/api/v1/videos/${id}`)
    .then(res => res.data);

const getVideos = () => api.get('/api/v1/videos')
    .then(res => res.data);

export {
    uploadVideo,
    getVideo,
    getVideos,
};