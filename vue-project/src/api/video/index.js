import axios from 'axios'

const api = axios.create({
    baseURL: `${import.meta.env.VUE_APP_API_BASE_URL}:${import.meta.env.VUE_APP_API_PORT}`,
    timeout: 5000,
})

const handleResponse = (res) => {
    const { status, data, msg } = res.data;
    if (status !== 0) {
        throw new Error(msg || 'service error');
    }
    return data;
};

const uploadVideo = form => api.post('/api/v1/videos', form)
    .then(handleResponse);

const getVideo = id => api.get(`/api/v1/videos/${id}`)
    .then(handleResponse);

const getVideos = () => api.get('/api/v1/videos')
    .then(handleResponse);

export {
    uploadVideo,
    getVideo,
    getVideos,
};
