import axios from 'axios';

const API = axios.create({ baseURL: 'http://localhost:8080' });

export const login = (userData) => API.post('/login', userData);
export const register = (userData) => API.post('/register', userData);
export const uploadFile = (fileData) => API.post('/upload', fileData);
export const saveShape = (shapeData) => API.post('/shapes', shapeData);
export const getShapes = () => API.get('/shapes');
