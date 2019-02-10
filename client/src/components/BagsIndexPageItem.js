import React from 'react';

export default function BagsIndexPageItem({bag: {name, brand, image_url, created_by}}) {
  return (
    <div>
      <img src={image_url} alt={name} />
    </div>  
  );
}

