import React from 'react';
import { Link } from 'react-router-dom';

export default function BagsIndexPageItem({bag: {id, name, brand, image_url, created_by}}) {
  return (
    <Link to={`/bags/${id}`}>
      <div>
        <img src={image_url} alt={name} />
      </div>  
    </Link>
  );
}

