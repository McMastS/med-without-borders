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
              name={inventory[i].name}
              distance={inventory[i].distance}
              price_per_unit={inventory[i].price_per_unit}
              quantity={inventory[i].quantity}
              />
          );
        })
      }
    </div>
  );
}

export default CardList;