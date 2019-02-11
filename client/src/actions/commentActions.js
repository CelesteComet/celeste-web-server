import axios from 'axios';
export const POST_COMMENT = 'POST_COMMENT';
export const RECEIVE_COMMENTS = 'RECEIVE_COMMENTS';
export const FETCH_COMMENTS = 'FETCH_COMMENTS';

export const postComment = (comment) => {
  return dispatch => {
    return axios.post(`/api/${comment.item_id}/comments`, comment)
      .then(res => {
        dispatch(fetchComments(comment.item_id));
      })
  }
}

export const fetchComments = (itemID) => {
  return dispatch => {
    return axios.get(`/api/${itemID}/comments`)
      .then(res => {
        debugger;
        dispatch(receiveComments(res.data));
      })
  }
}

export const receiveComments = (comments) => {
  return {
    type: RECEIVE_COMMENTS,
    payload: comments
  }
}