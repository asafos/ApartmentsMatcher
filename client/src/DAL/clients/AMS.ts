import axios from 'axios'

const AMSHttpClient = axios.create({
  baseURL: 'http://localhost:8080',
  timeout: 30000,
  withCredentials: true,
})

export { AMSHttpClient }
