import axios from 'axios';

const ajax = axios.create({
    baseURL: '/api',
    timeout: 10000,
    headers: {'Content-Type': 'application/json'}
});

export default ajax;