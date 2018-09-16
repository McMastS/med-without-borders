import React from 'react';

const Card = ({ name, distance, price, quantity }) => {
  return (
    <div className='tc grow bg-red br3 pa3 ma2 dib bw2 shadow-5'>
      <div>
        <h2>{name}</h2>
        <p>{distance}</p>
        <p>{price}</p>
        <p>{quantity}</p>
      </div>
    </div>
  );
}

export default Card;

// Possible card upgrade
// <article class="mw5 center bg-white br3 pa3 pa4-ns mv3 ba b--black-10">
//   <div class="tc">
//     <img src="http://tachyons.io/img/avatar_1.jpg" class="br-100 h4 w4 dib ba b--black-05 pa2" title="Photo of a kitty staring at you">
//     <h1 class="f3 mb2">Mimi W.</h1>
//     <h2 class="f5 fw4 gray mt0">CCO (Chief Cat Officer)</h2>
//   </div>
// </article>