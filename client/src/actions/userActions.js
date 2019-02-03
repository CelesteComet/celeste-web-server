import axios from 'axios';
import { closeModal } from './uiActions';

export const FETCH_USER         = 'FETCH_USER';
export const RECEIVE_USER       = 'RECEIVE_USER';
export const FETCH_LOGOUT_USER  = 'FETCH_LOGOUT_USER';
export const LOGOUT_USER        = 'LOGOUT_USER';
export const CLEAR_USER         = 'CLEAR_USER';
export const LOGIN_USER         = 'LOGIN_USER';

export const fetchUser = () => {
  return dispatch => {
    axios.get('/auth')
      .then(res => {
        const user = res.data;
        dispatch(retrieveUser(user));
      })
      .catch(err => {
        console.log(err);
      })
  }
}

export const createUser = (user) => {
  return dispatch => {
    let url = `${process.env.AUTH_URL}${"/users"}`;
    return axios.post(url, user)
      .then(res => {
        document.cookie = `JWT=${res.headers['jwt']}`;
        dispatch(fetchUser());
        dispatch(closeModal());
      })
      .catch(err => {
        console.log(err); 
      })
  }
}

export const loginUser = (user) => {
  return dispatch => {
    let url = `${process.env.AUTH_URL}${"/login"}`;
    return axios.post(url, user)
      .then(res => {
        document.cookie = `JWT=${res.headers['jwt']}`;
        dispatch(fetchUser());   
        dispatch(closeModal());     
      })
      .catch(err => {
        console.log(err);
      })
  }
}

export const logoutUser = () => {
  return dispatch => {
    axios.get('/auth/logout')
      .then(res => {
        dispatch(clearUser());
      })
  }
}

export const retrieveUser = (user) => {
  return {
    type: RECEIVE_USER,
    payload: user
  }
}

export const clearUser = () => {
  return {
    type: CLEAR_USER, 
    payload: null 
  }
}