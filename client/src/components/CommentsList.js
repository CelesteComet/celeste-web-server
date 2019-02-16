import React from 'react';
import CommentsListItem from './CommentsListItem';

function CommentsList({comments, user, deleteComment}) {
  return (
    <ul className="comments-list">
      {comments.map((comment) => {
        return <CommentsListItem 
                  key={comment.id} 
                  comment={comment} 
                  user={user} 
                  deleteComment={deleteComment}/>
      })}
    </ul>
  );
}

export default CommentsList;