import axios from 'axios';

export const RECEIVE_BAGS = 'RECEIVE_BAGS';
export const RECEIVE_BAG = 'RECEIVE_BAG';
export const FETCH_BAGS ='FETCH_BAGS';

export const receiveBags = bags => {
  return {
    type: RECEIVE_BAGS,
    payload: bags
  }
}

export const receiveBag = bag => {
  return {
    type: RECEIVE_BAG,
    payload: [bag]
  }
}

export const fetchBags = () => {
  return dispatch => {
    return axios.get('/api/bags')
      .then(res => {
        dispatch(receiveBags(res.data));
      })
  }
}

export const fetchBag = (i) => {
  return dispatch => {
    return axios.get(`/api/bags/${i}`)
      .then(res => {
        dispatch(receiveBag(res.data))
      })
  }
}
