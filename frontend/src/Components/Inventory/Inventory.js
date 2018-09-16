import React, { Component } from 'react';
import Dropdown from '../Dropdown/Dropdown.js';
import CardList from '../CardList/CardList.js';

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
      body: JSON.stringify({
        id: value,
        uuid: localStorage.getItem('uuid'),
      })
    })
        .then(res => res.json())
        .then(data => this.setState({ inventory : data, isLoading : false, }))
        
     // this.setState({ inventory: [{"id":"5b9cdc54cb18d0095c485dce","name":"Best Buy","quantity":30,"price_per_unit":10.3, medType: 'C0001'},{"id":"5b9cdc73cb18d00ddb4a2552","name":"Best Buy","quantity":30,"price_per_unit":10.3},{"id":"5b9cddd7cb18d01da2a9e2fa","name":"Best Buy","quantity":30,"price_per_unit":10.3},{"id":"5b9cde42cb18d022dd02f356","name":"Best Buy","quantity":30,"price_per_unit":10.3}]  })
    }
  }

    componentDidMount () {
      //this.fetchSuppliers();
      //this.setState({ inventory: [{"id":"5b9cdc54cb18d0095c485dce","name":"Best Buy","quantity":30,"price_per_unit":10.3}]})
    }

    render() {
      const { inventory, isLoading } = this.state;

      if (isLoading) {
        return <p> Loading... </p>
      }

      return (
        <div className="tc">
          <Dropdown onClick={this.fetchSuppliers} />
          <CardList medicine={this.state.medicine} inventory={inventory}/>
        </div>
      );
    }
  }

export default Inventory;