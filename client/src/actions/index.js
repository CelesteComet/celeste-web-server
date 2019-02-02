import axios from 'axios';

export const RECEIVE_BAGS = 'RECEIVE_BAGS';
export const FETCH_BAGS ='FETCH_BAGS'

export const receiveBags = bags => ({
  type: RECEIVE_BAGS,
  payload: bags 
})

export const fetchBags = () => {
  return dispatch => {
    return axios.get('/api/bags')
      .then(res => {
        console.log("SHOULDA DISPATCHED")
        dispatch(receiveBags(res.data));
      })
  }
}
