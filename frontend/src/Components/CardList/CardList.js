import React from 'react';
import Card from './Card';

const CardList = ({ inventory, medicine }) => {
  return (
    <div>
      <h2>Suppliers close to you are:</h2>
      {
        inventory.map((supplier, i) => {
          return (
            <Card
              name={inventory[i].name}
              distance={inventory[i].distance}
              price_per_unit={inventory[i].price_per_unit}
              quantity={inventory[i].quantity}
              address={inventory[i].address}
              phone_number={inventory[i].phone_number}
              />
          );
        }).sort(function (a, b) {
          return a.distance - b.distance;
        })
      }
    </div>
  );
}

export default CardList;