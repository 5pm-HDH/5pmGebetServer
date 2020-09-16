import axios from 'axios';
import config from '../config';

const ajax = axios.create({
    timeout: 10000,
    headers: {'Content-Type': 'application/json'},
    baseURL: config.base_url
});


export default ajax;