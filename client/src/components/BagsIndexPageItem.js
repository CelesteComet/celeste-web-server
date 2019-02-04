import React from 'react';

export default function BagsIndexPageItem({bag: {name, brand, image_url, created_by}}) {
  return (
    <div>
      <img src={image_url} alt={name} />
      <h1>{name}</h1>
      <p>{brand}</p>
      <p>Created By: {created_by}</p>
    </div>  
  );
}

