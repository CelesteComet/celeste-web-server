import React from 'react';
import md5 from 'md5';

function GravatarImage({gravatarHash}) {
  return <img src={`https://www.gravatar.com/avatar/${gravatarHash}`} alt={''} />
}

export default GravatarImage;