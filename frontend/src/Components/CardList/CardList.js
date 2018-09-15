import React from 'react';
import Card from './Card';

const CardList = ({ inventory, medicine }) => {
	return (
    <div>
      <h2>{medicine} suppliers close to you are:</h2>
      {
        inventory.map((supplier, i) => {
          return (
            <Card
              id={inventory[i].id}
              name={inventory[i].name}
              quantity={inventory[i].quantity}
              price_per_unit={inventory[i].price_per_unit}
              />
          );
        })
      }
    </div>
  );
}

export default CardList;

// GET INFO ON CARDS