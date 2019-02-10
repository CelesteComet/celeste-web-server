import { CLEAR_ERRORS , RECEIVE_ERRORS } from '../actions/errorActions';

const errorsReducer = (state = [], action) => {
  console.log(action.type);
  console.log(action.payload)
  switch (action.type) {
    case RECEIVE_ERRORS:
      return action.payload;
    case CLEAR_ERRORS:
      return [];
    default:
      return state
  }
}

export default errorsReducer;
