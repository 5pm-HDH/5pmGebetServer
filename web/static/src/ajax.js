import axios from 'axios';

const ajax = axios.create({
    timeout: 10000,
    headers: {'Content-Type': 'application/json'},
});

export default ajax;