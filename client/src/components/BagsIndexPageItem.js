import React from 'react';

export default function BagsIndexPageItem({bag: {name, brand, image_url, created_by}}) {
  return (
    <div>
      <h1>{name}</h1>
      <p>{brand}</p>
      <img src={image_url} alt={name} style={{width: "100px", height: "100px"}} />
      <p>Created By: {created_by}</p>
    </div>  
  );
}

