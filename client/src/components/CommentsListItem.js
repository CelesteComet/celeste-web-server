import React from 'react';
import moment from 'moment'
import GravatarImage from './GravatarImage';

function CommentsListItem({comment: {id, created_by, created_by_member, created_at, content, gravatar_hash}, user, deleteComment}) {
  const displayDelete = (user && user.id === created_by);
  return (
    <div className="item">

      <div className="item__avatar-container" >
        <GravatarImage gravatarHash={gravatar_hash} />
        <div className="item__line-down"></div>
      </div>

      <div className="item__content-container">
        <div>
          <h2><a>{ created_by_member }</a></h2>
          <small>{ moment(created_at).fromNow() }</small>
        </div>
        <p>{ content }</p>
      </div>

      <div className="item__utilities">
        { displayDelete && <button onClick={deleteComment.bind(null, id)}>DELETE</button> }
      </div>
    </div>


  );
}

export default CommentsListItem;