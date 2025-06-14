import axios from 'axios'

const api = axios.create({
    baseURL: import.meta.env.VITE_API_BASE_URL || 'http://localhost:3000',
    timeout: 5000,
})

const uploadVideo = form => api.post('/api/v1/videos', form).then(res => res.data);

const getVideo = id => api.get(`/api/v1/videos/${id}`);

const getVideos = () => api.get('/api/v1/videos').then(res => res.data);

export {
    uploadVideo,
    getVideo,
    getVideos,
};
