import axios from 'axios';

import { receiveErrors } from './errorActions';

export const POST_COMMENT = 'POST_COMMENT';
export const RECEIVE_COMMENTS = 'RECEIVE_COMMENTS';
export const FETCH_COMMENTS = 'FETCH_COMMENTS';
export const RECEIVE_COMMENT_DELETION = 'RECEIVE_COMMENT_DELETION';

export const postComment = (comment) => {
  return dispatch => {
    return axios.post(`/api/${comment.item_id}/comments`, comment)
      .then(res => {
        dispatch(fetchComments(comment.item_id));
      })
      .catch(err => {
        dispatch(receiveErrors(err));
      })
  }
}

export const fetchComments = (itemID) => {
  return dispatch => {
    return axios.get(`/api/${itemID}/comments`)
      .then(res => {
        dispatch(receiveComments(res.data));
      })
      .catch(err => {
        dispatch(receiveErrors(err));
      })
  }
}

export const deleteComment = (id) => {
  return dispatch => {
    return axios.delete(`/api/0/comments/${id}`)
      .then(res => {
        dispatch(receiveDeletion(res.data.id));
      })
      .catch(err => {
        dispatch(receiveErrors(err));
      })
  }
}

export const receiveComments = (comments) => {
  return {
    type: RECEIVE_COMMENTS,
    payload: comments
  }
}

export const receiveDeletion = (id) => {
  return {
    type: RECEIVE_COMMENT_DELETION,
    payload: id
  }
}
