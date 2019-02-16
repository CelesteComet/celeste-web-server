import { 
  POST_COMMENT, RECEIVE_COMMENTS, FETCH_COMMENTS, RECEIVE_COMMENT_DELETION
} from '../actions/commentActions';

const commentsReducer = (state = [], action) => {
  switch (action.type) {
    case POST_COMMENT:
      return action.payload;
    case RECEIVE_COMMENTS:
      return action.payload;
    case FETCH_COMMENTS:
      return action.payload;      
    case RECEIVE_COMMENT_DELETION:
      return state.filter(c => { return c.id !== action.payload }) // Return all comments that don't include the one deleted
    default:
      return state
  }
}

export default commentsReducer;
