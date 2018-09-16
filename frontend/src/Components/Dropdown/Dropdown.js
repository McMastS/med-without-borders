import React, { Component } from 'react';
import './Dropdown.css';

class Dropdown extends Component {
  constructor(props) {
    super(props);
    };
  // Make those buttons a loop dumbass
  render() {
    return (
      <div className="tc pa4">
        <button className="f6 link dim br-pill ph3 pv2 mb2 dib white bg-black" onClick={() => this.props.onClick(1)}> Aspirin </button>
        <button className="f6 link dim br-pill ph3 pv2 mb2 dib white bg-black" onClick={() => this.props.onClick(2)}> Amiloride </button>
        <button className="f6 link dim br-pill ph3 pv2 mb2 dib white bg-black" onClick={() => this.props.onClick(3)}> Amiodarone </button>
        <button className="f6 link dim br-pill ph3 pv2 mb2 dib white bg-black" onClick={() => this.props.onClick(4)}> Bisoprolol </button>
        <button className="f6 link dim br-pill ph3 pv2 mb2 dib white bg-black" onClick={() => this.props.onClick(5)}> Clopidogrel </button>
        <button className="f6 link dim br-pill ph3 pv2 mb2 dib white bg-black" onClick={() => this.props.onClick(6)}> Digoxin </button>
        <button className="f6 link dim br-pill ph3 pv2 mb2 dib white bg-black" onClick={() => this.props.onClick(7)}> Furosemide </button>
        <button className="f6 link dim br-pill ph3 pv2 mb2 dib white bg-black" onClick={() => this.props.onClick(8)}> Losartan </button>
        <button className="f6 link dim br-pill ph3 pv2 mb2 dib white bg-black" onClick={() => this.props.onClick(9)}> Methyldopa </button>
        <button className="f6 link dim br-pill ph3 pv2 mb2 dib white bg-black" onClick={() => this.props.onClick(10)}> Nifedipine </button>
        <button className="f6 link dim br-pill ph3 pv2 mb2 dib white bg-black" onClick={() => this.props.onClick(11)}> Spironolactone </button>
        <button className="f6 link dim br-pill ph3 pv2 mb2 dib white bg-black" onClick={() => this.props.onClick(12)}> Streptokinase </button>
        <button className="f6 link dim br-pill ph3 pv2 mb2 dib white bg-black" onClick={() => this.props.onClick(13)}> Verapamil </button>
      </div>
    );
  }
}

export default Dropdown;