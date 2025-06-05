import axios from 'axios'

const api = axios.create({
    baseURL: 'http://localhost:3000',
    timeout: 5000
})

const uploadVideo = form => api.post('/api/v1/videos', form)
    .then(res => res.data)
    .catch(error => {
        throw new Error(error.response?.data?.message || error.message)
    });

export {
    uploadVideo,
}
