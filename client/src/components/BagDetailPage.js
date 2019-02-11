import React, { Component } from 'react';
import { connect }  from 'react-redux';
import { fetchBag } from '../actions/bagActions';
import { postComment, fetchComments } from '../actions/commentActions';
import PageLoadSpinner from './PageLoadSpinner';
import MyStatefulEditor from './MyStatefulEditor';

class BagDetailPage extends Component {
  constructor(props) {
    super(props)

    this.state = {
      loading: true,
      comment: ""
    }

    this.handleChange = this.handleChange.bind(this);
    this.handleCommentSubmission = this.handleCommentSubmission.bind(this);
  }

  componentDidMount() {
    const { fetchBag } = this.props;
    const { id } = this.props.match.params;
    fetchBag(id);
    fetchComments(id);
  }

  componentDidUpdate(prevState, prevProps) {
    if (prevState.bag != this.props.bag) {
      this.setState({
        loading: false
      })
    }
  }

  handleChange(value) {
    this.setState({
      comment: value
    })
  }

  handleCommentSubmission() {
    let comment = {
      content: this.state.comment,
      created_by: 165,
      item_id: this.props.bag.id
    };

    this.props.postComment(comment)
    
  }

  render() {
    if (!this.props.bag || this.state.loading) {
      return <PageLoadSpinner />
    }

    const { name, brand, image_url, created_by_member } = this.props.bag;

    return (
      <div className="wrapper">
        <h1>{ name }</h1>
        <h1>{ brand }</h1>
        <h1>created by {created_by_member}</h1>
        <img src={image_url} alt={name} style={{width: '400px'}}></img>
        <MyStatefulEditor onChange={this.handleChange}/>
        <div onClick={this.handleCommentSubmission}>POST COMMENT</div>
        { this.props.comments.map((comment) => {
          return <p>{comment.created_by}: {comment.content}</p>
        })}
      </div>
    );
  }
}

const mSTP = state => {
  return { 
    bag: state.bags[0] ? state.bags[0] : null,
    comments: state.comments
  }
}

const mDTP = dispatch => {
  return { 
    fetchBag: (i) => {
      dispatch(fetchBag(i))
    },
    fetchComments: (id) => {
      dispatch(fetchComments(id))
    },
    postComment: (comment) => {
      dispatch(postComment(comment))
    },
  }
}

export default connect(mSTP, mDTP)(BagDetailPage);