import { FETCH_USER, RECEIVE_USER } from '../actions/userActions'

const usersReducer = (state = {}, action) => {
  switch (action.type) {
    case RECEIVE_USER:
      return action.payload;
    default:
      return state
  }
}

export default usersReducer;
