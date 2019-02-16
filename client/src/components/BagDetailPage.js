import React, { Component } from 'react';
import { connect }  from 'react-redux';
import { fetchBag } from '../actions/bagActions';
import { postComment, fetchComments, deleteComment } from '../actions/commentActions';
import PageLoadSpinner from './PageLoadSpinner';
import { Comment } from '../models';
import CommentsList from './CommentsList';

class BagDetailPage extends Component {
  constructor(props) {
    super(props)

    this.state = {
      loading: true,
      comment: new Comment
    }

    this.handleChange = this.handleChange.bind(this);
    this.handleCommentSubmission = this.handleCommentSubmission.bind(this);
  }

  componentDidMount() {
    const { fetchBag, fetchComments } = this.props;
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

  handleChange(e) {
    let comment = Object.assign({}, this.state.comment);
    comment.content = e.target.value;
    this.setState({ comment })
  }

  handleCommentSubmission() {
    let comment = this.state.comment;
    comment.item_id = this.props.bag.id
    this.props.postComment(comment);
  }

  render() {
    if (!this.props.bag || this.state.loading) {
      return <PageLoadSpinner />
    }

    const { name, brand, image_url, created_by_member } = this.props.bag;
    const { user, deleteComment } = this.props;

    return (
      <div className="wrapper bag-detail-page">
        <h1>{ name }</h1>
        <h1>{ brand }</h1>
        <h1>created by {created_by_member}</h1>
        <img src={image_url} alt={name} style={{width: '400px'}}></img>
        { user && 
          <div>
            <textarea onChange={this.handleChange} /> 
            <button className="button__submit" onClick={this.handleCommentSubmission}>POST COMMENT</button>
          </div>
        }
        <CommentsList comments={this.props.comments} user={user} deleteComment={deleteComment}/>
      </div>
    );
  }
}

const mSTP = state => {
  return { 
    bag: state.bags[0] ? state.bags[0] : null,
    comments: state.comments,
    user: state.user
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
    deleteComment: (id) => {
      dispatch(deleteComment(id))
    }
  }
}

export default connect(mSTP, mDTP)(BagDetailPage);