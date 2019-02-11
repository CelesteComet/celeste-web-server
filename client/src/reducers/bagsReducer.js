import { 
  FETCH_BAGS, 
  RECEIVE_BAGS,
  RECEIVE_BAG
} from '../actions/bagActions';

const bagsReducer = (state = [], action) => {
  switch (action.type) {
    case RECEIVE_BAGS:
      return action.payload;
    case RECEIVE_BAG:
      return action.payload;
    default:
      return state
  }
}

export default bagsReducer;
