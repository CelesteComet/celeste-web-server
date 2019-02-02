import axios from 'axios';

export const FETCH_USER = 'FETCH_USER';
export const RECEIVE_USER = 'RECEIVE_USER';

export const fetchUser = () => {
  return dispatch => {
    axios.get('/auth')
      .then(res => {
        const user = res.data;
        dispatch(retrieveUser(user));
      })
  }
}

export const retrieveUser = (user) => {
  return {
    type: RECEIVE_USER,
    payload: user
  }
}