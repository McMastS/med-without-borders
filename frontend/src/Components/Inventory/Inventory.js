import React, { Component } from 'react';
import Dropdown from '../Dropdown/Dropdown.js';
import CardList from '../CardList/CardList.js';
import Navbar from '../Navbar/Navbar.js';

const API = 'http://api.medwithoutborders.org/medicine/'

class Inventory extends Component {
  constructor(props) {
      super(props);
      this.state = {
        inventory: [],
        isLoading: false,
        medicine: 'Aspirin',
        medType: 1,
      };
    }
    // TODO fix medicine state change so that the inventory updates when clikcing the dropdown, consider going to list and searching
    fetchSuppliers = (value) => {
        fetch('http://api.medwithoutborders.org/medicine', {
        method: 'post',
        body: JSON.stringify(
          {
            id: value,
            uuid: localStorage.getItem('uuid'),
          }
        )
      })
        .then(res => res.json())
        .then(data => this.setState({ inventory : data, isLoading : false, }))
        
     // this.setState({ inventory: [{"id":"5b9cdc54cb18d0095c485dce","name":"Rocky View Hospital","quantity":30,"price_per_unit":10.3, "distance": 56.4, "medType": 1}, {"id":"5b9cdc54cb18d0095c485dce","name":"General Hospital","quantity":200,"price_per_unit":10.3, "distance": 63.7, "medType":2}]})
    }
  

    componentDidMount () {
      this.fetchSuppliers();
      //this.setState({ inventory: [{"id":"5b9cdc54cb18d0095c485dce","name":"Best Buy","quantity":30,"price_per_unit":10.3}]})
    }

    render() {
      const { inventory, isLoading } = this.state;

      if (isLoading) {
        return <p> Loading... </p>
      }

      return (
        <div className="tc">
          <Navbar />
          <Dropdown onClick={this.fetchSuppliers} />
          <CardList medicine={this.state.medicine} inventory={inventory}/>
        </div>
      );
    }
  }

export default Inventory;