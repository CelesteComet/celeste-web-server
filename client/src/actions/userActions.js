import axios from 'axios';
import { closeLoginForm } from './uiActions';
import { receiveErrors  } from './errorActions';

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
        dispatch(receiveUser(user));
      })
      .catch(err => {
        dispatch(receiveErrors(err.response.data));
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
        dispatch(closeLoginForm());
      })
      .catch(err => {
        dispatch(receiveErrors(err));
      })
  }
}

export const loginUser = (user) => {
  return dispatch => {
    let url = "/auth";
    return axios.post(url, user)
      .then(res => {
        dispatch(receiveUser(res));   
      })
      .catch(err => {
        dispatch(receiveErrors(err));
      })
  }
}

export const loginWithGoogle = () => {
  return dispatch => {
    let url = `${process.env.AUTH_URL}${`/auth/google`}`;
    return axios.get(url)
      .then(res => {
        console.log("WTF")
        debugger;
        console.log(res);
      })
      .catch(err => {
        console.log("WawdawdTF")
        debugger;
        console.log(err.response.data);
      })
  } 
}

export const logoutUser = () => {
  return dispatch => {
    axios.delete('/auth')
      .then(res => {
        dispatch(clearUser());
      })
  }
}

export const receiveUser = (user) => {
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