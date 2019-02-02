import { FETCH_BAGS, RECEIVE_BAGS } from '../actions'

const bagsReducer = (state = [], action) => {
  switch (action.type) {
    case RECEIVE_BAGS:
      return action.payload;
    default:
      return state
  }
}

export default bagsReducer;
