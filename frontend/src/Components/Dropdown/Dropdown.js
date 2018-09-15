import React, { Component } from 'react';
import './Dropdown.css';

class Dropdown extends Component {
  constructor(props) {
    super(props);
    
    this.state = {
      showMenu: false,
    };
    
    this.showMenu = this.showMenu.bind(this);
    this.closeMenu = this.closeMenu.bind(this);
  }
  
  showMenu(event) {
    event.preventDefault();
    
    this.setState({ showMenu: true }, () => {
      document.addEventListener('click', this.closeMenu);
    });
  }
  
  closeMenu(event) {
    if (!this.dropdownMenu.contains(event.target)) {
      this.setState({ showMenu: false }, () => {
        document.removeEventListener('click', this.closeMenu);
      });  
    }
  }
  // Make those buttons a loop dumbass
  render() {
    return (
      <div className="tc">
        <button onClick={this.showMenu}>
          Show menu
        </button>
        
        {
          this.state.showMenu
            ? (
              <div
                className="menu"
                ref={(element) => {
                  this.dropdownMenu = element;
                }}
              >
              <button className="dropdown-btn f6 link dim br1 ph3 pv2 mb2 dib white bg-red" onClick={() => this.props.onClick(1)}> Aspirin </button>
              <button className="dropdown-btn f6 link dim br1 ph3 pv2 mb2 dib white bg-red" onClick={() => this.props.onClick(2)}> Amiloride </button>
              <button className="dropdown-btn f6 link dim br1 ph3 pv2 mb2 dib white bg-red" onClick={() => this.props.onClick(3)}> Amiodarone </button>
              <button className="dropdown-btn f6 link dim br1 ph3 pv2 mb2 dib white bg-red" onClick={() => this.props.onClick(4)}> Bisoprolol </button>
              <button className="dropdown-btn f6 link dim br1 ph3 pv2 mb2 dib white bg-red" onClick={() => this.props.onClick(5)}> Clopidogrel </button>
              <button className="dropdown-btn f6 link dim br1 ph3 pv2 mb2 dib white bg-red" onClick={() => this.props.onClick(6)}> Digoxin </button>
              <button className="dropdown-btn f6 link dim br1 ph3 pv2 mb2 dib white bg-red" onClick={() => this.props.onClick(7)}> Furosemide </button>
              <button className="dropdown-btn f6 link dim br1 ph3 pv2 mb2 dib white bg-red" onClick={() => this.props.onClick(8)}> Losartan </button>
              <button className="dropdown-btn f6 link dim br1 ph3 pv2 mb2 dib white bg-red" onClick={() => this.props.onClick(9)}> Methyldopa </button>
              <button className="dropdown-btn f6 link dim br1 ph3 pv2 mb2 dib white bg-red" onClick={() => this.props.onClick(10)}> Nifedipine </button>
              <button className="dropdown-btn f6 link dim br1 ph3 pv2 mb2 dib white bg-red" onClick={() => this.props.onClick(11)}> Spironolactone </button>
              <button className="dropdown-btn f6 link dim br1 ph3 pv2 mb2 dib white bg-red" onClick={() => this.props.onClick(12)}> Streptokinase </button>
              <button className="dropdown-btn f6 link dim br1 ph3 pv2 mb2 dib white bg-red" onClick={() => this.props.onClick(13)}> Verapamil </button>
                
              </div>
            )
            : (
              null
            )
        }
      </div>
    );
  }
}

export default Dropdown;