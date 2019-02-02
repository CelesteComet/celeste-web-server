import { 
  FETCH_USER, 
  RECEIVE_USER,
  LOGOUT_USER,
  CLEAR_USER
} 
from '../actions/userActions'

const usersReducer = (state = null, action) => {
  switch (action.type) {
    case RECEIVE_USER:
      return action.payload;
    case CLEAR_USER:
      return action.payload;
    default:
      return state
  }
}

export default usersReducer;
