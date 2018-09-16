import React from 'react';

const Card = ({ name, distance, price, quantity }) => {
  return (
    <article className="mw5 center bg-white br3 pa3 pa4-ns mv3 ba b--black-10">
  		<div className="tc">
		    <img src="http://tachyons.io/img/avatar_1.jpg" className="br-100 h4 w4 dib ba b--black-05 pa2" title="Photo of a kitty staring at you"/>
		    <h1 className="f3 mb2">{name}</h1>
		    <h2 className="f5 fw4 gray mt0">{distance}</h2>
		    <h2 className="f5 fw4 gray mt0">{price}</h2>
		    <h2 className="f5 fw4 gray mt0">{quantity}</h2>
	  	</div>
		</article>
  );
}

export default Card;

