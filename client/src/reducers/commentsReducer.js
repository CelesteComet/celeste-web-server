import { 
  POST_COMMENT, RECEIVE_COMMENTS, FETCH_COMMENTS
} from '../actions/commentActions';

const commentsReducer = (state = [], action) => {
  switch (action.type) {
    case POST_COMMENT:
      return action.payload;
    case RECEIVE_COMMENTS:
      return action.payload;
    case FETCH_COMMENTS:
      return action.payload;      
    default:
      return state
  }
}

export default commentsReducer;
